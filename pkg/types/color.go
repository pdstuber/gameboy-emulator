package types

import "image/color"

/*
System.Drawing.ColorTranslator.FromHtml(Renga.Config.GetProperty("dmgColor1", "#9BBC0F")),
System.Drawing.ColorTranslator.FromHtml(Renga.Config.GetProperty("dmgColor2", "#8BAC0F")),
System.Drawing.ColorTranslator.FromHtml(Renga.Config.GetProperty("dmgColor3", "#306230")),
System.Drawing.ColorTranslator.FromHtml(Renga.Config.GetProperty("dmgColor4", "#0F380F")),
*/
type Color uint8

const (
	White     Color = 0b00
	LightGrey Color = 0b01
	DarkGrey  Color = 0b10
	Black     Color = 0b11
)

func (c Color) ToStandardColor() color.RGBA {
	// TODO use corret rgb codes for the grey
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
