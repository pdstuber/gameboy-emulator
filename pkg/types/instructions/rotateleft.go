package instructions

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type RotateLeft struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewRotateLeft(opcode types.Opcode) *RotateLeft {
	return &RotateLeft{
		opcode:                  opcode,
		durationInMachineCycles: 1,
	}
}

// "RL C" (rotate left C register), NOT to be confused with "RLC" (rotate left circular) rotates the C register one bit to the left.
// Previous carry flag becomes the least-significant bit, and previous Most Significant Bit becomes Carry.
func (i *RotateLeft) Execute(cpu types.CPU) (int, error) {
	var (
		sourceRegisterValue uint8
		setRegister         func(cpu types.CPU, value uint8)
	)

	switch i.opcode {
	case 0x10:
		sourceRegisterValue = cpu.GetRegisterB()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterB(value) }
	case 0x11:
		sourceRegisterValue = cpu.GetRegisterC()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterC(value) }
	case 0x12:
		sourceRegisterValue = cpu.GetRegisterD()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterD(value) }
	case 0x13:
		sourceRegisterValue = cpu.GetRegisterE()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterE(value) }
	case 0x14:
		sourceRegisterValue = cpu.GetRegisterH()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterH(value) }
	case 0x15:
		sourceRegisterValue = cpu.GetRegisterL()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterL(value) }
	case 0x17:
		sourceRegisterValue = cpu.GetRegisterA()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterA(value) }
	default:
		return 0, fmt.Errorf("unsupported opcode for xor command: %s", util.PrettyPrintOpcode(i.opcode))
	}

	carry := cpu.GetFlagCarry()
	result, newCarry := util.RotateLeftWithCarry(sourceRegisterValue, carry)

	if newCarry {
		cpu.GetFlagCarry()
	} else {
		cpu.UnsetFlagCarry()
	}

	setRegister(cpu, result)

	if result == 0 {
		cpu.SetFlagZero()
	} else {
		cpu.UnsetFlagZero()
	}

	cpu.UnsetFlagSubtraction()
	cpu.UnsetFlagHalfCarry()

	return i.durationInMachineCycles, nil
}
