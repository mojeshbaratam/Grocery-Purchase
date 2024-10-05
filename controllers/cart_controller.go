package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"grocery-purchase/models"
	"grocery-purchase/services"
)

// AddProductToCart adds a product to the user's cart.
func AddProductToCart(c *gin.Context) {
	var cart models.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*sql.DB)
	err := services.AddProductToCartService(db, &cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product to cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added to cart successfully"})
}

// ViewCart retrieves the cart contents for a specific user.
func ViewCart(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil || userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	db := c.MustGet("db").(*sql.DB)
	cart, err := services.GetCartService(db, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cart": cart})
}

// PurchaseCart handles purchasing all items in the cart.
func PurchaseCart(c *gin.Context) {
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil || userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	db := c.MustGet("db").(*sql.DB)

	// Clear the cart after purchase
	err = services.ClearCartService(db, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process purchase"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Purchase successful"})
}
