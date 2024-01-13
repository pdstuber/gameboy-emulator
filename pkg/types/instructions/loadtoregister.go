package instructions

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type LoadTo8BitRegister struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewLoadTo8BitRegister(opcode types.Opcode) *LoadTo8BitRegister {
	return &LoadTo8BitRegister{
		durationInMachineCycles: 2,
		opcode:                  opcode,
	}
}

func (i *LoadTo8BitRegister) Execute(cpu types.CPU) (int, error) {
	nn := cpu.ReadMemoryAndIncrementProgramCounter()

	switch i.opcode {
	case 0x0E:
		cpu.SetRegisterC(nn)
	case 0x1E:
		cpu.SetRegisterE(nn)
	case 0x2E:
		cpu.SetRegisterL(nn)
	case 0x3E:
		cpu.SetRegisterA(nn)
	default:
		return 0, fmt.Errorf("unsupported opcode for 8 bit load to register command: %s", util.PrettyPrintOpcode(i.opcode))
	}

	return i.durationInMachineCycles, nil
}
