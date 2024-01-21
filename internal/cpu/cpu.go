package cpu

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/internal/ppu"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/types/instructions"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
	"github.com/pkg/errors"
)

const (
	bitmaskFlagZero        = uint8(0x01 << 7)
	bitmaskFlagSubtraction = uint8(0x01 << 6)
	bitmaskFlagHalfCarry   = uint8(0x01 << 5)
	bitMaskFlagCarry       = uint8(0x01 << 4)
)

type CPU struct {
	pc uint16
	sp uint16
	a  uint8
	b  uint8
	c  uint8
	d  uint8
	e  uint8
	f  uint8
	h  uint8
	l  uint8

	flagZ                     bool
	flagN                     bool
	flagH                     bool
	flagC                     bool
	lastWorkingProgramCounter uint16
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

func (c *CPU) Tick() error {
	opcode := types.Opcode(c.ReadMemoryAndIncrementProgramCounter())

	instruction, err := c.decodeInstruction(opcode)
	if err != nil {
		return errors.Wrap(err, "could not decode instruction")
	}

	_, err = instruction.Execute(c)
	if err != nil {
		return err
	}
	c.lastWorkingProgramCounter = c.pc

	return nil
}

func (c *CPU) ReadMemoryAndIncrementProgramCounter() byte {
	data := c.memory.Read(types.Address(c.pc))
	c.pc++
	return data
}

func (c *CPU) WriteMemory(address types.Address, data byte) {
	c.memory.Write(address, data)
}

func (c *CPU) SetProgramCounter(value uint16) {
	c.pc = value
}

func (c *CPU) GetState() string {
	return fmt.Sprintf("PC: %s", util.PrettyPrintUINT16(c.lastWorkingProgramCounter))
}

func (c *CPU) decodeInstruction(opcode types.Opcode) (types.Instruction, error) {
	var instruction types.Instruction
	switch opcode {
	case 0x00:
		instruction = instructions.NewNOOP(opcode)
	case 0xCB:
		nextOpcode := types.Opcode(c.ReadMemoryAndIncrementProgramCounter())
		return c.decodePrefixedInstruction(nextOpcode)
	case 0xC3:
		instruction = instructions.NewJumpImmediate(opcode)
	case 0x01, 0x11, 0x21, 0x31:
		instruction = instructions.NewLoadTo16BitRegister(opcode)
	case 0xA8, 0xA9, 0xAA, 0xAB, 0xAC, 0xAD, 0xAE, 0xAF:
		instruction = instructions.NewXOR(opcode)
	case 0x32:
		instruction = instructions.NewLoadHLMinus(opcode)
	case 0x18, 0x20, 0x30, 0x28, 0x38:
		instruction = instructions.NewJumpConditionalRelative(opcode)
	case 0x0E, 0x1E, 0x2E, 0x3E:
		instruction = instructions.NewLoadTo8BitRegister(opcode)
	case 0xE0, 0xE2:
		instruction = instructions.NewLoadAccumulator(opcode)
	default:
		return nil, fmt.Errorf("unsupported opcode: %s", util.PrettyPrintOpcode(opcode))
	}

	return instruction, nil
}

func (c *CPU) decodePrefixedInstruction(opcode types.Opcode) (types.Instruction, error) {
	var instruction types.Instruction
	switch opcode {
	case 0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4A, 0x4B, 0x4C, 0x4D, 0x4E, 0x4F,
		0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5A, 0x5B, 0x5C, 0x5D, 0x5E, 0x5F,
		0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6A, 0x6B, 0x6C, 0x6D, 0x6E, 0x6F,
		0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7A, 0x7B, 0x7C, 0x7D, 0x7E, 0x7F:

		instruction = instructions.NewTestBit(opcode)
	default:
		return nil, fmt.Errorf("unsupported opcode: %s", util.PrettyPrintOpcode(opcode))
	}

	return instruction, nil
}

func (c *CPU) SetRegisterA(value uint8) {
	c.a = value
}

func (c *CPU) SetRegisterC(value uint8) {
	c.c = value
}

func (c *CPU) SetRegisterE(value uint8) {
	c.e = value
}

func (c *CPU) SetRegisterL(value uint8) {
	c.l = value
}
func (c *CPU) SetRegisterBC(value uint16) {
	c.b = uint8((value & 0xFF00) >> 8)
	c.c = uint8((value & 0xFF))
}

func (c *CPU) SetRegisterDE(value uint16) {
	c.d = uint8((value & 0xFF00) >> 8)
	c.e = uint8((value & 0xFF))
}

func (c *CPU) SetRegisterHL(value uint16) {
	c.h = uint8((value & 0xFF00) >> 8)
	c.l = uint8((value & 0xFF))
}

func (c *CPU) SetRegisterSP(value uint16) {
	c.sp = value
}

func (c *CPU) GetRegisterA() uint8 {
	return c.a
}

func (c *CPU) GetRegisterB() uint8 {
	return c.b
}

func (c *CPU) GetRegisterC() uint8 {
	return c.c
}

func (c *CPU) GetRegisterD() uint8 {
	return c.d
}

func (c *CPU) GetRegisterE() uint8 {
	return c.e
}

func (c *CPU) GetRegisterH() uint8 {
	return c.h
}

func (c *CPU) GetRegisterHL() uint16 {
	return uint16(c.h)<<8 | uint16(c.l)
}

func (c *CPU) GetRegisterL() uint8 {
	return c.l
}

func (c *CPU) UnsetFlagZero() {
	c.f = c.f & ^bitmaskFlagZero
}

func (c *CPU) UnsetFlagHalfCarry() {
	c.f = c.f & ^bitmaskFlagHalfCarry
}

func (c *CPU) UnsetFlagCarry() {
	c.f = c.f & ^bitMaskFlagCarry
}
func (c *CPU) UnsetFlagSubtraction() {
	c.f = c.f & ^bitmaskFlagSubtraction
}
func (c *CPU) SetFlagZero() {
	c.f = c.f | bitmaskFlagZero
}

func (c *CPU) SetFlagCarry() {
	c.f = c.f | bitMaskFlagCarry
}

func (c *CPU) SetFlagHalfCarry() {

	c.f = c.f | bitmaskFlagHalfCarry
}

func (c *CPU) SetFlagSubtraction() {

	c.f = c.f | bitmaskFlagSubtraction
}
func (c *CPU) GetFlagZero() bool {
	return c.f&bitmaskFlagZero != 0x00
}

func (c *CPU) GetFlagCarry() bool {
	return c.f&bitMaskFlagCarry != 0x00
}

func (c *CPU) GetFlagHalfCarry() bool {

	return c.f&bitmaskFlagHalfCarry != 0x00
}
func (c *CPU) GetFlagSubtraction() bool {

	return c.f&bitmaskFlagSubtraction != 0x00
}

func (c *CPU) GetProgramCounter() uint16 {
	return c.pc
}
