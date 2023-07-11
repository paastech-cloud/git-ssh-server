package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path"

	"github.com/rs/zerolog/log"
)

/**
* This function generates a RSA Key pair
* Usage: GenerateHostKey("example/key")
*
* it will generate:
*   - example/key
*   - example/key.pub
 */
<<<<<<< HEAD
func GenerateKeyPair(filepath string) {
=======
func GenerateHostKey(filepath string) {
>>>>>>> 5ad0b53 (feat: ssh key pair generation on init)
	// Creating parent directory if it doesn't exists
	MkParentdir(filepath)

	// Generate a new RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal().Err(err).Msg("Error while generating private key")
		return
	}

	// Convert private key with PEM format
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// Create a file for the private key
	privateKeyFile, err := os.Create(filepath)
	if err != nil {
		log.Fatal().Err(err).Msg("Error while creating private key file")
		return
	}
	defer privateKeyFile.Close()

	// Writing private key to file with PEM format
	err = pem.Encode(privateKeyFile, privateKeyPEM)
	if err != nil {
		log.Fatal().Err(err).Msg("Error while writting private key file")
		return
	}

	log.Info().Msg("Private key successfully generated at: " + filepath)

	publicKeyPath := filepath + ".pub"

	// Convert public key to PEM format
	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
	}

	// Create a file for the public key
	publicKeyFile, err := os.Create(publicKeyPath)
	if err != nil {
		log.Fatal().Err(err).Msg("Error while creating public key file")
		return
	}
	defer publicKeyFile.Close()

	// Writting public key to file with PEM format
	err = pem.Encode(publicKeyFile, publicKeyPEM)
	if err != nil {
		log.Fatal().Err(err).Msg("Error while writting public key file")
		return
	}

	log.Info().Msg("Public key successfully generated at: " + publicKeyPath)
}

func DirExists(path_dir string) bool {
	info, err := os.Stat(path_dir)
	if err == nil {
		if info.IsDir() {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func MkParentdir(filepath string) {
	parentdir := path.Dir(filepath)
	if DirExists(parentdir) {
		return
	}
	err := os.MkdirAll(parentdir, os.ModePerm)
	if err != nil {
		log.Fatal().Err(err).Msg("Error while making key directory")
		return
	}
}
