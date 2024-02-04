package instructions

import (
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
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
	lsb := cpu.ReadMemoryAndIncrementProgramCounter()
	msb := cpu.ReadMemoryAndIncrementProgramCounter()
	cpu.SetProgramCounter(util.UINT16FromUINT8(lsb, msb))

	return ji.durationInMachineCycles, nil
}
