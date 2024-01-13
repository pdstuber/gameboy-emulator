package instructions

import (
	"github.com/pdstuber/gameboy-emulator/pkg/types"
)

type JumpImmediate struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewJumpImmediate(opcode types.Opcode) *JumpImmediate {
	return &JumpImmediate{
		durationInMachineCycles: 3,
		opcode:                  opcode,
	}
}

func (ji *JumpImmediate) Execute(cpu types.CPU) (int, error) {
	nn := nextWord(cpu)

	cpu.SetProgramCounter(nn)

	return ji.durationInMachineCycles, nil
}
