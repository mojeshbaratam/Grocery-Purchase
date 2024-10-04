package services

import (
	"grocery-purchase/models"
	"grocery-purchase/repositories"
	"database/sql"
)

// CreateStore adds a new store to the database.
func CreateStore(db *sql.DB, store *models.Store) error {
	return repositories.CreateStore(db, store)
}

// GetStores retrieves all stores from the database.
func GetStores(db *sql.DB) ([]models.Store, error) {
	return repositories.GetStores(db)
}
