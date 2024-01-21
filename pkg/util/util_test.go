package util

import (
	"testing"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/stretchr/testify/require"
)

func TestCalculateTile(t *testing.T) {

	data := [16]byte{
		0b10100101,
		0b11000011,
		0b10011101,
		0b11010010,
		0b11100011,
		0b10101010,
		0b11001100,
		0b10110110,
		0b11010111,
		0b10110110,
		0b11111111,
		0b10000001,
		0b10101010,
		0b11000000,
		0b10000000,
		0b11111001,
	}
	tile := CalculateTile(data)

	var expectedTile types.Tile = [8][8]types.Color{
		{types.Black, types.DarkGrey, types.LightGrey, types.White, types.White, types.LightGrey, types.DarkGrey, types.Black},
		{types.Black, types.DarkGrey, types.White, types.Black, types.LightGrey, types.LightGrey, types.DarkGrey, types.LightGrey},
		{types.Black, types.LightGrey, types.Black, types.White, types.DarkGrey, types.White, types.Black, types.LightGrey},
		{types.Black, types.LightGrey, types.DarkGrey, types.DarkGrey, types.LightGrey, types.Black, types.DarkGrey, types.White},
		{types.Black, types.LightGrey, types.DarkGrey, types.Black, types.White, types.Black, types.Black, types.LightGrey},
		{types.Black, types.LightGrey, types.LightGrey, types.LightGrey, types.LightGrey, types.LightGrey, types.LightGrey, types.Black},
		{types.Black, types.DarkGrey, types.LightGrey, types.White, types.LightGrey, types.White, types.LightGrey, types.White},
		{types.Black, types.DarkGrey, types.DarkGrey, types.DarkGrey, types.DarkGrey, types.White, types.White, types.DarkGrey},
	}
	require.Equal(t, expectedTile, tile)
}

func TestCalculateColor(t *testing.T) {

	color1 := calculateColor(0, 0b10100101, 0b11000011)
	color2 := calculateColor(1, 0b10100101, 0b11000011)
	color3 := calculateColor(2, 0b10100101, 0b11000011)
	color4 := calculateColor(3, 0b10100101, 0b11000011)
	color5 := calculateColor(4, 0b10100101, 0b11000011)
	color6 := calculateColor(5, 0b10100101, 0b11000011)
	color7 := calculateColor(6, 0b10100101, 0b11000011)
	color8 := calculateColor(7, 0b10100101, 0b11000011)

	require.Equal(t, types.Black, color1)
	require.Equal(t, types.DarkGrey, color2)
	require.Equal(t, types.LightGrey, color3)
	require.Equal(t, types.White, color4)
	require.Equal(t, types.White, color5)
	require.Equal(t, types.LightGrey, color6)
	require.Equal(t, types.DarkGrey, color7)
	require.Equal(t, types.Black, color8)
}
