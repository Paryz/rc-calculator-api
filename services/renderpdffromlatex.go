package services

import (
	"fmt"
	"os"

	"github.com/rwestlund/gotex"
)

func RenderUsingGotex(param string) ([]byte, error) {

	url := os.Getenv("TEXLIVE_ADDRESS")
	if param == "" {
		param = "DUPA"
	}
	var document = fmt.Sprintf(`
        \documentclass[12pt]{article}
        \begin{document}
        This is a %s document.
        \end{document}
        `, param)
	return gotex.Render(document, gotex.Options{
		Command:   url,
		Runs:      1,
		Texinputs: "/my/asset/dir:/my/other/asset/dir"})
}
