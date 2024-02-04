package instructions

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type LoadFromRegister struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewLoadFromRegister(opcode types.Opcode) *LoadFromRegister {
	return &LoadFromRegister{
		durationInMachineCycles: 2,
		opcode:                  opcode,
	}
}

func (i *LoadFromRegister) Execute(cpu types.CPU) (int, error) {
	var (
		value   uint8
		address = types.Address(cpu.GetRegisterHL())
	)

	switch i.opcode {
	case 0x70:
		value = cpu.GetRegisterB()
	case 0x71:
		value = cpu.GetRegisterC()
	case 0x72:
		value = cpu.GetRegisterD()
	case 0x73:
		value = cpu.GetRegisterE()
	case 0x74:
		value = cpu.GetRegisterH()
	case 0x75:
		value = cpu.GetRegisterL()
	case 0x77:
		value = cpu.GetRegisterA()

	default:
		return 0, fmt.Errorf("unsupported opcode for 16 bit load command: %s", util.PrettyPrintOpcode(i.opcode))
	}

	cpu.WriteMemory(address, value)

	return i.durationInMachineCycles, nil
}
