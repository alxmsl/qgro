package output

import (
	"fmt"
	"io"

	"github.com/alxmsl/qrencode-go/qrencode"
)

const (
	TermBlack = iota + 40
	TermRed
	TermGreen
	TermYellow
	TermBlue
	TermMagenta
	TermCyan
	TermWhite
)

const (
	TermBlackBright = iota + 100
	TermRedBright
	TermGreenBright
	TermYellowBright
	TermBlueBright
	TermMagentaBright
	TermCyanBright
	TermWhiteBright
)

type TerminalSymbol struct {
	Background int
	Symbol     string
}

type TerminalOutput struct {
	black, white *TerminalSymbol
}

func NewTerminalOutput(black, white *TerminalSymbol) *TerminalOutput {
	return &TerminalOutput{black, white}
}

func (o *TerminalOutput) Output(grid *qrencode.Grid, w io.Writer) {
	white := fmt.Sprintf("\033[%dm%s\033[0m", o.white.Background, o.white.Symbol)
	black := fmt.Sprintf("\033[%dm%s\033[0m", o.black.Background, o.black.Symbol)
	newline := "\n"

	write(w, []byte(white))
	for i := 0; i <= grid.Width(); i++ {
		_, _ = w.Write([]byte(white))
	}
	write(w, []byte(newline))

	for i := 0; i < grid.Height(); i++ {
		write(w, []byte(white))
		for j := 0; j < grid.Width(); j++ {
			if grid.Get(j, i) {
				write(w, []byte(black))
			} else {
				write(w, []byte(white))
			}
		}
		write(w, []byte(white))
		write(w, []byte(newline))
	}
	write(w, []byte(white))
	for i := 0; i <= grid.Width(); i++ {
		write(w, []byte(white))
	}
	write(w, []byte(newline))
}

func write(w io.Writer, p []byte) {
	_, _ = w.Write(p)
}
