package output

import (
	"image/color"
	"image/jpeg"
	"io"

	"github.com/alxmsl/qrencode-go/qrencode"
)

type JpegOutput struct {
	imageOutput
	options *jpeg.Options
}

func NewJpegOutput(block, margin int, black, white color.Gray16, options *jpeg.Options) *JpegOutput {
	return &JpegOutput{
		imageOutput{
			block, margin,
			black, white,
		},
		options,
	}
}

func (o *JpegOutput) Output(grid *qrencode.Grid, w io.Writer) {
	_ = jpeg.Encode(w, o.Image(grid), o.options)
}
