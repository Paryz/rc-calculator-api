package templates

import (
	"fmt"
	"github.com/Paryz/rc-calculator-api/types"
)

// GenerateTemplate generates template for calculations of one row reinforcement
func GenerateTemplate(beam types.RcBeam) string {
	return fmt.Sprint(`
		\documentclass[12pt]{article}
		\usepackage{amsmath}
		\newcommand\fixedalign{\hspace{0.3\linewidth} &
                       \hspace{0.7\linewidth} \nonumber \\[-\baselineskip]}
		\begin{document}
		`,
		sectionDetails(beam),
		concreteDetails(beam),
		steelDetails(beam),
		minimumReinforcement(beam),
		requredReinforcement(beam),
	)
}

func sectionDetails(beam types.RcBeam) string {
	return fmt.Sprintf(`
		\section{RC Section Details}
		\begin{align*}\fixedalign
		h &= %.1fmm\\
		b &= %.1fmm \\
		c_{nom} &= %.1fmm \\
		\phi &= %.1fmm \\
		\phi_{S} &= %.1fmm \\
		M_{Ed} &= %.1fkNm
		\end{align*}
		`,
		beam.Height,
		beam.Width,
		beam.Cover,
		beam.MainBarDiameter,
		beam.LinkDiameter,
		beam.BendingMoment,
	)
}

func concreteDetails(beam types.RcBeam) string {
	return fmt.Sprintf(`
		\section{Concrete Details}
		\begin{align*}\fixedalign
		f_{ck} &= %.2fMPa \\
		\gamma_C &= %.2f \\
		\alpha_{cc} &= %.1f \\
		f_{cd} &= %.2fMPa \\
		f_{ctm} &= %.2fMPa
		\end{align*}
		`,
		beam.ConcreteClass,
		beam.ConcreteFactor,
		beam.AlphaCC,
		beam.Fcd,
		beam.Fctm,
	)
}

func steelDetails(beam types.RcBeam) string {
	return fmt.Sprintf(`
		\section{Steel Details}
		\begin{align*}\fixedalign
		f_{yk} &= %.1fMPa \\
		\gamma_S &= %.2f \\
		f_{yd} &= %.2fMPa
		\end{align*}
	`,
		beam.SteelClass,
		beam.SteelFactor,
		beam.Fyd,
	)
}

func minimumReinforcement(beam types.RcBeam) string {
	return fmt.Sprintf(`
		\section{Minimum Reinforcement}
		\begin{align*}\fixedalign
		d &= %.1fmm \\
		A_{s,min,1} &= %.1fmm^2 \\
		A_{s,min,2} &= %.1fmm^2 \\
		A_{s,min} &= %.1fmm^2
		\end{align*}
	`,
		beam.EffectiveHeight,
		minReinforcement1(beam),
		minReinforcement2(beam),
		beam.MinReinforcement,
	)
}

func requredReinforcement(beam types.RcBeam) string {
	return fmt.Sprintf(`
		\section{Required Reinforcement}
		\begin{align*}\fixedalign
		\varepsilon_{cu2} &= 0.0035 \\
		\varepsilon_{cu2} &= 0.0035 \\
		E_{S} &= 200000 MPa \\
		S_{C} &= %.1f \\
		\xi_{eff} &= %.2f \\
		\xi_{eff,lim} &= %.2f \\
		A_{s,b} &= %.1fmm^2 \\
		A_{s,b} &> A_{s,min} \\
		A_{s,req} &= %.1fmm^2
		\end{align*}
		\end{document}
	`,
		beam.Sc,
		beam.KsiEffective,
		beam.KsiEffectiveLim,
		*beam.BottomReqReinforcement,
		*beam.BottomReqReinforcement,
	)
}

func minReinforcement1(beam types.RcBeam) float64 {
	fctm := beam.Fctm
	steelClass := beam.SteelClass
	width := beam.Width
	effectiveHeight := beam.EffectiveHeight
	return 0.26 * (fctm / steelClass) * width * effectiveHeight
}

func minReinforcement2(beam types.RcBeam) float64 {
	width := beam.Width
	effectiveHeight := beam.EffectiveHeight
	return 0.0013 * width * effectiveHeight
}
