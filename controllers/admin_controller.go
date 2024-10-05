package controllers

import (
	"grocery-purchase/models"
	"grocery-purchase/services"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	"log"
)

func RegisterAdmin(c *gin.Context) {
	log.Println("Received request to create admin:")
	var admin models.Admin

	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// You may want to hash the password here before saving
	if err := services.CreateAdmin(c.MustGet("db").(*sql.DB), &admin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register admin"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "admin registered successfully"})
}

func CreateStore(c *gin.Context) {
	var store models.Store
	if err := c.ShouldBindJSON(&store); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	db := c.MustGet("db").(*sql.DB)

	if err := services.CreateStore(db, &store); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create store"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Store created successfully", "store": store})
}

func UpdateStore(c *gin.Context) {
	// Store the current store name from URL parameter
	storeName := c.Param("name")
	if storeName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store name"})
		return
	}

	var storeDetails struct {
		NewName     string `json:"new_name"`
		NewLocation string `json:"new_location"`
	}

	if err := c.ShouldBindJSON(&storeDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	db := c.MustGet("db").(*sql.DB)

	// Call the service to update the store
	updated, err := services.UpdateStore(db, storeName, storeDetails.NewName, storeDetails.NewLocation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update store"})
		return
	}

	if !updated {
		c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Store updated successfully"})
}

func RemoveStore(c *gin.Context) {
	storeName := c.Param("name")
	if storeName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store name"})
		return
	}

	db := c.MustGet("db").(*sql.DB)

	deleted, err := services.RemoveStore(db, storeName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete store"})
		return
	}

	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Store deleted successfully"})
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	db := c.MustGet("db").(*sql.DB)

	if err := services.CreateProduct(db, &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully"})
}

func UpdateProduct(c *gin.Context) {
	// Store the current store name from URL parameter
	productName := c.Param("name")
	if productName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Product name"})
		return
	}

	var productDetails struct {
		NewName    string  `json:"new_name"`
		NewPrice   float64 `json:"new_price"`
		NewStoreID int     `json:"new_store_id"`
	}

	if err := c.ShouldBindJSON(&productDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	db := c.MustGet("db").(*sql.DB)

	// Call the service to update the store
	updated, err := services.UpdateProduct(db, productName, productDetails.NewName, productDetails.NewPrice, productDetails.NewStoreID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Product"})
		return
	}

	if !updated {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func RemoveProduct(c *gin.Context) {
	productName := c.Param("name")
	if productName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product name"})
		return
	}

	db := c.MustGet("db").(*sql.DB)

	deleted, err := services.RemoveProduct(db, productName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}