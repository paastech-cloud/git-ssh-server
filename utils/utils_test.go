package utils_test

import (
	"testing"

	"github.com/paastech-cloud/git-ssh-server/utils"
	"github.com/stretchr/testify/assert"
)

// test the GetRepoName function
func TestGetRepoName(t *testing.T) {
	// test with a command in the form of `git-receive-pack /<repo-name>`
	command := "git-receive-pack /<repo-name>"
	repoName, err := utils.GetRepoName(command)

	assert.Nil(t, err)
	assert.Equal(t, "<repo-name>", repoName)

	// test with a command in the form of `git-receive-pack <repo-name>`
	command = "git-receive-pack <repo-name>"
	repoName, err = utils.GetRepoName(command)

	assert.Nil(t, err)
	assert.Equal(t, "<repo-name>", repoName)

	// test with a command in the form of `git-receive-pack /<repo-name>/`
	command = "git-receive-pack /<repo-name>/"
	repoName, err = utils.GetRepoName(command)

	assert.Nil(t, err)
	assert.Equal(t, "<repo-name>", repoName)

	// test with a command in the form of `git-receive-pack <repo-name>`
	command = "git-receive-pack <repo-name>/"
	repoName, err = utils.GetRepoName(command)

	assert.Nil(t, err)
	assert.Equal(t, "<repo-name>", repoName)

	// test with a command in the form of `not receive pack`
	command = "anythingelse <repo-name>"
	repoName, err = utils.GetRepoName(command)

	assert.NotNil(t, err)
	assert.Equal(t, "", repoName)

	// test with a command in the form of `git-receive-pack <repo-name> <repo-name>`
	command = "git-receive-pack <repo-name> <repo-name>"
	repoName, err = utils.GetRepoName(command)

	assert.NotNil(t, err)
	assert.Equal(t, "", repoName)

	// test with a command in the form of `git-receive-pack`
	command = "git-receive-pack"
	repoName, err = utils.GetRepoName(command)

	assert.NotNil(t, err)
	assert.Equal(t, "", repoName)
}
