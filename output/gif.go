package output

import (
	"image/color"
	"image/gif"
	"io"

	"github.com/alxmsl/qrencode-go/qrencode"
)

type GifOutput struct {
	imageOutput
	options *gif.Options
}

func NewGifOutput(block, margin int, black, white color.Gray16, options *gif.Options) *GifOutput {
	return &GifOutput{
		imageOutput{
			block, margin,
			black, white,
		},
		options,
	}
}

func (o *GifOutput) Output(grid *qrencode.Grid, w io.Writer) {
	_ = gif.Encode(w, o.Image(grid), o.options)
}
