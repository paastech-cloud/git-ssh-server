package services

import (
	"database/sql"
	"fmt"
	"os"
)

// Queries the database to check if the user is authorized to access the repository
// If the user is authorized, return true
// If the user is not authorized, return false
//
// The user being authorized means that the ssh key is associated with a user
// that has access to the repository
func IsUserAuthorized(key string, repoName string) (bool, error) {
	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("config.GIT_POSTGRESQL_USERNAME"),
		os.Getenv("config.GIT_POSTGRESQL_PASSWORD"),
		os.Getenv("config.GIT_POSTGRESQL_HOST"),
		os.Getenv("config.GIT_POSTGRESQL_PORT"),
		os.Getenv("config.GIT_POSTGRESQL_DATABASE_NAME"),
	)

	db, err := sql.Open("postgres", dbUrl)

	if err != nil {
		return false, err
	}

	defer db.Close()

	var count int

	err = db.QueryRow(
		"SELECT COUNT(*) FROM ssh_keys WHERE key = $1 AND user_id IN (SELECT id FROM users WHERE user_id IN (SELECT user_id FROM repositories WHERE id = $2))",
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
