package models

import (
	"database/sql"
	"errors"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	LoginID   string `json:"login_id"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}

func CreateUser(db *sql.DB, name, email, loginID, hashedPassword string) error {
	query := `INSERT INTO users (name, email, login_id, password) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, name, email, loginID, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}

// GetUserPassword retrieves the hashed password for a given login ID
func GetUserPassword(db *sql.DB, loginID string) (string, error) {
	var hashedPassword string
	query := `SELECT password FROM users WHERE login_id = ?`
	err := db.QueryRow(query, loginID).Scan(&hashedPassword)
	if err == sql.ErrNoRows {
		return "", errors.New("user not found")
	}
	if err != nil {
		return "", err
	}
	return hashedPassword, nil
}
