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
	var setRegister func(cpu types.CPU, value1, value2 uint8)

	switch i.opcode {
	case 0xC1:
		setRegister = func(cpu types.CPU, value1, value2 uint8) {
			cpu.SetRegisterB(value1)
			cpu.SetRegisterC(value2)
		}
	case 0xD1:
		setRegister = func(cpu types.CPU, value1, value2 uint8) {
			cpu.SetRegisterD(value1)
			cpu.SetRegisterE(value2)
		}
	case 0xE1:
		setRegister = func(cpu types.CPU, value1, value2 uint8) {
			cpu.SetRegisterH(value1)
			cpu.SetRegisterL(value2)
		}
	default:
		return 0, fmt.Errorf("unsupported opcode for push to stack command: %s", util.PrettyPrintOpcode(i.opcode))
	}

	sp := cpu.GetRegisterSP()
	sp += 1

	msb := cpu.ReadMemory(types.Address(sp))
	sp += 1
	lsb := cpu.ReadMemory(types.Address(sp))

	setRegister(cpu, msb, lsb)

	cpu.SetRegisterSP(sp)

	return i.durationInMachineCycles, nil
}
