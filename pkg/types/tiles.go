package types

import "image/color"

type Color uint8

const (
	White     Color = 0b00
	LightGrey Color = 0b01
	DarkGrey  Color = 0b10
	Black     Color = 0b11
)

type Tile [8][8]Color

func (c Color) ToStandardColor() color.RGBA {
	// TODO use corret rgb codes for the grey
	switch c {
	case White:
		return color.White
	case Black:
		return color.Black
	case LightGrey:
		return color.Gray16{0x0000}
	case DarkGrey:
		return color.Gray16{}
	}

	return color.Opaque
}
