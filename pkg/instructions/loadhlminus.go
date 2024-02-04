package instructions

import (
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
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

	hl := cpu.GetRegisterHL() - 1

	cpu.SetRegisterH(util.GetLeastSignificantBits(hl))
	cpu.SetRegisterL(util.GetMostSignificantBits(hl))

	return i.durationInMachineCycles, nil
}
