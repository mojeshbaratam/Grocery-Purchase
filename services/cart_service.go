package services

import (
	"grocery-purchase/models"
	"grocery-purchase/repositories"
	"database/sql"
)

// CreateCartItem adds an item to the user's cart.
func CreateCartItem(db *sql.DB, cartItem *models.Cart) error {
	return repositories.CreateCartItem(db, cartItem)
}
