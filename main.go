package main

import (

	"server/core"
	"server/core/user"

	"server/routes"

	"github.com/gin-gonic/gin"
	//log "github.com/sirupsen/logrus"
)

func main() {
    // 1. Initialize your DB manager
    dbManager := &core.DataBase{}
    dbManager.ConnectPGDB() // This populates dbManager.DB
    
    // 2. CRITICAL: Assign the connection to the user package's global DB
    user.DB = dbManager.DB 

    // 3. Now you can safely migrate and start the server
    user.DB.AutoMigrate(&user.User{})

    r := gin.Default()
    authHandler := &user.AuthHandler{}
    routes.RegisterRoutes(r, authHandler)

    r.Run(":8080")
}
