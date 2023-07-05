package utils

import (
	"errors"
	"strings"
)

// Extract the repository name from the command
//
// The command can be in the form of:
// - git-receive-pack '<repo-name>'
// - git-receive-pack '/<repo-name>'
// - git-receive-pack '<repo-name>/'
// - git-receive-pack '/<repo-name>/'
//
// if the command doesn't have the right format, an error is returned
func GetRepoName(command string) (string, error) {
	if !strings.HasPrefix(command, "git-receive-pack") {
		return "", errors.New("command executed was not git-receive-pack")
	}

	operands := strings.Split(command, " ")

	if len(operands) != 2 {
		return "", errors.New("command executed was not formatted correctly")
	}

	repoName := operands[1]

	repoName = strings.Trim(repoName, "'")
	repoName = strings.TrimPrefix(repoName, "/")
	repoName = strings.TrimSuffix(repoName, "/")

	return repoName, nil
}
