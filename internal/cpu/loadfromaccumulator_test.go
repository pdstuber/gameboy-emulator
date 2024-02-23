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
	cpu.SetProgramCounter(uint16(0x0025))

	cpu.WriteMemory(types.Address(0x0025), 0x10)
	cpu.WriteMemory(types.Address(0x0026), 0x80)
	i1 := instructions.NewLoadTo16BitRegister(types.Opcode(0x21))
	i1.Execute(cpu)

	i := instructions.NewLoadFromAccumulator(types.Opcode(0x22))

	i.Execute(cpu)

	require.Equal(t, uint8(0x77), cpu.ReadMemory(types.Address(0x8010)))
	require.Equal(t, uint8(0x80), cpu.GetRegisterH())
	require.Equal(t, uint8(0x11), cpu.GetRegisterL())
	require.Equal(t, uint16(0x8011), cpu.GetRegisterHL())
}
