package controllers

import (
	"grocery-purchase/services"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
)

func GetStores(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)

	stores, err := services.GetStores(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch stores"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"stores": stores})
}
