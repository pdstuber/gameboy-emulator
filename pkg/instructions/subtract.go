package instructions

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
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
	var valueToSubtract uint8

	switch i.opcode {
	case 0x90:
		valueToSubtract = cpu.GetRegisterB()
	default:
		return 0, fmt.Errorf("unsupported opcode for subtract command: %s", util.PrettyPrintOpcode(i.opcode))
	}

	a := cpu.GetRegisterA()

	half_carry := (a & 0xF) < (valueToSubtract & 0xF)
	carry := a < valueToSubtract
	zero := a == uint8(0x00)

	result := a - valueToSubtract
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

	cpu.SetFlagSubtraction()

	if half_carry {
		cpu.SetFlagHalfCarry()
	} else {
		cpu.UnsetFlagHalfCarry()
	}

	return i.durationInMachineCycles, nil
}
