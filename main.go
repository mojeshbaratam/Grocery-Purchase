package main

import (
	"log"
	"grocery-purchase/config"
	"grocery-purchase/controllers"
	"grocery-purchase/server"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize the database connection
	db, err := server.InitDB(&cfg) // Pass the pointer to cfg
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize the HTTP server
	router := gin.Default()

	// Pass DB connection to the handlers via middleware
	router.Use(func(c *gin.Context) {
		c.Set("db", db) // store db instance in Gin context
		c.Next()
	})

	// Setup routes
	router.POST("/user/login", controllers.UserLogin)
	router.POST("/admin/login", controllers.AdminLogin)
	router.POST("/admin/store", controllers.CreateStore)
	router.POST("/admin/product", controllers.CreateProduct)
	router.GET("/admin/stores", controllers.GetStores)
	// router.GET("/user/product", controllers.GetProducts)

		// Setup routes
	router.POST("/user/register", controllers.RegisterUser) // Add this line


	// Run the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
