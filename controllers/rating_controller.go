// controllers/rating_controller.go
package controllers

import (
	"database/sql"
	"net/http"
	"grocery-purchase/models"
	"grocery-purchase/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Create a new rating
func CreateRating(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	var rating models.Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := services.CreateRating(db, &rating); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Rating created successfully"})
}

// Get ratings for a product
func GetRatings(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	productID, err := strconv.Atoi(c.Param("product_id"))

	ratings, err := services.GetRatingsByProductID(db, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ratings)
}