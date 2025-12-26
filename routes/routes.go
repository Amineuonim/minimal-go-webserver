package routes

import (
	"server/handlers"
	"github.com/gin-gonic/gin"

	"time"
	"github.com/gin-contrib/timeout"
	
	"server/constants"


)

var timer = timeout.New(timeout.WithTimeout(constants.REQUESTTIMEOUTMS *time.Millisecond))

func RegisterRoutes(r *gin.Engine) {
	// basic routes
	//r.GET("/", handlers.Home)
	//r.GET("/health", handlers.Health)

	// group example (like Django namespaces)
	api := r.Group("/api")
	{
		api.GET("/health", timer, handlers.Health)
		api.GET("/testdb",timer,handlers.TestDb)
	}
}