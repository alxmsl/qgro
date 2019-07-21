package output

import (
	"image"
	"image/color"

	"github.com/alxmsl/qrencode-go/qrencode"
)

type imageOutput struct {
	block, margin int
	black, white  color.Gray16
}

func (o *imageOutput) Image(grid *qrencode.Grid) image.Image {
	width := o.block * (2*o.margin + grid.Width())
	height := o.block * (2*o.margin + grid.Height())
	i := image.NewGray16(image.Rect(0, 0, width, height))
	for y := 0; y < o.block*o.margin; y++ {
		for x := 0; x < width; x++ {
			i.Set(x, y, o.white)
			i.Set(x, height-1-y, o.white)
		}
	}
	for y := o.block * o.margin; y < height-o.block*o.margin; y++ {
		for x := 0; x < o.block*o.margin; x++ {
			i.Set(x, y, o.white)
			i.Set(width-1-x, y, o.white)
		}
	}
	for y, w, h := 0, grid.Width(), grid.Height(); y < h; y++ {
		for x := 0; x < w; x++ {
			x0 := o.block * (x + o.margin)
			y0 := o.block * (y + o.margin)
			c := o.white
			if grid.IsDark(x, y) {
				c = o.black
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
