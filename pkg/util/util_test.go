package util

import (
	"testing"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/stretchr/testify/require"
)

func TestCalculateTile(t *testing.T) {
	data := []byte{
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
	color1 := calculateColor(7, 0b10100101, 0b11000011)
	color2 := calculateColor(6, 0b10100101, 0b11000011)
	color3 := calculateColor(5, 0b10100101, 0b11000011)
	color4 := calculateColor(4, 0b10100101, 0b11000011)
	color5 := calculateColor(3, 0b10100101, 0b11000011)
	color6 := calculateColor(2, 0b10100101, 0b11000011)
	color7 := calculateColor(1, 0b10100101, 0b11000011)
	color8 := calculateColor(0, 0b10100101, 0b11000011)

	require.Equal(t, types.Black, color1)
	require.Equal(t, types.DarkGrey, color2)
	require.Equal(t, types.LightGrey, color3)
	require.Equal(t, types.White, color4)
	require.Equal(t, types.White, color5)
	require.Equal(t, types.LightGrey, color6)
	require.Equal(t, types.DarkGrey, color7)
	require.Equal(t, types.Black, color8)
}

func TestCalculateColor2(t *testing.T) {
	color1 := calculateColor(7, 0b10011101, 0b11010010)
	color2 := calculateColor(6, 0b10011101, 0b11010010)
	color3 := calculateColor(5, 0b10011101, 0b11010010)
	color4 := calculateColor(4, 0b10011101, 0b11010010)
	color5 := calculateColor(3, 0b10011101, 0b11010010)
	color6 := calculateColor(2, 0b10011101, 0b11010010)
	color7 := calculateColor(1, 0b10011101, 0b11010010)
	color8 := calculateColor(0, 0b10011101, 0b11010010)

	require.Equal(t, types.Black, color1)
	require.Equal(t, types.DarkGrey, color2)
	require.Equal(t, types.White, color3)
	require.Equal(t, types.Black, color4)
	require.Equal(t, types.LightGrey, color5)
	require.Equal(t, types.LightGrey, color6)
	require.Equal(t, types.DarkGrey, color7)
	require.Equal(t, types.LightGrey, color8)
}

func Test_getLeastSignificatBits(t *testing.T) {
	number := uint16(0b1101100110111001)
	lsb := GetLeastSignificantBits(number)

	expected := uint8(0b10111001)
	require.Equal(t, expected, lsb)
	number = uint16(0x002B)
	lsb = GetLeastSignificantBits(number)

	expected = uint8(0x2B)
	require.Equal(t, expected, lsb)
}

func Test_getMostSignificatBits(t *testing.T) {
	number := uint16(0b1101100110111001)
	msb := GetMostSignificantBits(number)

	expected := uint8(0b11011001)
	require.Equal(t, expected, msb)
}

func Test_RotateLeftWithCarry(t *testing.T) {
	number := uint8(0b11011001)
	result, newCarry := RotateLeftWithCarry(number, true)

	expectedResult := uint8(0b10110011)

	expectedCarry := true

	require.Equal(t, expectedResult, result)
	require.Equal(t, expectedCarry, newCarry)

	number = uint8(0b01011001)
	result, newCarry = RotateLeftWithCarry(number, true)

	expectedResult = uint8(0b10110011)

	expectedCarry = false

	require.Equal(t, expectedResult, result)
	require.Equal(t, expectedCarry, newCarry)
}

func Test_SetCarryBit(t *testing.T) {
	number := uint8(0b11011001)
	carry := uint8(0b00000001)
	result := setCarryBit(number, carry)

	expectedResult := uint8(0b11011001)
	require.Equal(t, expectedResult, result)

	number = uint8(0b11011000)
	carry = uint8(0b00000001)
	result = setCarryBit(number, carry)

	expectedResult = uint8(0b11011001)
	require.Equal(t, expectedResult, result)

	number = uint8(0b11011001)
	carry = uint8(0b00000000)
	result = setCarryBit(number, carry)

	expectedResult = uint8(0b11011001)
	require.Equal(t, expectedResult, result)
}

func Test_TestBit(t *testing.T) {
	number := uint8(0b10011111)
	result := TestBit(number, 7)

	require.Equal(t, true, result)

	number = uint8(0b10101010)
	result = TestBit(number, 7)

	require.Equal(t, true, result)

	number = uint8(0b01111111)
	result = TestBit(number, 7)

	require.Equal(t, false, result)
}
