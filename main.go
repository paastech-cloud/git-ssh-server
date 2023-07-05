package main

import (
	"log"
	"os"

	"github.com/gliderlabs/ssh"
	"github.com/joho/godotenv"
	"github.com/paastech-cloud/git-ssh-server/config"
	"github.com/paastech-cloud/git-ssh-server/handlers"
	"github.com/paastech-cloud/git-ssh-server/logger"
)

/**
 * @description: SSH server with authorization via public key.
 */
func main() {
	err := godotenv.Load()

	if err != nil && os.Getenv("ENV") == "development" {
		log.Fatal("Error loading .env file")
	}

	logger.Init()

	err = config.CheckConfig()

	if err != nil {
		log.Fatal(err)
	}

	s := &ssh.Server{
		Addr:             ":2222",
		Handler:          handlers.ReceivePackHandler,
		PublicKeyHandler: handlers.AuthenticateUser,
	}

	logger.InfoLogger.Println("starting ssh server on port 2222...")

	log.Fatal(s.ListenAndServe())
}
