package controllers

import (
	"net/http"

	"github.com/Paryz/rc-calculator-api/services"
	"github.com/gin-gonic/gin"
)

// RenderPdfFromLatex is self explanatory
func RenderPdfFromLatex(c *gin.Context) {
	pdf, err := services.RenderUsingGotex(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Type", "application/pdf")
	c.Writer.Write(pdf)
}
