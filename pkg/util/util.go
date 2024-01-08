package util

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
)

func PrettyPrintOpcode(opcode types.Opcode) string {
	return fmt.Sprintf("0x%02X", opcode)
}

func PrettyPrintAddress(address types.Address) string {
	return fmt.Sprintf("0x%04X", address)
}
