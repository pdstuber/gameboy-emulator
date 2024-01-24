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
	case 0xA0:
		registerValue = cpu.GetRegisterB()
	case 0xA1:
		registerValue = cpu.GetRegisterC()
	case 0xA2:
		registerValue = cpu.GetRegisterD()
	case 0xA3:
		registerValue = cpu.GetRegisterE()
	case 0xA4:
		registerValue = cpu.GetRegisterH()
	case 0xA5:
		registerValue = cpu.GetRegisterL()
	case 0xA7:
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
