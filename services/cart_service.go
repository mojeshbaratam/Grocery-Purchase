package services

import (
	"database/sql"
	"grocery-purchase/models"
	"grocery-purchase/repositories"
)

// AddProductToCartService handles adding a product to the cart.
func AddProductToCartService(db *sql.DB, cart *models.Cart) error {
	return repositories.AddProductToCart(db, cart)
}

// GetCartService retrieves the cart contents for a specific user.
func GetCartService(db *sql.DB, userID int) ([]models.Cart, error) {
	return repositories.GetCartByUserID(db, userID)
}

// ClearCartService clears the cart after purchase.
func ClearCartService(db *sql.DB, userID int) error {
	return repositories.ClearCart(db, userID)
}
