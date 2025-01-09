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

type Movie struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
	Genre string `json:"genre"`
	Date  string `json:"release_date"`
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

func GetMovieList(db *sql.DB) ([]Movie, error) {
	var movies []Movie
	query := `SELECT title, url,  release_date, genre FROM movies`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var movie Movie
		if err := rows.Scan(&movie.Title, &movie.URL, &movie.Date, &movie.Genre); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	// Check for any error after iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}
