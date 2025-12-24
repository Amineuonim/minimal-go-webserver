package main

import (
	//"server/core"
	"server/routes"
	"server/constants"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

func main(){

	router := gin.Default()
	routes.RegisterRoutes(router)

	// Start the HTTPS server
	if err := router.Run(constants.PORT); err != nil {
		log.Fatal("Oh Oh! looks like i couldn't start :(\ncauses:\n", err)
	}


}
