package instructions

import "github.com/pdstuber/gameboy-emulator/pkg/types"

func nextWord(cpu types.CPU) uint16 {
	lsb := cpu.ReadMemoryAndIncrementProgramCounter()
	msb := cpu.ReadMemoryAndIncrementProgramCounter()

	return uint16(lsb) | uint16(msb)<<8
}
