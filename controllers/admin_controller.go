package controllers

import (
	"grocery-purchase/models"
	"grocery-purchase/services"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
)

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
