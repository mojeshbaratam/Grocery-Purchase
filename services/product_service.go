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

// GetProducts retrieves all products from the database.
func GetProducts(db *sql.DB) ([]models.Product, error) {
	return repositories.GetProducts(db)
}

func GetProductsByName(db *sql.DB, name string) ([]models.Product, error) {
	return repositories.GetProductsByName(db, name)
}

func UpdateProduct(db *sql.DB, name string, newName string, newPrice float64, newStoreId int) (bool, error) {
	return repositories.UpdateProduct(db, name, newName, newPrice, newStoreId)
}

func RemoveProduct(db *sql.DB, productName string) (bool, error) {
	return repositories.RemoveProduct(db, productName)
}