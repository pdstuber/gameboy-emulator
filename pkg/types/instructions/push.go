package instructions

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type Push struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewPush(opcode types.Opcode) *Push {
	return &Push{
		durationInMachineCycles: 4,
		opcode:                  opcode,
	}
}

func (i *Push) Execute(cpu types.CPU) (int, error) {
	var value uint16

	switch i.opcode {
	case 0xC5:
		value = cpu.GetRegisterBC()
	case 0xD5:
		value = cpu.GetRegisterDE()
	case 0xE5:
		value = cpu.GetRegisterHL()
	case 0xF5:
		value = cpu.GetRegisterAF()
	default:
		return 0, fmt.Errorf("unsupported opcode for push to stack command: %s", util.PrettyPrintOpcode(i.opcode))
	}

	sp := cpu.GetRegisterSP()
	sp -= 1
	msb := util.GetMostSignificantBits(value)
	lsb := util.GetLeastSignificantBits(value)

	cpu.WriteMemory(types.Address(sp), msb)
	sp -= 1
	cpu.WriteMemory(types.Address(sp), lsb)

	return i.durationInMachineCycles, nil
}
