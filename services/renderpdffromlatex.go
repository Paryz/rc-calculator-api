package services

import (
	"os"

	"github.com/Paryz/rc-calculator-api/templates"
	"github.com/Paryz/rc-calculator-api/types"
	"github.com/gin-gonic/gin"
	"github.com/rwestlund/gotex"
)

func RenderUsingGotex(c *gin.Context) ([]byte, error) {
	var beam types.RcBeam

	if err := c.ShouldBindQuery(&beam); err != nil {
		return []byte(""), err
	}

	document := templates.FetchTemplate(beam)

	url := os.Getenv("TEXLIVE_ADDRESS")

	options := gotex.Options{
		Command:   url,
		Runs:      1,
		Texinputs: "/my/asset/dir:/my/other/asset/dir"}

	return gotex.Render(document, options)
}
