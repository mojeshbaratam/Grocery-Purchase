package services

import (
	"grocery-purchase/models"
	"grocery-purchase/repositories"
	"database/sql"
)

// CreateUser registers a new user.
func CreateUser(db *sql.DB, user *models.User) error {
	return repositories.CreateUser(db, user)
}
