package controllers

import (
	"grocery-purchase/services"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
)

func GetProducts(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	products, err := services.GetProducts(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func GetProductsByName(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	productName := c.Param("name")

	products, err := services.GetProductsByName(db, productName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(products) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No products found"})
		return
	}

	c.JSON(http.StatusOK, products)
}