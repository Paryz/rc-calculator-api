package controllers

import (
	"net/http"

	"github.com/Paryz/rc-calculator-api/services"
	"github.com/Paryz/rc-calculator-api/types"
	"github.com/gin-gonic/gin"
)

// RenderPdfFromLatex is self explanatory
func RenderPdfFromLatex(c *gin.Context) {
	var beam types.RcBeam

	if err := c.ShouldBindQuery(&beam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pdf, err := services.RenderUsingGotex(beam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gotex error"})
		return
	}
	c.Header("Content-Type", "application/pdf")
	c.Writer.Write(pdf)
}
