package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
	"math/rand"
)

type Image struct{
	width, height int
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.width, m.height)
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) At(x, y int) color.Color {
	p := uint8(rand.Intn(255))
	return color.RGBA{p, p, 255, 255}
}

func main() {
	width := 200
	height := 200
	m := Image{width, height}
	pic.ShowImage(m)
}
