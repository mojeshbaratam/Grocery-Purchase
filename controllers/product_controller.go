package controllers

import (
	"grocery-purchase/models"
	"grocery-purchase/services"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
)

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

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully", "product": product})
}
