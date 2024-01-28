package cpu

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetRegisterBC(t *testing.T) {
	cpu := New(true, nil)
	value := uint16(0b1010101011111111)

	cpu.SetRegisterBC(value)

	require.Equal(t, uint8(0b10101010), cpu.GetRegisterB())
	require.Equal(t, uint8(0b11111111), cpu.GetRegisterC())
}
