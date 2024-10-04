package services

import (
	"grocery-purchase/models"
	"grocery-purchase/repositories"
	"database/sql"
)

// CreateAdmin creates a new admin in the database.
func CreateAdmin(db *sql.DB, admin *models.Admin) error {
	return repositories.CreateAdmin(db, admin)
}

// AuthenticateAdmin checks the admin's credentials and returns true if valid.
func AuthenticateAdmin(db *sql.DB, username, password string) (bool, error) {
	admin, err := repositories.GetAdminByUsername(db, username)
	if err != nil {
		return false, err
	}
	// Implement password check (hashing logic can be added here)
	return admin.Password == password, nil
}
