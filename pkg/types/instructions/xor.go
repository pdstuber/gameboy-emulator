package instructions

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type XOR struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewXOR(opcode types.Opcode) *XOR {
	return &XOR{
		opcode:                  opcode,
		durationInMachineCycles: 1,
	}
}

func (i *XOR) Execute(cpu types.CPU) (int, error) {
	var sourceRegisterValue = cpu.GetRegisterA()
	var targetRegisterValue uint16

	switch i.opcode {
	case 0xA8:
		targetRegisterValue = cpu.GetRegisterB()
	case 0xA9:
		targetRegisterValue = cpu.GetRegisterC()
	case 0xAA:
		targetRegisterValue = cpu.GetRegisterD()
	case 0xAB:
		targetRegisterValue = cpu.GetRegisterE()
	case 0xAC:
		targetRegisterValue = cpu.GetRegisterH()
	case 0xAD:
		targetRegisterValue = cpu.GetRegisterL()
	case 0xAE:
		targetRegisterValue = cpu.GetRegisterHL()
	case 0xAF:
		targetRegisterValue = cpu.GetRegisterA()
	default:
		return 0, fmt.Errorf("unsupported opcode for xor command: %s", util.PrettyPrintOpcode(i.opcode))
	}

	result := sourceRegisterValue ^ targetRegisterValue

	cpu.SetRegisterA(result)

	if result == 0 {
		cpu.SetFlagZ(true)
	} else {
		cpu.SetFlagZ(false)
	}
	cpu.SetFlagN(false)
	cpu.SetFlagH(false)
	cpu.SetFlagC(false)

	return i.durationInMachineCycles, nil
}
