package main

import (
	"log"
	"myproject/database"
	"myproject/routes"
	"net/http"
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

	routes.RegisterRoutes()

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
