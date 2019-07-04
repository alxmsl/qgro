package qrencode

import "bytes"

type cell struct {
	bool
	uint8
}

type Grid struct {
	width, height int
	cells         []cell
}

func NewGrid(width, height int) *Grid {
	return &Grid{width, height, make([]cell, width*height)}
}

func (g *Grid) Get(x, y int) uint8 {
	return g.cells[x+y*g.width].uint8
}

func (g *Grid) Set(x, y int, v uint8) {
	g.cells[x+y*g.width].bool = true
	g.cells[x+y*g.width].uint8 = v
}

func (g *Grid) Clear() {
	for i := range g.cells {
		g.cells[i].bool = false
		g.cells[i].uint8 = 0x0
	}
}

func (g *Grid) SetDark(x, y int, v bool) {
	if v {
		g.Set(x, y, ElBlack)
	} else {
		g.Set(x, y, ElWhite)
	}
}

func (g *Grid) IsDark(x, y int) bool {
	return g.Get(x, y) > ElWhite
}

func (g *Grid) IsEmpty(x, y int) bool {
	return !g.cells[x+y*g.width].bool
}

func (g *Grid) Height() int {
	return g.height
}

func (g *Grid) Width() int {
	return g.width
}

func (g *Grid) String() string {
	b := bytes.Buffer{}
	for y, w, h := 0, g.Width(), g.Height(); y < h; y++ {
		for x := 0; x < w; x++ {
			if g.IsEmpty(x, y) {
				b.WriteString(" ")
			} else if g.IsDark(x, y) {
				b.WriteString("#")
			} else {
				b.WriteString("_")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}
