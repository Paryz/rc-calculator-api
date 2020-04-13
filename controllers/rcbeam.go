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

	if err := c.ShouldBindJSON(&beam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var pdf, err = services.RenderUsingGotex("")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gotex error"})
		return
	}
	c.Header("Content-Type", "application/pdf")
	c.Writer.Write(pdf)
}

// RenderPdf is self explanatory
func RenderPdf(c *gin.Context) {
	var pdf []byte
	var err error
	if test := c.Param("test"); test == "" {
		pdf, err = services.RenderUsingGotex("")
	} else {
		pdf, err = services.RenderUsingGotex(test)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gotex error"})
		return
	}
	c.Header("Content-Type", "application/pdf")
	c.Writer.Write(pdf)
}
