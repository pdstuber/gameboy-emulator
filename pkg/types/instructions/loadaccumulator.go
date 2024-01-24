package instructions

import (
	"github.com/pdstuber/gameboy-emulator/pkg/types"
)

type LoadAccumulator struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewLoadAccumulator(opcode types.Opcode) *LoadAccumulator {
	return &LoadAccumulator{
		durationInMachineCycles: 2,
		opcode:                  opcode,
	}
}

func (i *LoadAccumulator) Execute(cpu types.CPU) (int, error) {
	de := cpu.GetRegisterDE()

	n := cpu.ReadMemory(types.Address(de))

	cpu.SetRegisterA(n)

	return i.durationInMachineCycles, nil
}
