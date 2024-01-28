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
	var (
		lsb uint8
		msb uint8
	)

	switch i.opcode {
	case 0xC5:
		lsb = cpu.GetRegisterB()
		msb = cpu.GetRegisterC()
	case 0xD5:
		lsb = cpu.GetRegisterD()
		msb = cpu.GetRegisterE()
	case 0xE5:
		lsb = cpu.GetRegisterH()
		msb = cpu.GetRegisterL()
	default:
		return 0, fmt.Errorf("unsupported opcode for push to stack command: %s", util.PrettyPrintOpcode(i.opcode))
	}

	sp := cpu.GetRegisterSP()

	cpu.WriteMemory(types.Address(sp), msb)
	sp -= 1
	cpu.WriteMemory(types.Address(sp), lsb)
	sp -= 1
	cpu.SetRegisterSP(sp)

	return i.durationInMachineCycles, nil
}
