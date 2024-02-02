package instructions

import (
	"github.com/pdstuber/gameboy-emulator/pkg/types"
)

type Subtract struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewSubtract(opcode types.Opcode) *Subtract {
	return &Subtract{
		durationInMachineCycles: 2,
		opcode:                  opcode,
	}
}

func (i *Subtract) Execute(cpu types.CPU) (int, error) {
	n := int(cpu.ReadMemoryAndIncrementProgramCounter())
	a := int(cpu.GetRegisterA())

	half_carry := ((a & 0xF) - (n & 0xF)) < 0
	carry := (a - n) < 0
	zero := a == n

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

	cpu.SetRegisterA(uint8(a - n))

	return i.durationInMachineCycles, nil
}
