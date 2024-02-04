package instructions

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type LoadAccumulator struct {
	opcode types.Opcode
}

func NewLoadAccumulator(opcode types.Opcode) *LoadAccumulator {
	return &LoadAccumulator{
		opcode: opcode,
	}
}

func (i *LoadAccumulator) Execute(cpu types.CPU) (int, error) {
	var (
		value  uint8
		cycles int
	)
	switch i.opcode {
	case 0x1A:
		de := cpu.GetRegisterDE()

		value = cpu.ReadMemory(types.Address(de))
		cycles = 2
	case 0xF0:
		n := cpu.ReadMemoryAndIncrementProgramCounter()
		value = cpu.ReadMemory(types.Address(util.UINT16FromUINT8(n, 0xFF)))

		// TODO https://gbdev.gg8.se/wiki/articles/Video_Display#FF44_-_LY_-_LCDC_Y-Coordinate_.28R.29
		cpu.WriteMemory(types.Address(0xFF44), 0x90)
		cycles = 3

	default:
		return 0, fmt.Errorf("unsupported opcode for 16 bit load command: %s", util.PrettyPrintOpcode(i.opcode))
	}

	cpu.SetRegisterA(value)

	return cycles, nil
}
