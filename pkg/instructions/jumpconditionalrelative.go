package instructions

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type JumpConditionalRelative struct {
	opcode types.Opcode
}

func NewJumpConditionalRelative(opcode types.Opcode) *JumpConditionalRelative {
	return &JumpConditionalRelative{
		opcode: opcode,
	}
}

func (i *JumpConditionalRelative) Execute(cpu types.CPU) (int, error) {
	var (
		offset    = int8(cpu.ReadMemoryAndIncrementProgramCounter())
		cycles    = 2
		condition bool
	)

	switch i.opcode {
	case 0x20:
		condition = !cpu.GetFlagZero()
	case 0x30:
		condition = !cpu.GetFlagCarry()
	case 0x18:
		condition = true
	case 0x28:
		condition = cpu.GetFlagZero()
	case 0x38:
		condition = cpu.GetFlagCarry()

	default:
		return 0, fmt.Errorf("unsupported opcode for jump conditional command: %s", util.PrettyPrintOpcode(i.opcode))
	}
	if condition {
		pc := cpu.GetProgramCounter()
		cpu.SetProgramCounter(uint16(int(pc) + int(offset)))
		cycles = 3
	}
	return cycles, nil
}
