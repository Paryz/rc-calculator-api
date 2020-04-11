package main

import (
	"fmt"
	"github.com/rwestlund/gotex"
	"os"
)

func RenderUsingGotex() ([]byte, error) {

	url := os.Getenv("TEXLIVE_ADDRESS")
	test := "DUPA"
	var document = fmt.Sprintf(`
        \documentclass[12pt]{article}
        \begin{document}
        This is a %s document.
        \end{document}
        `, test)
	return gotex.Render(document, gotex.Options{
		Command:   url,
		Runs:      1,
		Texinputs: "/my/asset/dir:/my/other/asset/dir"})
}
