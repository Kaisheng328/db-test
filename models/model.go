package models

import (
	"myproject/database"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func CreateUser(name, email string) error {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err := database.DB.Exec(query, name, email)
	return err
}

func GetAllUsers() ([]User, error) {
	query := "SELECT id, name, email, created_at FROM users"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
