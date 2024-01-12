package instructions

import (
	"github.com/pdstuber/gameboy-emulator/pkg/types"
)

type JumpImmediate struct {
	durationInMachineCycles int
	opcode                  uint16
}

func NewJumpImmediate() *JumpImmediate {
	return &JumpImmediate{
		durationInMachineCycles: 3,
		opcode:                  0xC3,
	}
}

func (ji *JumpImmediate) Execute(cpu types.CPU) (int, error) {
	nn := nextWord(cpu)

	cpu.SetProgramCounter(types.Address(nn))

	return ji.durationInMachineCycles, nil
}
