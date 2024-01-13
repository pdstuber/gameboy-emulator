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
