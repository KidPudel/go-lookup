package main

import (
	"golang.org/x/tour/pic"
	"image/color"
	"image"
	//"fmt"
)


// blueprint, define a type, by giving it's name and what it is going to be at a base
type Image struct {}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 192, 108)
}

func (img Image) At(x, y int) color.Color {
	//fmt.Printf("%T\n", x)
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
