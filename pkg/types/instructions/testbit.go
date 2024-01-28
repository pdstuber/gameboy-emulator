package instructions

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

type TestBit struct {
	durationInMachineCycles int
	opcode                  types.Opcode
}

func NewTestBit(opcode types.Opcode) *TestBit {
	return &TestBit{
		opcode:                  opcode,
		durationInMachineCycles: 1,
	}
}

func (i *TestBit) Execute(cpu types.CPU) (int, error) {
	var (
		testBit     uint8
		valueToTest uint8
	)
	switch i.opcode {

	case 0x7C:
		testBit = 7
		valueToTest = cpu.GetRegisterH()
	default:
		return 0, fmt.Errorf("unsupported opcode for test bit command: %s", util.PrettyPrintOpcode(i.opcode))
	}

	result := util.TestBit(valueToTest, testBit)

	if result {
		cpu.UnsetFlagZero()
	} else {
		cpu.SetFlagZero()
	}

	cpu.UnsetFlagSubtraction()
	cpu.SetFlagHalfCarry()

	return i.durationInMachineCycles, nil
}
