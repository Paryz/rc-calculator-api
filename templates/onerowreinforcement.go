package templates

import (
	"fmt"
	"github.com/Paryz/rc-calculator-api/types"
)

func FetchTemplate(beam types.RcBeam) string {

	return fmt.Sprintf(`
        \documentclass[12pt]{article}
        \begin{document}
        This is a %s document.
        \end{document}
        `, beam.Height)
}
