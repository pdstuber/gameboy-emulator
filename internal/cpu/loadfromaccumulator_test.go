package cpu

import (
	"testing"

	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/pkg/instructions"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/stretchr/testify/require"
)

func Test_LoadFromAccumulator_Execute(t *testing.T) {
	memory := memory.New()
	cpu := New(true, memory)

	cpu.SetRegisterA(0x77)
	cpu.SetRegisterH(0x55)
	cpu.SetRegisterL(0x66)

	i := instructions.NewLoadFromAccumulator(types.Opcode(0x22))

	i.Execute(cpu)

	require.Equal(t, uint8(0x77), cpu.ReadMemory(types.Address(0x6655)))
	require.Equal(t, uint16(0x5666), cpu.GetRegisterHL())
}
