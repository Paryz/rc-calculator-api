package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rwestlund/gotex"
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
	var document = `
        \documentclass[12pt]{article}
        \begin{document}
        This is a LaTeX document.
        \end{document}
        `
	var pdf, err = gotex.Render(document, gotex.Options{
		Command:   "/Library/TeX/texbin/pdflatex",
		Runs:      1,
		Texinputs: "/my/asset/dir:/my/other/asset/dir"})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Type", "application/pdf")
	c.Writer.Write(pdf)
}

// RcBeam type is a beam representation
type RcBeam struct {
	Height                   string `json:"height" binding:"required"`
	Width                    string `json:"width" binding:"required"`
	Cover                    string `json:"cover" binding:"required"`
	TopCover                 string `json:"topCover" binding:"required"`
	ConcreteClass            string `json:"concreteClass" binding:"required"`
	SteelClass               string `json:"steelClass" binding:"required"`
	ConcreteFactor           string `json:"concreteFactor" binding:"required"`
	SteelFactor              string `json:"steelFactor" binding:"required"`
	AlphaCC                  string `json:"alphaCC" binding:"required"`
	LinkDiameter             string `json:"linkDiameter" binding:"required"`
	MainBarDiameter          string `json:"mainBarDiameter" binding:"required"`
	BendingMoment            string `json:"bendingMoment" binding:"required"`
	Fcd                      string `json:"fcd" binding:"required"`
	Fctm                     string `json:"fctm" binding:"required"`
	Fyd                      string `json:"fyd" binding:"required"`
	EffectiveHeight          string `json:"effectiveHeight" binding:"required"`
	MinReinforcement         string `json:"minReinforcement" binding:"required"`
	Sc                       string `json:"sC" binding:"required"`
	KsiEffective             string `json:"ksiEffective" binding:"required"`
	KsiEffectiveLim          string `json:"ksiEffectiveLim" binding:"required"`
	BottomReinforcementPrime string `json:"bottomReinforcementPrime" binding:"required"`
	BendingMomentPrime       string `json:"bendingMomentPrime" binding:"required"`
	BendingMomentDelta       string `json:"bendingMomentDelta" binding:"required"`
	TopReqReinforcement      string `json:"topReqReinforcement" binding:"required"`
	BottomReqReinforcement   string `json:"bottomReqReinforcement" binding:"required"`
	MaximumReinforcement     string `json:"maximumReinforcement" binding:"required"`
}
