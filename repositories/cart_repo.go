package repositories

import (
	"database/sql"
	"grocery-purchase/models"
)

func CreateCartItem(db *sql.DB, cartItem *models.Cart) error {
	query := "INSERT INTO carts (user_id, product_id, quantity) VALUES (?, ?, ?)"
	_, err := db.Exec(query, cartItem.UserID, cartItem.ProductID, cartItem.Quantity)
	return err
}
