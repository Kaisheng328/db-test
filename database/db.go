package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDatabase() error {
	var err error
	DB, err = sql.Open("sqlite3", "app.db")
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	log.Println("Database connected successfully")
	return nil
}

func CreateUsersTable() error {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        email TEXT UNIQUE NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`
	_, err := DB.Exec(query)
	if err != nil {
		return err
	}

	log.Println("Users table created successfully")
	return nil
}
func CreateMovieTable() error {
	query := `
    CREATE TABLE IF NOT EXISTS movies (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    url TEXT NOT NULL
);`
	_, err := DB.Exec(query)
	if err != nil {
		return err
	}

	log.Println("Movie table created successfully")
	return nil
}
