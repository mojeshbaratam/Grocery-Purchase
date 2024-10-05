package repositories

import (
	"database/sql"
	"grocery-purchase/models"
)

// AddProductToCart adds a product to the user's cart.
func AddProductToCart(db *sql.DB, cart *models.Cart) error {
	query := "INSERT INTO carts (user_id, product_id, quantity) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE quantity = quantity + ?"
	_, err := db.Exec(query, cart.UserID, cart.ProductID, cart.Quantity, cart.Quantity)
	return err
}

// GetCartByUserID retrieves the cart for a specific user.
func GetCartByUserID(db *sql.DB, userID int) ([]models.Cart, error) {
	query := "SELECT id, user_id, product_id, quantity FROM carts WHERE user_id = ?"
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var carts []models.Cart
	for rows.Next() {
		var cart models.Cart
		if err := rows.Scan(&cart.ID, &cart.UserID, &cart.ProductID, &cart.Quantity); err != nil {
			return nil, err
		}
		carts = append(carts, cart)
	}

	return carts, nil
}

// ClearCart clears the cart for a specific user after purchase.
func ClearCart(db *sql.DB, userID int) error {
	query := "DELETE FROM carts WHERE user_id = ?"
	_, err := db.Exec(query, userID)
	return err
}
