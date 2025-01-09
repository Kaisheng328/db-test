package routes

import (
	"myproject/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/signup", handlers.SignupHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/movie", handlers.MovieHandler)
}
