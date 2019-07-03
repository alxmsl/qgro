package qrencode

import (
	"bytes"
)

type Grid struct {
	width, height int
	bits          []uint8
}

func NewGrid(width, height int) *Grid {
	return &Grid{width, height, make([]uint8, 2*width*height)}
}

func (g *Grid) String() string {
	b := bytes.Buffer{}
	for y, w, h := 0, g.Width(), g.Height(); y < h; y++ {
		for x := 0; x < w; x++ {
			if g.Empty(x, y) {
				b.WriteString(" ")
			} else if g.Get(x, y) {
				b.WriteString("#")
			} else {
				b.WriteString("_")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (g *Grid) Width() int {
	return g.width
}

func (g *Grid) Height() int {
	return g.height
}

func (g *Grid) Empty(x, y int) bool {
	return g.bits[2*(x+y*g.width)] == ElWhite
}

func (g *Grid) Get(x, y int) bool {
	return g.Value(x, y) > ElWhite
}

func (g *Grid) Value(x, y int) uint8 {
	return g.bits[2*(x+y*g.width)+1]
}

func (g *Grid) Set(x, y int, v bool) {
	if v {
		g.SetValue(x, y, ElBlack)
	} else {
		g.SetValue(x, y, ElWhite)
	}
}

func (g *Grid) SetValue(x, y int, v uint8) {
	g.bits[2*(x+y*g.width)] = 0x1
	g.bits[2*(x+y*g.width)+1] = v
}

func (g *Grid) Clear() {
	for i, _ := range g.bits {
		g.bits[i] = ElWhite
	}
}
