package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		api.POST("/renderpdf", RenderPdfFromLatex)
	}

	router.Run(":3000")
}

// RenderPdfFromLatex is self explanatory
func RenderPdfFromLatex(c *gin.Context) {
	var beam RcBeam

	if err := c.ShouldBindJSON(&beam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var pdf, err = RenderUsingGotex()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gotex error"})
		return
	}
	c.Header("Content-Type", "application/pdf")
	c.Writer.Write(pdf)
}
