package instructions

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type LoadRegister struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewLoadRegister(opcode types.Opcode) *LoadRegister {
	return &LoadRegister{
		durationInMachineCycles: 1,
		opcode:                  opcode,
	}
}

func (i *LoadRegister) Execute(cpu types.CPU) (int, error) {
	var (
		sourceValue uint8
		setRegister func(cpu types.CPU, value uint8)
	)

	switch i.opcode {
	case 0x40:
		sourceValue = cpu.GetRegisterB()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterB(value) }
	case 0x41:
		sourceValue = cpu.GetRegisterC()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterB(value) }
	case 0x42:
		sourceValue = cpu.GetRegisterD()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterB(value) }
	case 0x43:
		sourceValue = cpu.GetRegisterE()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterB(value) }
	case 0x44:
		sourceValue = cpu.GetRegisterH()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterB(value) }
	case 0x45:
		sourceValue = cpu.GetRegisterL()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterB(value) }
	case 0x47:
		sourceValue = cpu.GetRegisterA()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterB(value) }
	case 0x48:
		sourceValue = cpu.GetRegisterB()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterC(value) }
	case 0x49:
		sourceValue = cpu.GetRegisterC()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterC(value) }
	case 0x4A:
		sourceValue = cpu.GetRegisterD()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterC(value) }
	case 0x4B:
		sourceValue = cpu.GetRegisterE()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterC(value) }
	case 0x4C:
		sourceValue = cpu.GetRegisterH()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterC(value) }
	case 0x4D:
		sourceValue = cpu.GetRegisterL()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterC(value) }
	case 0x4F:
		sourceValue = cpu.GetRegisterA()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterC(value) }
	case 0x78:
		sourceValue = cpu.GetRegisterB()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterA(value) }
	case 0x79:
		sourceValue = cpu.GetRegisterC()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterA(value) }
	case 0x7A:
		sourceValue = cpu.GetRegisterD()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterA(value) }
	case 0x7B:
		sourceValue = cpu.GetRegisterE()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterA(value) }
	case 0x7C:
		sourceValue = cpu.GetRegisterH()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterA(value) }
	case 0x7D:
		sourceValue = cpu.GetRegisterL()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterA(value) }
	case 0x57:
		sourceValue = cpu.GetRegisterA()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterD(value) }
	case 0x67:
		sourceValue = cpu.GetRegisterA()
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterH(value) }
	default:
		return 0, fmt.Errorf("unsupported opcode for load register command: %s", util.PrettyPrintOpcode(i.opcode))
	}

	setRegister(cpu, sourceValue)

	return i.durationInMachineCycles, nil
}
