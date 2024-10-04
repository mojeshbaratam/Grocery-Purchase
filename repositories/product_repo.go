package repositories

import (
	"database/sql"
	"grocery-purchase/models"
)

func CreateProduct(db *sql.DB, product *models.Product) error {
	query := "INSERT INTO products (name, price, store_id) VALUES (?, ?, ?)"
	_, err := db.Exec(query, product.Name, product.Price, product.StoreID)
	return err
}
