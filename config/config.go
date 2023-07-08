package config

import (
	"fmt"
	"os"
)

const (
	GIT_REPOSITORIES_FULL_BASE_PATH_KEY = "GIT_REPOSITORIES_FULL_BASE_PATH"
	GIT_HOST_SIGNER_PATH_KEY            = "GIT_HOST_SIGNER_PATH"
	GIT_POSTGRESQL_USERNAME_KEY         = "GIT_POSTGRESQL_USERNAME"
	GIT_POSTGRESQL_PASSWORD_KEY         = "GIT_POSTGRESQL_PASSWORD"
	GIT_POSTGRESQL_DATABASE_NAME_KEY    = "GIT_POSTGRESQL_DATABASE_NAME"
	GIT_POSTGRESQL_PORT_KEY             = "GIT_POSTGRESQL_PORT"
	GIT_POSTGRESQL_HOST_KEY             = "GIT_POSTGRESQL_HOST"
)

var (
	RepositoriesBasePath string
	PathToHostSigner     string
	PostgreSQLUsername   string
	PostgreSQLPassword   string
	PostgreSQLDatabase   string
	PostgreSQLPort       string
	PostgreSQLHost       string
)

// CheckConfig checks if all required environment variables are set
func CheckConfig() error {
	// make an array of all required environment variables
	requiredEnvVars := []string{
		GIT_REPOSITORIES_FULL_BASE_PATH_KEY,
		GIT_HOST_SIGNER_PATH_KEY,
		GIT_POSTGRESQL_USERNAME_KEY,
		GIT_POSTGRESQL_PASSWORD_KEY,
		GIT_POSTGRESQL_DATABASE_NAME_KEY,
		GIT_POSTGRESQL_PORT_KEY,
		GIT_POSTGRESQL_HOST_KEY,
	}

	// make an array of all missing environment variables
	missingEnvVars := []string{}

	// loop through all required environment variables
	// and check if they are set, if not, add them to the missingEnvVars array
	for _, envVar := range requiredEnvVars {
		// if the environment variable is not set, return an error
		if os.Getenv(envVar) == "" {
			missingEnvVars = append(missingEnvVars, envVar)
		}
	}

	// if there are missing environment variables, return an error
	if len(missingEnvVars) > 0 {
		return fmt.Errorf("missing environment variables: %v", missingEnvVars)
	}

	// set the global variables
	RepositoriesBasePath = os.Getenv(GIT_REPOSITORIES_FULL_BASE_PATH_KEY)
	PathToHostSigner = os.Getenv(GIT_HOST_SIGNER_PATH_KEY)
	PostgreSQLUsername = os.Getenv(GIT_POSTGRESQL_USERNAME_KEY)
	PostgreSQLPassword = os.Getenv(GIT_POSTGRESQL_PASSWORD_KEY)
	PostgreSQLDatabase = os.Getenv(GIT_POSTGRESQL_DATABASE_NAME_KEY)
	PostgreSQLPort = os.Getenv(GIT_POSTGRESQL_PORT_KEY)
	PostgreSQLHost = os.Getenv(GIT_POSTGRESQL_HOST_KEY)

	return nil
}
