package cpu

import (
	"testing"

	"github.com/pdstuber/gameboy-emulator/pkg/instructions"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/stretchr/testify/require"
)

func Test_RotateLeft_Execute(t *testing.T) {
	cpu := New(true, nil)

	i := instructions.NewRotateLeft(types.Opcode(0x11))
	cpu.SetRegisterC(0xCE)
	cpu.SetFlagCarry()

	i.Execute(cpu)

	require.Equal(t, false, cpu.GetFlagZero())
	require.Equal(t, true, cpu.GetFlagCarry())
	require.Equal(t, uint8(0x9D), cpu.GetRegisterC())

	cpu.SetRegisterC(0xCE)
	cpu.UnsetFlagCarry()

	i.Execute(cpu)

	/*
		sourceRegisterValue = 0x80
		carry = false
		newCarry = true
		result = 0
	*/
	require.Equal(t, false, cpu.GetFlagZero())
	require.Equal(t, true, cpu.GetFlagCarry())
	require.Equal(t, uint8(0x9C), cpu.GetRegisterC())
}
