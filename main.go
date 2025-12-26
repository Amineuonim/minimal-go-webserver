package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"server/constants"
	"server/routes"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	constants.Db.ConnectPGDB()
	router := gin.Default()
	routes.RegisterRoutes(router)

	server := &http.Server{
		Addr:    constants.PORT,
		Handler: router,
	}

	// Start server in background
	go func() {
		log.Infof("Server listening on %s", constants.PORT)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server failed to start:", err)
		}
	}()

	// Listen for interrupt signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Warn("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(constants.TIMEOUT)*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error("Server forced to shutdown:", err)
	} else {
		log.Info("Server shut down cleanly")
	}
}
