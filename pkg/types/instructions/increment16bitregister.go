package instructions

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type Increment16BitRegister struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewIncrement16BitRegister(opcode types.Opcode) *Increment16BitRegister {
	return &Increment16BitRegister{
		durationInMachineCycles: 1,
		opcode:                  opcode,
	}
}

func (i *Increment16BitRegister) Execute(cpu types.CPU) (int, error) {
	switch i.opcode {
	case 0x13:
		oldValue := cpu.GetRegisterDE()
		newValue := oldValue + 1
		cpu.SetRegisterD(util.GetMostSignificantBits(newValue))
		cpu.SetRegisterE(util.GetLeastSignificantBits(newValue))
	case 0x23:
		oldValue := cpu.GetRegisterHL()
		newValue := oldValue + 1
		cpu.SetRegisterH(util.GetMostSignificantBits(newValue))
		cpu.SetRegisterL(util.GetLeastSignificantBits(newValue))
	default:
		return 0, fmt.Errorf("unsupported opcode for increment 16 bit register command: %s", util.PrettyPrintOpcode(i.opcode))
	}

	return i.durationInMachineCycles, nil
}
