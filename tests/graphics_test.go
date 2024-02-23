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
func Test_Graphics(t *testing.T) {
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

	err = executeNInstructions(cpu, 4034)
	require.NoError(t, err)
	require.Equal(t, uint16(0x0039), cpu.GetProgramCounter())
}
