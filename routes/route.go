package routes

import (
	"myproject/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// User routes
	userGroup := router.Group("/users")
	{
		userGroup.POST("/", handlers.CreateUser) // Endpoint to create a user
		userGroup.GET("/", handlers.GetUsers)    // Endpoint to fetch all users
	}
}
