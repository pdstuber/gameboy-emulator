package cpu

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type CPU struct {
	pc     types.Address
	memory *memory.Memory
}

func New(memory *memory.Memory) *CPU {
	return &CPU{
		pc:     0x0000,
		memory: memory,
	}
}

func (c *CPU) FetchAndExecuteNextInstruction() error {
	opcode := c.memory.Read(c.pc)
	defer func() {
		c.pc++
	}()

	switch opcode {
	default:
		return fmt.Errorf("unsupported opcode: %s", util.PrettyPrintByte(opcode))
	}
}
