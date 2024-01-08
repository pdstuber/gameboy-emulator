package cpu

import (
	"context"
	"fmt"
	"log"

	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/internal/ppu"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/types/instructions"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
	"github.com/pkg/errors"
)

type CPU struct {
	pc                        types.Address
	lastWorkingProgramCounter types.Address
	memory                    *memory.Memory
	ppu                       *ppu.PPU
	errorChannel              chan error
}

func New(memory *memory.Memory, ppu *ppu.PPU) *CPU {
	cpu := new(CPU)
	cpu.memory = memory
	cpu.ppu = ppu
	cpu.errorChannel = make(chan error)

	loadDefaults(cpu)

	return cpu
}
func loadDefaults(cpu *CPU) {
	// skip boot logic
	cpu.pc = 0x100
}

func (cpu *CPU) Start(ctx context.Context) error {
	log.Println("starting cpu")
	go func() {
		for {
			cycles, err := cpu.FetchAndExecuteNextInstruction()
			if err != nil {
				cpu.errorChannel <- err
				break
			}
			cpu.ppu.NotifyCycles(cycles)
		}
	}()

	select {
	case err := <-cpu.errorChannel:
		return err
	case <-ctx.Done():
		return nil
	}
}

func (c *CPU) FetchAndExecuteNextInstruction() (int, error) {
	opcode := types.Opcode(c.ReadMemoryAndIncrementProgramCounter())

	instruction, err := decodeInstruction(opcode)
	if err != nil {
		return 0, errors.Wrap(err, "could not decode instruction")
	}

	cycles, err := instruction.Execute(c)
	if err != nil {
		return 0, err
	}
	c.lastWorkingProgramCounter = c.pc

	return cycles, nil
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
	return fmt.Sprintf("PC: %s", util.PrettyPrintAddress(c.lastWorkingProgramCounter))
}

func decodeInstruction(opcode types.Opcode) (types.Instruction, error) {
	var instruction types.Instruction
	switch opcode {
	case 0xc3:
		instruction = instructions.NewJumpImmediate()

	default:
		return nil, fmt.Errorf("unsupported opcode: %s", util.PrettyPrintOpcode(opcode))
	}

	return instruction, nil
}
