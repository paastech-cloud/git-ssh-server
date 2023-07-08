package handlers

import (
	"context"
	"io"
	"os"
	"os/exec"

	"github.com/gliderlabs/ssh"
	"github.com/paastech-cloud/git-ssh-server/logger"
	"github.com/paastech-cloud/git-ssh-server/utils"
)

func receivePack(session ssh.Session, repoName string) error {
	ctx, cancel := context.WithCancel(session.Context())
	defer cancel()

	fullRepoPath := os.Getenv("GIT_REPOSITORIES_FULL_BASE_PATH") + "/" + repoName
	cmd := exec.CommandContext(ctx, "git-receive-pack", fullRepoPath)

	// Set the environment variable IMAGE_NAME to the repository name
	cmd.Env = append(os.Environ(),
		"IMAGE_NAME="+repoName,
	)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}
	defer stdout.Close()

	stderr, err := cmd.StderrPipe()
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}
	defer stderr.Close()

	stdin, err := cmd.StdinPipe()
	if err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}
	defer stdin.Close()

	if err := cmd.Start(); err != nil {
		logger.ErrorLogger.Println(err)
		return err
	}

	go func() {
		defer stdin.Close()
		if _, err = io.Copy(stdin, session); err != nil {
			logger.ErrorLogger.Println(err)
			return
		}
	}()

	go func() {
		defer stdout.Close()
		if _, err = io.Copy(session, stdout); err != nil {
			logger.ErrorLogger.Println(err)
			return
		}
	}()

	go func() {
		defer stderr.Close()
		if _, err = io.Copy(session.Stderr(), stderr); err != nil {
			logger.ErrorLogger.Println(err)
			return
		}
	}()

	err = cmd.Wait()
	if err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			logger.ErrorLogger.Println(err)
		}
		return err
	}

	return nil
}

// This handler is once the user has authenticated
// It will run git-receive-pack only
//
// It takes the command from the environment variable GIT_SSH_COMMAND which is handled by gliderlabs/ssh
// It then parses the command to get the repository name and then runs git-receive-pack
// The output of git-receive-pack is then sent to the user
//
// If there is an error, it will be sent to the user
// The errors include:
// - The user is not allowed to access the repository (which appears for the user as if the repository does not exist)
// - The repository does not exist
// - The command is not git-receive-pack
// - The command is not formatted correctly
func ReceivePackHandler(session ssh.Session) {
	userCommand := session.RawCommand()

	repoName, err := utils.GetRepoName(userCommand)

	if err != nil {
		logger.WarningLogger.Println(err)
		return
	}

	fullSshKey := utils.ParsePublicKey(session.PublicKey())

	CanUserEditRepository, err := utils.CanUserEditRepository(fullSshKey, repoName)

	if err != nil {
		logger.ErrorLogger.Println(err)
		return
	}

	if !CanUserEditRepository {
		logger.WarningLogger.Printf("user with public key %s unauthorized to access repository %s", fullSshKey, repoName)
		return
	} else {
		logger.InfoLogger.Printf("user with public key %s authorized to access repository %s", fullSshKey, repoName)
	}

	err = receivePack(session, repoName)

	if err != nil {
		logger.ErrorLogger.Println(err)
		session.Exit(1)
		return
	}

	// TODO get the exit code of the command and send it to the user instead of exiting with 1
	if err := session.Exit(0); err != nil {
		logger.ErrorLogger.Println(err)
	}
}
