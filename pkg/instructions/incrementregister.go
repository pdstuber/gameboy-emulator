package instructions

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type IncrementRegister struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewIncrementRegister(opcode types.Opcode) *IncrementRegister {
	return &IncrementRegister{
		durationInMachineCycles: 1,
		opcode:                  opcode,
	}
}

func (i *IncrementRegister) Execute(cpu types.CPU) (int, error) {
	var (
		registerValue uint8
		setRegister   func(cpu types.CPU, value uint8)
	)

	switch i.opcode {
	case 0x0C:
		registerValue = cpu.GetRegisterC()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterC(value) }
	case 0x04:
		registerValue = cpu.GetRegisterB()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterB(value) }
	case 0x24:
		registerValue = cpu.GetRegisterH()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterH(value) }
	default:
		return 0, fmt.Errorf("unsupported opcode for increment register command: %s", util.PrettyPrintOpcode(i.opcode))
	}
	if registerValue == uint8(0xFF) {
		cpu.SetFlagZero()
	} else {
		cpu.UnsetFlagZero()
	}
	result := registerValue + 0x1
	half_carry := (((registerValue & 0xF) + (0x1 & 0xF)) & 0x10) == 0x10

	setRegister(cpu, result)

	cpu.UnsetFlagSubtraction()

	if half_carry {
		cpu.SetFlagHalfCarry()
	} else {
		cpu.UnsetFlagHalfCarry()
	}

	return i.durationInMachineCycles, nil
}
