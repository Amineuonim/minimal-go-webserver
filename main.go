package main

import (
	"log"

	"server/constants"
	"server/core"
	"server/core/user"
	"server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	dbManager := &core.DataBase{}
	dbManager.ConnectPGDB()

	user.DB = dbManager.DB
	constants.Db.DB = dbManager.DB

	if err := user.EnsureSchema(dbManager.DB); err != nil {
		log.Fatalf("failed to ensure database schema: %v", err)
	}

	r := gin.Default()
	authHandler := &user.AuthHandler{}
	routes.RegisterRoutes(r, authHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
