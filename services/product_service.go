package services

import (
	"grocery-purchase/models"
	"grocery-purchase/repositories"
	"database/sql"
)

// CreateProduct adds a new product to the database.
func CreateProduct(db *sql.DB, product *models.Product) error {
	return repositories.CreateProduct(db, product)
}
