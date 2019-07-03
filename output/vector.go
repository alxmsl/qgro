package output

import (
	"io"

	"github.com/alxmsl/qrencode-go/qrencode"
)

type (
	VectorImage struct {
		Filename string
		Size     int
	}

	VectorOutput struct {
		block int
		x, y  int

		image     *VectorImage
		processor VectorProcessor
	}

	VectorProcessor func(*svg.SVG, int, int, int, uint8)
)

var (
	CircleProcessor = func(s *svg.SVG, x, y, block int, v uint8) {
		if v == qrencode.ElWhite {
			s.Circle(x+block/2, y+block/2, block/2, "fill:white;stroke:none")
		} else {
			s.Circle(x+block/2, y+block/2, block/2, "fill:black;stroke:none")
		}
	}

	RectangleProcessor = func(s *svg.SVG, x, y, block int, v uint8) {
		if v == qrencode.ElWhite {
			s.Rect(x, y, block, block, "fill:white;stroke:none")
		} else {
			s.Rect(x, y, block, block, "fill:black;stroke:none")
		}
	}
)

func NewVectorOutput(block, x, y int, processor VectorProcessor, image *VectorImage) *VectorOutput {
	return &VectorOutput{block, x, y, image, processor}
}

func (o *VectorOutput) Output(grid *qrencode.Grid, w io.Writer) {
	s := svg.New(w)

	width := (grid.Width() * o.block) + (o.block * 8)
	o.x = o.block * 4
	o.y = o.block * 4
	s.Start(width, width)

	currY := o.y
	for x := 0; x < grid.Width(); x++ {
		currX := o.x
		for y := 0; y < grid.Height(); y++ {
			o.processor(s, currX, currY, o.block, grid.Value(x, y))
			currX += o.block
		}
		currY += o.block
	}
	if o.image != nil {
		height := (grid.Height() * o.block) + (o.block * 8)
		s.Image((width-o.image.Size)/2, (height-o.image.Size)/2, o.image.Size, o.image.Size, o.image.Filename)
	}
	s.End()
}
