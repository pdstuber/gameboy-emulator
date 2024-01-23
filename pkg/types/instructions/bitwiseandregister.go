package instructions

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type BitwiseAndRegister struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewBitwiseAndRegister(opcode types.Opcode) *BitwiseAndRegister {
	return &BitwiseAndRegister{
		durationInMachineCycles: 1,
		opcode:                  opcode,
	}
}

func (i *BitwiseAndRegister) Execute(cpu types.CPU) (int, error) {
	var (
		registerValue uint8
		valueA        = cpu.GetRegisterA()
	)

	switch i.opcode {
	case 0x0A:
		registerValue = cpu.GetRegisterB()
	case 0x1A:
		registerValue = cpu.GetRegisterC()
	case 0x2A:
		registerValue = cpu.GetRegisterD()
	case 0x3A:
		registerValue = cpu.GetRegisterE()
	case 0x4A:
		registerValue = cpu.GetRegisterH()
	case 0x5A:
		registerValue = cpu.GetRegisterL()
	case 0x7A:
		registerValue = cpu.GetRegisterA()
	default:
		return 0, fmt.Errorf("unsupported opcode for increment register command: %s", util.PrettyPrintOpcode(i.opcode))
	}

	result := valueA & registerValue

	cpu.SetRegisterA(result)

	if result == 0x0 {
		cpu.SetFlagZero()
	} else {
		cpu.UnsetFlagZero()
	}

	cpu.UnsetFlagSubtraction()
	cpu.UnsetFlagCarry()
	cpu.SetFlagHalfCarry()

	return i.durationInMachineCycles, nil
}
