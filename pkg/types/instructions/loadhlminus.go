package instructions

import (
	"github.com/pdstuber/gameboy-emulator/pkg/types"
)

type LoadHLMinus struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewLoadHLMinus(opcode types.Opcode) *LoadHLMinus {
	return &LoadHLMinus{
		opcode:                  opcode,
		durationInMachineCycles: 2,
	}
}

func (i *LoadHLMinus) Execute(cpu types.CPU) (int, error) {
	cpu.WriteMemory(types.Address(cpu.GetRegisterHL()), byte(cpu.GetRegisterA()))
	cpu.SetRegisterHL(cpu.GetRegisterHL() - 1)

	return i.durationInMachineCycles, nil
}
