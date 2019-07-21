package output

import (
	"image/color"
	"image/png"
	"io"

	"github.com/alxmsl/qrencode-go/qrencode"
)

type PngOutput struct {
	imageOutput
}

func NewPngOutput(block, margin int, black, white color.Gray16) *PngOutput {
	return &PngOutput{
		imageOutput{
			block, margin,
			black, white,
		},
	}
}

func (o *PngOutput) Output(grid *qrencode.Grid, w io.Writer) {
	_ = png.Encode(w, o.Image(grid))
}
