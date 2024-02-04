package instructions

import (
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type Return struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewReturn(opcode types.Opcode) *Return {
	return &Return{
		durationInMachineCycles: 4,
		opcode:                  opcode,
	}
}

func (i *Return) Execute(cpu types.CPU) (int, error) {
	sp := cpu.GetRegisterSP()

	lsb := cpu.ReadMemory(types.Address(sp + 1))
	msb := cpu.ReadMemory(types.Address(sp + 2))

	cpu.SetProgramCounter(util.UINT16FromUINT8(lsb, msb))
	cpu.SetRegisterSP(sp + 2)

	return i.durationInMachineCycles, nil
}
