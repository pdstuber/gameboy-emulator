package instructions

import (
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type Call struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewCall(opcode types.Opcode) *Call {
	return &Call{
		durationInMachineCycles: 6,
		opcode:                  opcode,
	}
}

func (i *Call) Execute(cpu types.CPU) (int, error) {
	nn := nextWord(cpu)
	sp := cpu.GetRegisterSP()
	sp -= 1

	pc := cpu.GetProgramCounter()
	var pc_lsb uint8 = util.GetLeastSignificantBits(pc)
	var pc_msb uint8 = util.GetMostSignificantBits(pc)

	cpu.WriteMemory(types.Address(sp), pc_msb)
	sp -= 1
	cpu.WriteMemory(types.Address(sp), pc_lsb)
	cpu.SetRegisterSP(sp)
	cpu.SetProgramCounter(nn)

	return i.durationInMachineCycles, nil
}
