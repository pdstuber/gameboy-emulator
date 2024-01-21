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

func X() types.Tile {
	var tile types.Tile

	data := []byte{
		0x3C, 0x7E, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x7E, 0x5E, 0x7E, 0x0A, 0x7C, 0x56, 0x38, 0x7C,
	}
	// data for one tile occupies 16 bytes
	for i := 0; i < 16; i += 2 {

		byte1 := data[i]
		byte2 := data[i+1]

		for j := 7; j >= 0; j-- {
			mask := uint8(1 << j)
			//fmt.Printf("%08b\n", mask)
			lsb := byte1 & mask
			msb := byte2 & mask
			//fmt.Printf("lsb: %08b\n", lsb>>j)
			//fmt.Printf("msb: %08b\n", msb>>j)
			//fmt.Printf("%2b\n", lsb|msb)
			//fmt.Printf("0x%02X\n", msb)
			//fmt.Printf("0x%02X\n", lsb)
			x := msb>>j + lsb>>j
			fmt.Printf("msb: %d\n", x)
			var color types.Color = types.Color(uint16(lsb) | uint16(msb)<<8)

			tile[i/2][j] = color
		}
	}

	return tile
}

func CalculateTile(data [16]byte) types.Tile {
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
