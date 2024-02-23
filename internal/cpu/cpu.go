package cpu

import (
	"fmt"
	"log"

	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/pkg/instructions"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
	"github.com/pkg/errors"
)

const (
	bitmaskFlagZero        = uint8(0x01 << 7)
	bitmaskFlagSubtraction = uint8(0x01 << 6)
	bitmaskFlagHalfCarry   = uint8(0x01 << 5)
	bitMaskFlagCarry       = uint8(0x01 << 4)

	// ClockSpeed = 4194304
	ClockSpeed = 6000
)

type CPU struct {
	pc   uint16
	sp   uint16
	a    uint8
	b    uint8
	c    uint8
	d    uint8
	e    uint8
	f    uint8
	h    uint8
	l    uint8
	lcdc uint8

	lastWorkingProgramCounter uint16
	memory                    *memory.Memory
	errorChannel              chan error
}

func New(bootRomLoaded bool, memory *memory.Memory) *CPU {
	cpu := new(CPU)
	cpu.memory = memory
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

func (c *CPU) Tick(cyclesToExecute int) error {
	i := 0
	for i <= cyclesToExecute {
		cycles, err := c.DecodeAndExecuteNextInstruction()
		if err != nil {
			return err
		}
		c.lastWorkingProgramCounter = c.pc

		// fetching is also one cycle
		i += cycles + 1
	}
	return nil
}

func (c *CPU) DecodeAndExecuteNextInstruction() (int, error) {
	opcode := types.Opcode(c.ReadMemoryAndIncrementProgramCounter())

	instruction, err := c.decodeInstruction(opcode)
	if err != nil {
		return 0, errors.Wrap(err, "could not decode instruction")
	}

	cycles, err := instruction.Execute(c)
	if err != nil {
		return 0, errors.Wrap(err, "could not execute instruction")
	}
	return cycles, nil
}

func (c *CPU) ReadMemoryAndIncrementProgramCounter() byte {
	data := c.ReadMemory(types.Address(c.pc))
	c.pc++
	return data
}

func (c *CPU) ReadMemory(address types.Address) byte {
	return c.memory.Read(address)
}

func (c *CPU) WriteMemory(address types.Address, data byte) {
	c.memory.Write(address, data)
	if address == types.Address(0x0091) {
		log.Println("should not happen")
	}
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
	case 0x06, 0x16, 0x26, 0x0E, 0x1E, 0x2E, 0x3E:
		instruction = instructions.NewLoadTo8BitRegister(opcode)
	case 0xE0, 0xE2, 0x22, 0xEA:
		instruction = instructions.NewLoadFromAccumulator(opcode)
	case 0x04, 0x0C, 0x14, 0x1C, 0x24, 0x2C, 0x34, 0x3C:
		instruction = instructions.NewIncrementRegister(opcode)
	case 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x77:
		instruction = instructions.NewLoadFromRegister(opcode)
	case 0xA0, 0xA1, 0xA2, 0xA3, 0xA4, 0xA5, 0xA7:
		instruction = instructions.NewBitwiseAndRegister(opcode)
	case 0xCD:
		instruction = instructions.NewCall(opcode)
	case 0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x47, 0x48, 0x49, 0x4A, 0x4B, 0x4C, 0x4D, 0x4F,
		0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x57, 0x58, 0x59, 0x5A, 0x5B, 0x5C, 0x5D, 0x5F,
		0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x67, 0x68, 0x69, 0x6A, 0x6B, 0x6C, 0x6D, 0x6F,
		0x78, 0x79, 0x7A, 0x7B, 0x7C, 0x7D, 0x7F:
		instruction = instructions.NewLoadRegister(opcode)
	case 0x1A, 0xF0:
		instruction = instructions.NewLoadAccumulator(opcode)
	case 0xC5, 0xD5, 0xE5, 0xF5:
		instruction = instructions.NewPush(opcode)
	case 0x17:
		instruction = instructions.NewRotateLeft(opcode)
	case 0xC1, 0xD1, 0xE1, 0xF1:
		instruction = instructions.NewPop(opcode)
	case 0x05, 0x15, 0x25, 0x0D, 0x1D, 0x2D, 0x3D:
		instruction = instructions.NewDecrementRegister(opcode)
	case 0x13, 0x23:
		instruction = instructions.NewIncrement16BitRegister(opcode)
	case 0xC9:
		instruction = instructions.NewReturn(opcode)
	case 0xFE:
		instruction = instructions.NewCompare(opcode)
	case 0x90:
		instruction = instructions.NewSubtract(opcode)
	case 0xBE:
		instruction = instructions.NewCompareIndirect(opcode)
	case 0x86:
		instruction = instructions.NewAddIndirect(opcode)
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
	case 0x10, 0x11, 0x12, 0x13, 0x14, 0x15:
		instruction = instructions.NewRotateLeft(opcode)
	default:
		return nil, fmt.Errorf("unsupported prefixed opcode: %s", util.PrettyPrintOpcode(opcode))
	}

	return instruction, nil
}

func (c *CPU) SetRegisterA(value uint8) {
	c.a = value
}

func (c *CPU) SetRegisterB(value uint8) {
	c.b = value
}

func (c *CPU) SetRegisterC(value uint8) {
	c.c = value
}

func (c *CPU) SetRegisterD(value uint8) {
	c.d = value
}

func (c *CPU) SetRegisterH(value uint8) {
	c.h = value
}

func (c *CPU) SetRegisterE(value uint8) {
	c.e = value
}

func (c *CPU) SetRegisterL(value uint8) {
	c.l = value
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

func (c *CPU) GetRegisterDE() uint16 {
	return uint16(c.d)<<8 | uint16(c.e)
}

func (c *CPU) GetRegisterBC() uint16 {
	return uint16(c.b)<<8 | uint16(c.c)
}

func (c *CPU) GetRegisterAF() uint16 {
	return uint16(c.a)<<8 | uint16(c.f)
}

func (c *CPU) GetRegisterSP() uint16 {
	return c.sp
}

func (c *CPU) GetRegisterL() uint8 {
	return c.l
}

func (c *CPU) GetRegisterLCDC() uint8 {
	return c.lcdc
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
