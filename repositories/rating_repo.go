// repositories/rating_repo.go
package repositories

import (
	"database/sql"
	"grocery-purchase/models"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

// Create a new rating
func CreateRating(db *sql.DB, rating *models.Rating) error {
	query := "INSERT INTO ratings (user_id, product_id, rating, description) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, rating.UserID, rating.ProductID, rating.Rating, rating.Description)
	if err != nil {
		if err, ok := err.(*mysql.MySQLError); ok && err.Number == 1062 {
			return fmt.Errorf("rating by this user for this product already exists")
		}
		return err
	}
	return nil
}

// Get ratings for a product
func GetRatingsByProductID(db *sql.DB, productID int) ([]models.Rating, error) {
	query := "SELECT id, user_id, product_id, rating, description FROM ratings WHERE product_id = ?"
	rows, err := db.Query(query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ratings []models.Rating

	for rows.Next() {
		var rating models.Rating
		err := rows.Scan(&rating.ID, &rating.UserID, &rating.ProductID, &rating.Rating, &rating.Description)
		if err != nil {
			return nil, err
		}
		ratings = append(ratings, rating)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ratings, nil
}