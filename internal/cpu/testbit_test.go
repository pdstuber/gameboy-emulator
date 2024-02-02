package cpu

import (
	"testing"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/types/instructions"
	"github.com/stretchr/testify/require"
)

func Test_TestBit_Execute(t *testing.T) {

	cpu := New(true, nil)

	i := instructions.NewTestBit(types.Opcode(0x7C))

	cpu.SetRegisterH(0x90)
	cpu.SetRegisterL(0x00)

	i.Execute(cpu)

	require.Equal(t, false, cpu.GetFlagZero())

	cpu.SetRegisterH(0x7F)
	cpu.SetRegisterL(0xFF)

	i.Execute(cpu)

	require.Equal(t, true, cpu.GetFlagZero())
}
