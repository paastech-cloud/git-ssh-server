package main

import (
	"os"

	"github.com/gliderlabs/ssh"
	"github.com/joho/godotenv"
	"github.com/paastech-cloud/git-ssh-server/config"
	"github.com/paastech-cloud/git-ssh-server/handlers"
	"github.com/paastech-cloud/git-ssh-server/logger"
	"github.com/paastech-cloud/git-ssh-server/utils"

	"github.com/rs/zerolog/log"
)

/**
 * TODO: Generate ssh key at start on the host volume,
 * if it is already there, just use it.
 *
 * Use env vars loaded in config to determine the path
 *
 * @description: SSH server with authorization via public key.
 */
func main() {
	err := godotenv.Load()

	logger.Setup()

	if err != nil && os.Getenv("ENV") == "development" {
		log.Fatal().Err(err).Msg("error loading .env file")
	}

	err = config.CheckConfig()

	if err != nil {
		log.Fatal().Err(err)
	}

	// check if path to host signer exists
	if _, err := os.Stat(config.PathToHostSigner); err != nil {
		log.Info().Err(err).Msg("host signer file does not exist, generating a key pair..")
		keyErr := utils.GenerateKeyPair(config.PathToHostSigner)
		if keyErr != nil {
			log.Fatal().Err(keyErr).Msg("Error while generating RSA Key Pair")
		}
	}

	// read host file from host
	hostSigner := ssh.HostKeyFile(config.PathToHostSigner)

	s := &ssh.Server{
		Addr:             ":2222",
		Handler:          handlers.ReceivePackHandler,
		PublicKeyHandler: handlers.AuthenticateUser,
	}

	// add host key to server
	err = s.SetOption(hostSigner)

	if err != nil {
		log.Fatal().Err(err).Msg("error setting host signer")
	}

	log.Info().Msg("server listening on port 2222")

	log.Fatal().Err(s.ListenAndServe())
}
