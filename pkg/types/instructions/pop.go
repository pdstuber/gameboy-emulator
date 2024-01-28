package instructions

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type Pop struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewPop(opcode types.Opcode) *Pop {
	return &Pop{
		durationInMachineCycles: 3,
		opcode:                  opcode,
	}
}

func (i *Pop) Execute(cpu types.CPU) (int, error) {
	var setRegister func(cpu types.CPU, value uint16)

	switch i.opcode {
	case 0xC1:
		setRegister = func(cpu types.CPU, value uint16) { cpu.SetRegisterBC(value) }
	case 0xD1:
		setRegister = func(cpu types.CPU, value uint16) { cpu.SetRegisterDE(value) }
	case 0xE1:
		setRegister = func(cpu types.CPU, value uint16) { cpu.SetRegisterHL(value) }
	case 0xF1:
		setRegister = func(cpu types.CPU, value uint16) { cpu.SetRegisterAF(value) }
	default:
		return 0, fmt.Errorf("unsupported opcode for push to stack command: %s", util.PrettyPrintOpcode(i.opcode))
	}

	sp := cpu.GetRegisterSP()
	sp += 1

	value := wordFromAddress(cpu, types.Address(sp))

	sp += 1
	setRegister(cpu, value)

	cpu.SetRegisterSP(sp)

	return i.durationInMachineCycles, nil
}
