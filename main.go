package main

import (
	"log"
	"myproject/database"
	"myproject/routes"
	"net/http"
)

// becoming trailer website and portfolio firstt
func main() {
	// Initialize the database
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	if err := database.InitDatabase(); err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer database.DB.Close()

	// Run migrations
	if err := database.CreateUsersTable(); err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}
	if err := database.CreateMovieTable(); err != nil {
		log.Fatalf("Failed to create Movie table: %v", err)
	}
	routes.RegisterRoutes()

	log.Println("Server running on http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
