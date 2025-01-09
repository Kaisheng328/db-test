package handlers

import (
	"encoding/json"
	"log"
	"myproject/database"
	"myproject/models"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	LoginID  string `json:"login_id"`
	Password string `json:"password"`
}

type LoginRequest struct {
	LoginID  string `json:"login_id"`
	Password string `json:"password"`
}
type MovieRequest struct {
	Title string `json:"title"` // Optional: if you want to fetch a specific movie by title
	URL   string `json:"url"`   // Optional: for filtering by URL
	Genre string `json:"genre"`
	Date  string `json:"release_date"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var req SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error creating account", http.StatusInternalServerError)
		return
	}

	// Save the user to the database
	err = models.CreateUser(database.DB, req.Name, req.Email, req.LoginID, string(hashedPassword))
	if err != nil {
		http.Error(w, "User already exists or database error", http.StatusConflict)
		log.Printf("Error creating user: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Get the hashed password for the given login ID
	hashedPassword, err := models.GetUserPassword(database.DB, req.LoginID)
	if err != nil {
		http.Error(w, "Invalid login ID or password", http.StatusUnauthorized)
		log.Printf("Login error: %v", err)
		return
	}

	// Compare the hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)); err != nil {
		http.Error(w, "Invalid login ID or password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}
func MovieHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the request body to extract the query parameters (title and/or URL)
	var req MovieRequest

	// Call the GetMovieList function to fetch the list of movies from the database
	movies, err := models.GetMovieList(database.DB)
	if err != nil {
		http.Error(w, "Failed to retrieve movies", http.StatusInternalServerError)
		log.Printf("Error fetching movies: %v", err)
		return
	}

	// Filter the movie list by title if the title is provided
	if req.Title != "" {
		var filteredMovies []models.Movie
		for _, movie := range movies {
			if containsIgnoreCase(movie.Title, req.Title) {
				filteredMovies = append(filteredMovies, movie)
			}
		}
		movies = filteredMovies
	}

	// Filter the movie list by URL if the URL is provided
	if req.URL != "" {
		var filteredMovies []models.Movie
		for _, movie := range movies {
			if containsIgnoreCase(movie.URL, req.URL) {
				filteredMovies = append(filteredMovies, movie)
			}
		}
		movies = filteredMovies
	}
	if req.Date != "" {
		var filteredMovies []models.Movie
		for _, movie := range movies {
			if containsIgnoreCase(movie.Date, req.Date) {
				filteredMovies = append(filteredMovies, movie)
			}
		}
		movies = filteredMovies
	}

	if req.Genre != "" {
		var filteredMovies []models.Movie
		for _, movie := range movies {
			if containsIgnoreCase(movie.Genre, req.Genre) {
				filteredMovies = append(filteredMovies, movie)
			}
		}
		movies = filteredMovies
	}
	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Send the movies data as a JSON response
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		http.Error(w, "Failed to encode movies data", http.StatusInternalServerError)
		log.Printf("Error encoding movies response: %v", err)
	}
}

func containsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
