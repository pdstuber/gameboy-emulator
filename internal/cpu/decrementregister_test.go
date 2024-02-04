package cpu

import (
	"testing"

	"github.com/pdstuber/gameboy-emulator/pkg/instructions"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/stretchr/testify/require"
)

func Test_DecrementRegister_Execute(t *testing.T) {
	cpu := New(true, nil)
	i := instructions.NewDecrementRegister(types.Opcode(0x05))

	cpu.SetRegisterB(0x45)

	i.Execute(cpu)

	require.Equal(t, uint8(0x44), cpu.GetRegisterB())
	require.Equal(t, false, cpu.GetFlagZero())

	cpu.SetRegisterB(0x00)

	i.Execute(cpu)

	require.Equal(t, false, cpu.GetFlagZero())
	require.Equal(t, uint8(0xFF), cpu.GetRegisterB())

	cpu.SetRegisterB(0x01)

	i.Execute(cpu)

	require.Equal(t, true, cpu.GetFlagZero())
	require.Equal(t, uint8(0x00), cpu.GetRegisterB())
}
