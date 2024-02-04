package instructions

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type LoadTo8BitRegister struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewLoadTo8BitRegister(opcode types.Opcode) *LoadTo8BitRegister {
	return &LoadTo8BitRegister{
		durationInMachineCycles: 2,
		opcode:                  opcode,
	}
}

func (i *LoadTo8BitRegister) Execute(cpu types.CPU) (int, error) {
	var (
		nn          = cpu.ReadMemoryAndIncrementProgramCounter()
		setRegister func(cpu types.CPU, value uint8)
	)

	switch i.opcode {
	case 0x06:
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterB(value) }
	case 0x16:
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterD(value) }
	case 0x26:
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterH(value) }
	case 0x0E:
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterC(value) }
	case 0x1E:
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterE(value) }
	case 0x2E:
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterL(value) }
	case 0x3E:
		setRegister = func(cpu types.CPU, value uint8) { cpu.SetRegisterA(value) }
	default:
		return 0, fmt.Errorf("unsupported opcode for 8 bit load to register command: %s", util.PrettyPrintOpcode(i.opcode))
	}

	setRegister(cpu, nn)

	return i.durationInMachineCycles, nil
}
