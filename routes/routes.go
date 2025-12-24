package routes

import (
	"server/endpoints"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// basic routes
	//r.GET("/", handlers.Home)
	//r.GET("/health", handlers.Health)

	// group example (like Django namespaces)
	api := r.Group("/api")
	{
		api.GET("/ping", endpoints.Ping)
	}
}