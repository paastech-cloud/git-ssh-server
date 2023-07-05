package handlers

import (
	_ "github.com/lib/pq"
	"github.com/paastech-cloud/git-ssh-server/logger"
	"github.com/paastech-cloud/git-ssh-server/utils"

	"github.com/gliderlabs/ssh"
)

// AuthenticateUser authenticates the user by checking if the user added their ssh key to the database
func AuthenticateUser(ctx ssh.Context, key ssh.PublicKey) bool {
	fullKeyString := utils.ParsePublicKey(key)

	logger.InfoLogger.Println("user connected with public key: ", fullKeyString)

	userAuthorized, err := utils.IsUserAuthorized(fullKeyString)

	if err != nil {
		logger.ErrorLogger.Println(err)
		return false
	}

	if !userAuthorized {
		logger.WarningLogger.Printf("user with public key %s unauthorized", fullKeyString)
		return false
	}

	logger.InfoLogger.Printf("user with public key %s authorized", fullKeyString)

	return true
}
