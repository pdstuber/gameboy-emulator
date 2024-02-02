package util

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
)

func PrettyPrintOpcode(opcode types.Opcode) string {
	return fmt.Sprintf("0x%02X", opcode)
}

func PrettyPrintUINT16(value uint16) string {
	return fmt.Sprintf("0x%04X", value)
}

func PrettyPrintBinary(value uint8) string {
	return fmt.Sprintf("0b%08b", value)
}

func CalculateTile(data []byte) types.Tile {
	if len(data) != 16 {
		panic(fmt.Sprintf("tile data must be exactly 16 bytes, but was %d", len(data)))
	}
	var tile types.Tile

	for row := 0; row < 8; row++ {
		firstByte := data[2*row]
		secondByte := data[2*row+1]
		for column := 0; column < 8; column++ {
			tile[row][column] = calculateColor(7-column, firstByte, secondByte)
		}
	}

	return tile
}

func calculateColor(column int, firstByte, secondByte byte) types.Color {
	lsb := ((secondByte >> column) & 0x1) << 1
	msb := (firstByte >> column) & 0x1

	return types.Color(lsb | msb)
}

func GetLeastSignificantBits(number uint16) uint8 {
	return uint8(number & 0xFF)
}

func GetMostSignificantBits(number uint16) uint8 {
	return uint8((number & 0xFF00) >> 8)
}

func RotateLeftWithCarry(number uint8, carry bool) (uint8, bool) {
	newCarry := uint8(number & 0b10000000)

	var c uint8
	if carry {
		c = 0b00000001
	} else {
		c = 0x00000000
	}

	return setCarryBit((number << 1), c), newCarry != uint8(0x0)
}

func setCarryBit(number uint8, carry uint8) uint8 {
	return number | carry
}

func TestBit(number uint8, bitToTest uint8) bool {
	return number&uint8(1<<bitToTest) != 0x00
}
func UINT16FromUINT8(lsb, msb uint8) uint16 {
	return uint16(lsb) | uint16(msb)<<8
}
