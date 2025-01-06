package main

import (
	"log"
	"myproject/database"
	"myproject/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	if err := database.InitDatabase(); err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer database.DB.Close()

	// Run migrations
	if err := database.CreateUsersTable(); err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}

	// Set up Gin router
	router := gin.Default()
	routes.RegisterRoutes(router)

	// Start the server
	log.Println("Server running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
