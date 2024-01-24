package instructions

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type LoadFromAccumulator struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewLoadFromAccumulator(opcode types.Opcode) *LoadFromAccumulator {
	return &LoadFromAccumulator{
		durationInMachineCycles: 3,
		opcode:                  opcode,
	}
}

func (i *LoadFromAccumulator) Execute(cpu types.CPU) (int, error) {
	var (
		lsb uint8
		msb uint8 = 0xFF
	)

	switch i.opcode {
	case 0xE0:
		lsb = cpu.ReadMemoryAndIncrementProgramCounter()
	case 0xE2:
		lsb = cpu.GetRegisterC()

	default:
		return 0, fmt.Errorf("unsupported opcode for 16 bit load command: %s", util.PrettyPrintOpcode(i.opcode))
	}

	address := uint16(lsb) | uint16(msb)<<8

	cpu.WriteMemory(types.Address(address), cpu.GetRegisterA())

	return i.durationInMachineCycles, nil
}
