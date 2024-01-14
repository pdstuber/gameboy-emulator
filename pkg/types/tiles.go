package types

import "image/color"

type Color uint8

const (
	White Color = iota
	LightGrey
	DarkGrey
	Black
)

type Tile [8][8]Color

func (c Color) ToStandardColor() color.Color {

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
