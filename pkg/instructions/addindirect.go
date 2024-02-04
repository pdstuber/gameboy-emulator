package instructions

import (
	"github.com/pdstuber/gameboy-emulator/pkg/types"
)

type AddIndirect struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewAddIndirect(opcode types.Opcode) *AddIndirect {
	return &AddIndirect{
		durationInMachineCycles: 2,
		opcode:                  opcode,
	}
}

func (i *AddIndirect) Execute(cpu types.CPU) (int, error) {
	var (
		hl         = cpu.GetRegisterHL()
		valueToAdd = cpu.ReadMemory(types.Address(hl))
		a          = cpu.GetRegisterA()
	)
	half_carry := (a&0xF)+(valueToAdd&0xF) > 0xF
	carry := a+valueToAdd > 0xFF

	result := a + valueToAdd
	zero := result == 0

	cpu.SetRegisterA(result)

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

	cpu.UnsetFlagSubtraction()

	if half_carry {
		cpu.SetFlagHalfCarry()
	} else {
		cpu.UnsetFlagHalfCarry()
	}

	return i.durationInMachineCycles, nil
}
