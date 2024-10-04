package controllers

import (
	"grocery-purchase/models"
	"grocery-purchase/services"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
)

func AdminLogin(c *gin.Context) {
	var loginData models.Admin
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	db := c.MustGet("db").(*sql.DB)

	authenticated, err := services.AuthenticateAdmin(db, loginData.Username, loginData.Password)
	if err != nil || !authenticated {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Admin login successful"})
}
