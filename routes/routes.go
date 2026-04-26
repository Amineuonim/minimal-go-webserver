package routes

import (
	"server/constants"
	"server/handlers"
	"server/core/user" // Import your user package
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

var timer = timeout.New(timeout.WithTimeout(constants.REQUESTTIMEOUTMS * time.Millisecond))

// We pass the authHandler here so routes can access the DB logic
func RegisterRoutes(r *gin.Engine, authHandler *user.AuthHandler) {
	
	// Public Auth Routes
	auth := r.Group("/auth")
	{
		auth.POST("/signup", timer, authHandler.Signup)
		auth.POST("/login", timer, authHandler.Login)
	}

	// Protected API Group
	api := r.Group("/api")
//	api.Use(user.AuthMiddleware()) // Protect all routes in this group
	{
		api.GET("/health", timer, handlers.Health)
		api.GET("/testdb", timer, handlers.TestDb)
		
		// Example: Get current user profile
		// api.GET("/me", timer, authHandler.GetProfile) 
	}
}