package style

import "github.com/muesli/termenv"

var (
	color   = termenv.EnvColorProfile().Color
	Keyword = termenv.Style{}.Foreground(color("204")).Background(color("235")).Styled
	Help    = termenv.Style{}.Foreground(color("241")).Styled
)
