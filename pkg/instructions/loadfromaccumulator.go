package instructions

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type LoadFromAccumulator struct {
	opcode types.Opcode
}

func NewLoadFromAccumulator(opcode types.Opcode) *LoadFromAccumulator {
	return &LoadFromAccumulator{
		opcode: opcode,
	}
}

func (i *LoadFromAccumulator) Execute(cpu types.CPU) (int, error) {
	var (
		lsb_value uint8
		msb_value uint8 = 0xFF
		cycles    int
	)

	switch i.opcode {
	case 0xE0:
		lsb_value = cpu.ReadMemoryAndIncrementProgramCounter()
		cycles = 3
	case 0xEA:
		lsb_value = cpu.ReadMemoryAndIncrementProgramCounter()
		msb_value = cpu.ReadMemoryAndIncrementProgramCounter()
		cycles = 4
	case 0xE2:
		lsb_value = cpu.GetRegisterC()
		cycles = 2
	case 0x22:
		lsb_value = cpu.GetRegisterH()
		msb_value = cpu.GetRegisterL()
		newValue := util.UINT16FromUINT8(lsb_value, msb_value) + 1
		cpu.SetRegisterH(util.GetLeastSignificantBits(newValue))
		cpu.SetRegisterL(util.GetMostSignificantBits(newValue))
		cycles = 2

	default:
		return 0, fmt.Errorf("unsupported opcode for 16 bit load command: %s", util.PrettyPrintOpcode(i.opcode))
	}

	address := util.UINT16FromUINT8(lsb_value, msb_value)

	cpu.WriteMemory(types.Address(address), cpu.GetRegisterA())

	return cycles, nil
}
