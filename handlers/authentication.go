package handlers

import (
	_ "github.com/lib/pq"
	"github.com/paastech-cloud/git-ssh-server/utils"
	"github.com/rs/zerolog/log"

	"github.com/gliderlabs/ssh"
)

// AuthenticateUser authenticates the user by checking if the user added their ssh key to the database
func AuthenticateUser(ctx ssh.Context, key ssh.PublicKey) bool {
	fullKeyString := utils.ParsePublicKey(key)

	log.Debug().Msg("user connected with public key: " + fullKeyString)

	userAuthorized, err := utils.IsUserAuthorized(fullKeyString)

	if err != nil {
		log.Error().Err(err).Msg("something went wrong while checking if user is authorized")
		return false
	}

	if !userAuthorized {
		log.Debug().Msgf("user with public key %s unauthorized", fullKeyString)
		return false
	}

	log.Debug().Msgf("user with public key %s authorized", fullKeyString)

	return true
}
