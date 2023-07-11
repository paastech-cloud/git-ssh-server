package utils

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/paastech-cloud/git-ssh-server/config"
)

// Returns a connection to the database
func getConnection() (*sql.DB, error) {
	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.PostgreSQLUsername,
		config.PostgreSQLPassword,
		config.PostgreSQLHost,
		config.PostgreSQLPort,
		config.PostgreSQLDatabase,
	)

	db, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil
}

// Queries the database to check if the user added their ssh key to the database
func IsUserAuthorized(key string) (bool, error) {
	db, err := getConnection()

	if err != nil {
		return false, err
	}

	defer db.Close()

	var count int

	err = db.QueryRow(
		"SELECT COUNT(*) FROM ssh_keys WHERE value LIKE $1",
		key+" %",
	).Scan(&count)

	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

// Queries the database to check if the user is authorized to access the repository
// If the user is authorized, return true
// If the user is not authorized, return false
//
// The user being authorized means that the ssh key is associated with a user
// that has access to the repository
func CanUserEditRepository(key string, repoName string) (bool, error) {
	db, err := getConnection()

	if err != nil {
		log.Println(err)
		return false, err
	}

	defer db.Close()

	var count int

	err = db.QueryRow(
		"SELECT COUNT(*) FROM ssh_keys WHERE value = $1 AND user_id IN (SELECT id FROM users WHERE user_id IN (SELECT user_id FROM projects WHERE id = $2))",
		key,
		repoName,
	).Scan(&count)

	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}
