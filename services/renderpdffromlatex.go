package services

import (
	"os"

	"github.com/Paryz/rc-calculator-api/templates"
	"github.com/Paryz/rc-calculator-api/types"
	"github.com/rwestlund/gotex"
)

func RenderUsingGotex(beam types.RcBeam) ([]byte, error) {

	url := os.Getenv("TEXLIVE_ADDRESS")
	document := templates.FetchTemplate(beam)
	return gotex.Render(document, gotex.Options{
		Command:   url,
		Runs:      1,
		Texinputs: "/my/asset/dir:/my/other/asset/dir"})
}
