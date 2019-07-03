package output

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/alxmsl/qrencode-go/qrencode"
)

const (
	Gif = iota
	Jpeg
	Png
)

type Interface interface {
	Output(*qrencode.Grid, io.Writer)
}

type ImageOutput struct {
	format, block, margin int

	gif  *gif.Options
	jpeg *jpeg.Options
}

func NewImageOutput(format, block, margin int, gif *gif.Options, jpeg *jpeg.Options) *ImageOutput {
	switch format {
	case Gif:
		fallthrough
	case Jpeg:
		fallthrough
	case Png:
		return &ImageOutput{format, block, margin, gif, jpeg}
	default:
		panic(fmt.Sprintf("invalid format: %d", format))
	}
}

func (o *ImageOutput) Output(grid *qrencode.Grid, w io.Writer) {
	switch o.format {
	case Gif:
		_ = gif.Encode(w, o.Image(grid), o.gif)
	case Jpeg:
		_ = jpeg.Encode(w, o.Image(grid), o.jpeg)
	case Png:
		_ = png.Encode(w, o.Image(grid))
	}
}

func (o *ImageOutput) Image(grid *qrencode.Grid) image.Image {
	width := o.block * (2*o.margin + grid.Width())
	height := o.block * (2*o.margin + grid.Height())
	i := image.NewGray16(image.Rect(0, 0, width, height))
	for y := 0; y < o.block*o.margin; y++ {
		for x := 0; x < width; x++ {
			i.Set(x, y, color.White)
			i.Set(x, height-1-y, color.White)
		}
	}
	for y := o.block * o.margin; y < height-o.block*o.margin; y++ {
		for x := 0; x < o.block*o.margin; x++ {
			i.Set(x, y, color.White)
			i.Set(width-1-x, y, color.White)
		}
	}
	for y, w, h := 0, grid.Width(), grid.Height(); y < h; y++ {
		for x := 0; x < w; x++ {
			x0 := o.block * (x + o.margin)
			y0 := o.block * (y + o.margin)
			c := color.White
			if grid.Get(x, y) {
				c = color.Black
			}
			for dy := 0; dy < o.block; dy++ {
				for dx := 0; dx < o.block; dx++ {
					i.Set(x0+dx, y0+dy, c)
				}
			}
		}
	}
	return i
}
