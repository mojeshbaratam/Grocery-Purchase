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
	router.POST("/admin/login", controllers.AdminLogin)
	router.POST("/admin/register", controllers.RegisterAdmin)
	
	router.POST("/admin/store", controllers.CreateStore)
	router.GET("/admin/stores", controllers.GetStores)
	router.PUT("/admin/store/name/:name", controllers.UpdateStore)
	router.DELETE("/admin/store/name/:name", controllers.RemoveStore) 

	router.POST("/admin/product", controllers.CreateProduct)
	router.GET("/products", controllers.GetProducts)
	router.GET("/products/name/:name", controllers.GetProductsByName)
	router.PUT("/admin/product/name/:name", controllers.UpdateProduct)
	router.DELETE("/admin/product/name/:name", controllers.RemoveProduct) 

	router.POST("/ratings", controllers.CreateRating)
	router.GET("/products/:product_id/ratings", controllers.GetRatings)

	router.POST("/user/register", controllers.RegisterUser) 
	router.POST("/user/login", controllers.UserLogin)
	router.POST("/user/cart/add", controllers.AddProductToCart)   // Add product to cart
	router.GET("/user/cart/view", controllers.ViewCart)           // View the cart
	router.POST("/user/cart/purchase", controllers.PurchaseCart)  // Purchase items in the cart
	

	// Run the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
