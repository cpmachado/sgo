package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	Width  int
	Height int
	Mat    [][]uint8
}

// ColorModel returns the Image's color model.
func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.Width, i.Height)
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (i *Image) At(x int, y int) color.Color {
	if x < 0 || x >= i.Width || y < 0 || y >= i.Height {
		return color.Black
	}
	return color.Gray{Y: i.Mat[x][y]}
}

func NewImage(dx, dy int, f func(uint8, uint8) uint8) *Image {
	m := new(Image)
	res := make([][]uint8, dy)

	for y := range dy {
		res[y] = make([]uint8, dx)
		for x := range dx {
			res[y][x] = f(uint8(x), uint8(y))
		}
	}
	m.Height = dx
	m.Width = dy
	m.Mat = res

	return m
}

func main() {
	xorf := func(x, y uint8) uint8 { return x ^ y }
	m := NewImage(256, 256, xorf)
	pic.ShowImage(m)
}
