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
	pc types.Address
	bc uint16
	sp uint16
	hl uint16
	de uint16
	a  uint16
	b  uint16
	c  uint16
	d  uint16
	e  uint16
	h  uint16
	l  uint16

	flagZ                     bool
	flagN                     bool
	flagH                     bool
	flagC                     bool
	lastWorkingProgramCounter types.Address
	memory                    *memory.Memory
	ppu                       *ppu.PPU
	errorChannel              chan error
}

func New(bootRomLoaded bool, memory *memory.Memory, ppu *ppu.PPU) *CPU {
	cpu := new(CPU)
	cpu.memory = memory
	cpu.ppu = ppu
	cpu.errorChannel = make(chan error)

	if !bootRomLoaded {
		loadDefaults(cpu)
	}

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
	case 0xC3:
		instruction = instructions.NewJumpImmediate()
	case 0x01, 0x11, 0x21, 0x31:
		instruction = instructions.New16BitLoad(opcode)
	case 0xA8, 0xA9, 0xAA, 0xAB, 0xAC, 0xAD, 0xAE, 0xAF:
		instruction = instructions.NewXOR(opcode)
	default:
		return nil, fmt.Errorf("unsupported opcode: %s", util.PrettyPrintOpcode(opcode))
	}

	return instruction, nil
}
func (c *CPU) SetRegisterA(value uint16) {
	c.a = value
}
func (c *CPU) SetRegisterBC(value uint16) {
	c.bc = value
}
func (c *CPU) SetRegisterDE(value uint16) {
	c.de = value
}
func (c *CPU) SetRegisterHL(value uint16) {
	c.hl = value
}
func (c *CPU) SetRegisterSP(value uint16) {
	c.sp = value
}

func (c *CPU) GetRegisterA() uint16 {
	return c.a
}
func (c *CPU) GetRegisterB() uint16 {
	return c.b
}
func (c *CPU) GetRegisterC() uint16 {
	return c.c
}
func (c *CPU) GetRegisterD() uint16 {
	return c.d
}
func (c *CPU) GetRegisterE() uint16 {
	return c.e
}
func (c *CPU) GetRegisterH() uint16 {
	return c.h
}
func (c *CPU) GetRegisterHL() uint16 {
	return c.hl
}
func (c *CPU) GetRegisterL() uint16 {
	return c.l
}

func (c *CPU) SetFlagZ(value bool) {
	c.flagZ = value
}
func (c *CPU) SetFlagN(value bool) {
	c.flagN = value
}
func (c *CPU) SetFlagH(value bool) {
	c.flagH = value
}
func (c *CPU) SetFlagC(value bool) {
	c.flagC = value

}
