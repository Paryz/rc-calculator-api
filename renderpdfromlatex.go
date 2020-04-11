package main

import "github.com/rwestlund/gotex"

func RenderUsingGotex() ([]byte, error) {

	var document = `
        \documentclass[12pt]{article}
        \begin{document}
        This is a LaTeX document.
        \end{document}
        `
	return gotex.Render(document, gotex.Options{
		Command:   "/Library/TeX/texbin/pdflatex",
		Runs:      1,
		Texinputs: "/my/asset/dir:/my/other/asset/dir"})
}
