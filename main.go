package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Paryz/rc-calculator-api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router.GET("/", hello)

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", pong)
		api.GET("/renderpdf", controllers.RenderPdfFromLatex)
	}

	router.Run(":" + port)
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
}
func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
