package tests

import (
	_ "embed"
	"testing"

	"github.com/pdstuber/gameboy-emulator/internal/cpu"
	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/stretchr/testify/require"
)

// https://gbdev.gg8.se/wiki/articles/Gameboy_Bootstrap_ROM
func Test_Graphics_Routine(t *testing.T) {
	memory := memory.New()
	err := memory.Load(testData, types.Address(0x0000))
	require.NoError(t, err)

	err = memory.Load(testData2[0x100:], types.Address(0x100))

	require.NoError(t, err)

	var cpu types.CPU = cpu.New(true, memory)

	cpu.SetProgramCounter(0x0000)

	err = executeNInstructions(cpu, 210)
	require.NoError(t, err)
	require.Equal(t, uint16(0x0027), cpu.GetProgramCounter())

	err = executeNInstructions(cpu, 2)
	require.NoError(t, err)
	require.Equal(t, uint8(0xCE), cpu.GetRegisterA())
	require.Equal(t, uint16(0x0095), cpu.GetProgramCounter())

	err = executeNInstructions(cpu, 1)
	require.NoError(t, err)
	require.Equal(t, uint16(0x0096), cpu.GetProgramCounter())
	require.Equal(t, uint8(0xCE), cpu.GetRegisterC())

	err = executeNInstructions(cpu, 1)
	require.NoError(t, err)
	require.Equal(t, uint16(0x0098), cpu.GetProgramCounter())
	require.Equal(t, uint8(0x04), cpu.GetRegisterB())

	err = executeNInstructions(cpu, 1)
	require.NoError(t, err)
	require.Equal(t, uint16(0x0099), cpu.GetProgramCounter())
	require.Equal(t, uint8(0xCE), cpu.ReadMemory(types.Address(0xFEFD)))
	require.Equal(t, uint8(0x04), cpu.ReadMemory(types.Address(0xFEFC)))
	require.Equal(t, uint16(0xFEFB), cpu.GetRegisterSP())

	err = executeNInstructions(cpu, 1)
	require.NoError(t, err)
	require.Equal(t, uint16(0x009b), cpu.GetProgramCounter())
	require.Equal(t, uint8(0x9C), cpu.GetRegisterC())
	require.Equal(t, uint8(0xCE), cpu.GetRegisterA())

	err = executeNInstructions(cpu, 1)
	require.NoError(t, err)
	require.Equal(t, uint16(0x009c), cpu.GetProgramCounter())
	require.Equal(t, uint8(0x9C), cpu.GetRegisterC())
	require.Equal(t, uint8(0x9D), cpu.GetRegisterA())

	err = executeNInstructions(cpu, 4)
	require.NoError(t, err)
	require.Equal(t, uint16(0x00a1), cpu.GetProgramCounter())
	require.Equal(t, uint8(0x03), cpu.GetRegisterB())
	require.Equal(t, false, cpu.GetFlagZero())
	require.Equal(t, uint8(0x9D), cpu.GetRegisterC())
	require.Equal(t, uint8(0x3B), cpu.GetRegisterA())

	err = executeNInstructions(cpu, 8)
	require.NoError(t, err)
	require.Equal(t, uint16(0x00a1), cpu.GetProgramCounter())
	require.Equal(t, uint8(0x02), cpu.GetRegisterB())
	require.Equal(t, false, cpu.GetFlagZero())

	err = executeNInstructions(cpu, 8)
	require.NoError(t, err)
	require.Equal(t, uint16(0x00a1), cpu.GetProgramCounter())
	require.Equal(t, uint8(0x01), cpu.GetRegisterB())
	require.Equal(t, false, cpu.GetFlagZero())

	err = executeNInstructions(cpu, 8)
	require.NoError(t, err)
	require.Equal(t, uint16(0x00a1), cpu.GetProgramCounter())
	require.Equal(t, uint8(0x00), cpu.GetRegisterB())
	require.Equal(t, true, cpu.GetFlagZero())

	err = executeNInstructions(cpu, 1)
	require.NoError(t, err)
	require.Equal(t, uint16(0x00a3), cpu.GetProgramCounter())
}
