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

func CalculateTile(data []byte) types.Tile {
	if len(data) != 16 {
		panic("tile data must be exactly 16 bytes")
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
