package main

import (
	"os"

	"github.com/gliderlabs/ssh"
	"github.com/joho/godotenv"
	"github.com/paastech-cloud/git-ssh-server/config"
	"github.com/paastech-cloud/git-ssh-server/handlers"
	"github.com/paastech-cloud/git-ssh-server/logger"

	"github.com/rs/zerolog/log"
)

/**
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
		log.Fatal().Err(err).Msg("host signer file does not exist")
	}

	hostSigner := ssh.HostKeyFile(config.PathToHostSigner)

	s := &ssh.Server{
		Addr:             ":2222",
		Handler:          handlers.ReceivePackHandler,
		PublicKeyHandler: handlers.AuthenticateUser,
	}

	s.SetOption(hostSigner)

	log.Info().Msg("server listening on port 2222")

	log.Fatal().Err(s.ListenAndServe())
}
