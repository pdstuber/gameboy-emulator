package instructions

import (
	"github.com/pdstuber/gameboy-emulator/pkg/types"
)

type CompareIndirect struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewCompareIndirect(opcode types.Opcode) *CompareIndirect {
	return &CompareIndirect{
		durationInMachineCycles: 2,
		opcode:                  opcode,
	}
}

func (i *CompareIndirect) Execute(cpu types.CPU) (int, error) {
	var (
		hl              = cpu.GetRegisterHL()
		valueToSubtract = cpu.ReadMemory(types.Address(hl))
		a               = cpu.GetRegisterA()
	)

	zero := a == valueToSubtract
	half_carry := (a & 0xF) < (valueToSubtract & 0xF)
	carry := a < valueToSubtract

	if carry {
		cpu.SetFlagCarry()
	} else {
		cpu.UnsetFlagCarry()
	}
	if zero {
		cpu.SetFlagZero()
	} else {
		cpu.UnsetFlagZero()
	}

	cpu.SetFlagSubtraction()

	if half_carry {
		cpu.SetFlagHalfCarry()
	} else {
		cpu.UnsetFlagHalfCarry()
	}

	return i.durationInMachineCycles, nil
}
