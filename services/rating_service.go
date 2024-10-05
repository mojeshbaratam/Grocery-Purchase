// services/rating_service.go
package services

import (
	"database/sql"
	"grocery-purchase/models"
	"grocery-purchase/repositories"
)

// Create a new rating
func CreateRating(db *sql.DB, rating *models.Rating) error {
	return repositories.CreateRating(db, rating)
}

// Get ratings for a product
func GetRatingsByProductID(db *sql.DB, productID int) ([]models.Rating, error) {
	return repositories.GetRatingsByProductID(db, productID)
}