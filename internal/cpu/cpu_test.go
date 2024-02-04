package cpu

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetRegisterHL(t *testing.T) {
	cpu := New(true, nil)
	cpu.SetRegisterH(0x8F)
	cpu.SetRegisterL(0xFF)

	require.Equal(t, uint8(0x8F), cpu.GetRegisterH())
	require.Equal(t, uint8(0xFF), cpu.GetRegisterL())
}

func TestFlagZero(t *testing.T) {
	cpu := New(true, nil)
	cpu.UnsetFlagZero()

	require.Equal(t, false, cpu.GetFlagZero())

	cpu.SetFlagZero()

	require.Equal(t, true, cpu.GetFlagZero())
}
