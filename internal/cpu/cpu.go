package cpu

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/types/instructions"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
	"github.com/pkg/errors"
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
	opcode := types.Opcode(c.ReadMemoryAndIncrementProgramCounter())

	instruction, err := decodeInstruction(opcode)
	if err != nil {
		return errors.Wrap(err, "could not decode instruction")
	}

	return instruction.Execute(c)
}

func (c *CPU) ReadMemoryAndIncrementProgramCounter() byte {
	data := c.memory.Read(c.pc)
	c.pc++
	return data
}

func (c *CPU) SetProgramCounter(address types.Address) {
	c.pc = address
}

func (c *CPU) GetState() string {
	return fmt.Sprintf("PC: %s", util.PrettyPrintAddress(c.pc))
}

func decodeInstruction(opcode types.Opcode) (types.Instruction, error) {
	var instruction types.Instruction
	switch opcode {
	case 0xc3:
		instruction = new(instructions.JumpImmediate)

	default:
		return nil, fmt.Errorf("unsupported opcode: %s", util.PrettyPrintOpcode(opcode))
	}

	return instruction, nil
}
