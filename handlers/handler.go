package handlers

import (
	"encoding/json"
	"log"
	"myproject/database"
	"myproject/models"
	"net/http"

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
