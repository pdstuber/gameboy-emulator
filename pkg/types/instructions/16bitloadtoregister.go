package instructions

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type LoadTo16BitRegister struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewLoadTo16BitRegister(opcode types.Opcode) *LoadTo16BitRegister {
	return &LoadTo16BitRegister{
		durationInMachineCycles: 3,
		opcode:                  opcode,
	}
}

func (i *LoadTo16BitRegister) Execute(cpu types.CPU) (int, error) {
	lsb := cpu.ReadMemoryAndIncrementProgramCounter()
	msb := cpu.ReadMemoryAndIncrementProgramCounter()

	switch i.opcode {
	case 0x01:
		cpu.SetRegisterB(lsb)
		cpu.SetRegisterC(msb)
	case 0x11:
		cpu.SetRegisterD(lsb)
		cpu.SetRegisterE(msb)
	case 0x21:
		cpu.SetRegisterH(lsb)
		cpu.SetRegisterL(msb)
	case 0x31:
		cpu.SetRegisterSP(util.UINT16FromUINT8(lsb, msb))
	default:
		return 0, fmt.Errorf("unsupported opcode for 16 bit load command: %s", util.PrettyPrintOpcode(i.opcode))
	}

	return i.durationInMachineCycles, nil
}
