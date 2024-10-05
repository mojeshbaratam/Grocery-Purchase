package services

import (
	"grocery-purchase/repositories"
	"database/sql"
)

// AuthenticateUser checks the user's credentials and returns true if valid.
func AuthenticateUser(db *sql.DB, username, password string) (bool, error) {
	user, err := repositories.GetUserByUsername(db, username)
	if err != nil {
		return false, err
	}
	// Implement password check (hashing logic can be added here)
	return user.Password == password, nil
}