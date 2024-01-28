package types

import "image/color"

type Color uint8

const (
	White     Color = 0b00
	LightGrey Color = 0b01
	DarkGrey  Color = 0b10
	Black     Color = 0b11
)

func (c Color) ToStandardColor() color.RGBA {
	switch c {
	case White:
		return color.RGBA{R: 255, G: 255, B: 255, A: 255}
	case Black:
		return color.RGBA{R: 0, G: 0, B: 0, A: 255}
	case LightGrey:
		return color.RGBA{R: 211, G: 211, B: 211, A: 255}
	case DarkGrey:
		return color.RGBA{R: 128, G: 128, B: 128, A: 255}
	}

	return color.RGBA{R: 0, G: 0, B: 0, A: 0}
}
