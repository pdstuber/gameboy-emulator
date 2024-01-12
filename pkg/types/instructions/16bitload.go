package instructions

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type Load16Bit struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func New16BitLoad(opcode types.Opcode) *Load16Bit {
	return &Load16Bit{
		durationInMachineCycles: 3,
		opcode:                  opcode,
	}
}

func (i *Load16Bit) Execute(cpu types.CPU) (int, error) {

	nn := nextWord(cpu)

	switch i.opcode {
	case 0x01:
		cpu.SetRegisterBC(nn)
	case 0x11:
		cpu.SetRegisterDE(nn)
	case 0x21:
		cpu.SetRegisterHL(nn)
	case 0x31:
		cpu.SetRegisterSP(nn)
	default:
		return 0, fmt.Errorf("unsupported opcode for 16 bit load command: %s", util.PrettyPrintOpcode(i.opcode))
	}

	return i.durationInMachineCycles, nil
}
