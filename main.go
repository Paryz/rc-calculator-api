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

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		api.GET("/dupa", controllers.RenderPdf)
		api.GET("/dupa/:test", controllers.RenderPdf)
		api.POST("/renderpdf", controllers.RenderPdfFromLatex)
	}

	router.Run(":" + port)
}
