package cpu

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/types/instructions"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type CPU struct {
	pc     types.Address
	memory *memory.Memory
}

func New(memory *memory.Memory) *CPU {
	cpu := new(CPU)
	cpu.memory = memory
	loadDefaults(cpu)

	return cpu
}
func loadDefaults(cpu *CPU) {
	// skip boot logic
	cpu.pc = 0x100
}

func (c *CPU) FetchAndExecuteNextInstruction() error {
	opcode := types.Opcode(c.memory.Read(c.pc))
	defer func() {
		c.pc++
	}()

	var instruction types.Instruction
	switch opcode {
	case 0x3c:
		instruction = new(instructions.JumpImmediate)

	default:
		return fmt.Errorf("unsupported opcode: %s", util.PrettyPrintOpcode(opcode))
	}

	return instruction.Execute()
}
