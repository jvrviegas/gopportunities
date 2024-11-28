package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize the router using Gin's default configuration
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Run the server
	router.Run()
}
