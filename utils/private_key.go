package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
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
func GenerateKeyPair(filepath string) error {
	// Creating parent directory if it doesn't exists
	err := MkParentdir(filepath)
	if err != nil {
		return err
	}

	// Generate a new RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return fmt.Errorf("Error while generating private key: " + err.Error())
	}

	// Convert private key with PEM format
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// Create a file for the private key
	privateKeyFile, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("Error while creating private key file: " + err.Error())
	}
	defer privateKeyFile.Close()

	// Writing private key to file with PEM format
	err = pem.Encode(privateKeyFile, privateKeyPEM)
	if err != nil {
		return fmt.Errorf("Error while writting private key file: " + err.Error())
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
		return fmt.Errorf("Error while creating public key file: " + err.Error())
	}
	defer publicKeyFile.Close()

	// Writting public key to file with PEM format
	err = pem.Encode(publicKeyFile, publicKeyPEM)
	if err != nil {
		return fmt.Errorf("Error while writting public key file: " + err.Error())
	}

	log.Info().Msg("Public key successfully generated at: " + publicKeyPath)
	return nil
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

func MkParentdir(filepath string) error {
	parentdir := path.Dir(filepath)
	if DirExists(parentdir) {
		return nil
	}
	err := os.MkdirAll(parentdir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Error while making key directory: " + err.Error())
	}
	return nil
}
