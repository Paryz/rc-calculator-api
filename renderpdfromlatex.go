package main

import (
	"github.com/rwestlund/gotex"
	"os"
)

func RenderUsingGotex() ([]byte, error) {

	url := os.Getenv("TEXLIVE_ADDRESS")
	var document = `
        \documentclass[12pt]{article}
        \begin{document}
        This is a LaTeX document.
        \end{document}
        `
	return gotex.Render(document, gotex.Options{
		Command:   url,
		Runs:      1,
		Texinputs: "/my/asset/dir:/my/other/asset/dir"})
}
