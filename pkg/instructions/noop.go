package instructions

import (
	"github.com/pdstuber/gameboy-emulator/pkg/types"
)

type NOOP struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewNOOP(opcode types.Opcode) *NOOP {
	return &NOOP{
		opcode:                  opcode,
		durationInMachineCycles: 0,
	}
}

func (i *NOOP) Execute(cpu types.CPU) (int, error) {
	return i.durationInMachineCycles, nil
}
