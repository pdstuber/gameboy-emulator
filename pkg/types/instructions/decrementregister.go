package instructions

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type DecrementRegister struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewDecrementRegister(opcode types.Opcode) *DecrementRegister {
	return &DecrementRegister{
		durationInMachineCycles: 1,
		opcode:                  opcode,
	}
}

func (i *DecrementRegister) Execute(cpu types.CPU) (int, error) {
	var (
		registerValue uint8
		setRegister   func(cpu types.CPU, value uint8)
	)

	switch i.opcode {
	case 0x05:
		registerValue = cpu.GetRegisterB()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterB(value) }
	case 0x15:
		registerValue = cpu.GetRegisterD()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterD(value) }
	case 0x3D:
		registerValue = cpu.GetRegisterA()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterA(value) }
	case 0x0D:
		registerValue = cpu.GetRegisterC()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterC(value) }
	case 0x1D:
		registerValue = cpu.GetRegisterE()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterE(value) }
	default:
		return 0, fmt.Errorf("unsupported opcode for decrement register command: %s", util.PrettyPrintOpcode(i.opcode))
	}
	result := uint8(registerValue - 0x1)
	half_carry := (((registerValue & 0xF) + (0x1 & 0xF)) & 0x10) == 0x10

	setRegister(cpu, result)

	if result == 0x00 {
		cpu.SetFlagZero()
	} else {
		cpu.UnsetFlagZero()
	}

	cpu.SetFlagSubtraction()

	if half_carry {
		cpu.SetFlagHalfCarry()
	} else {
		cpu.UnsetFlagHalfCarry()
	}

	return i.durationInMachineCycles, nil
}
