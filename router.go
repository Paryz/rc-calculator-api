package main

import (
	"github.com/Paryz/rc-calculator-api/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// setupRoutes sets all routes so main function is clean
func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", hello)

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", pong)
		api.GET("/renderpdf", controllers.RenderPdfFromLatex)
	}
	return router
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
}
func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
