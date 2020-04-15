package templates

import (
	"fmt"
	"github.com/Paryz/rc-calculator-api/types"
)

// GenerateTemplate generates template for calculations of one row reinforcement
func GenerateTemplate(beam types.RcBeam) string {
	return fmt.Sprintf(`
		\documentclass[12pt]{article}
		\usepackage{amsmath}
		\newcommand\fixedalign{\hspace{0.3\linewidth} &
                       \hspace{0.7\linewidth} \nonumber \\[-\baselineskip]}
		\begin{document}
		\section{RC Section Details}
		\begin{align*}\fixedalign
		h &= %smm\\
		b &= %smm \\
		c_{nom} &= 30mm \\
		\phi &= 25mm \\
		\phi_{S} &= 10mm \\
		M_{Ed} &= 200kNm
		\end{align*}
		\section{Concrete Details}
		\begin{align*}\fixedalign
		f_{ck} &= 30MPa \\
		\gamma_C &= 1.4 \\
		\alpha_{cc} &= 1.0 \\
		f_{cd} &= 17.46MPa \\
		f_{ctm} &= 1.7MPa
		\end{align*}
		\section{Steel Details}
		\begin{align*}\fixedalign
		f_{yk} &= 500MPa \\
		\gamma_S &= 1.15 \\
		f_{yd} &= 434.70MPa
		\end{align*}
		\section{Minimum Reinforcement}
		\begin{align*}\fixedalign
		d &= 550mm \\
		A_{s,min,1} &= 100mm^2 \\
		A_{s,min,2} &= 150mm^2 \\
		A_{s,min} &= 100mm^2
		\end{align*}
		\section{Required Reinforcement}
		\begin{align*}\fixedalign
		\varepsilon_{cu2} &= 0.0035 \\
		\varepsilon_{cu2} &= 0.0035 \\
		E_{S} &= 200000 MPa \\
		S_{C} &= 1.7 \\
		\xi_{eff} &= 1.8 \\
		\xi_{eff,lim} &= 0.90 \\
		A_{s,b} &= 1000mm^2 \\
		A_{s,b} &> A_{s,min} \\
		A_{s,req} &= 1500mm^2
		\end{align*}
		\end{document}
		`, beam.Height, beam.Width)
}
