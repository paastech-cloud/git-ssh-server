package handlers

import (
	"os"
	"os/exec"

	"github.com/gliderlabs/ssh"
	"github.com/paastech-cloud/git-ssh-server/logger"
	"github.com/paastech-cloud/git-ssh-server/utils"
)

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
		_, _ = session.Stderr().Write([]byte("invalid command"))
		return
	}

	fullSshKey := utils.ParsePublicKey(session.PublicKey())

	CanUserEditRepository, err := utils.CanUserEditRepository(fullSshKey, repoName)

	if err != nil {
		logger.ErrorLogger.Println(err)
		_, _ = session.Stderr().Write([]byte("unknown error"))
		return
	}

	if !CanUserEditRepository {
		logger.WarningLogger.Printf("user with public key %s unauthorized to access repository %s", fullSshKey, repoName)
		_, _ = session.Stderr().Write([]byte("repository does not exist"))
		return
	}

	fullRepoPath := os.Getenv("GIT_REPOSITORIES_FULL_BASE_PATH") + "/" + repoName

	command := exec.Command("git-receive-pack", fullRepoPath)

	// set the stdout of the command to our session
	command.Stdout = session
	command.Stdin = session

	logger.InfoLogger.Printf("running command: %s", command.String())

	// run the command
	if err := command.Run(); err != nil {
		logger.ErrorLogger.Println(err)
		_, _ = session.Stderr().Write([]byte(err.Error()))
		return
	}
}
