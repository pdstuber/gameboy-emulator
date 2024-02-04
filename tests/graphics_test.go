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

	var cpu types.CPU = cpu.New(true, memory)

	cpu.SetProgramCounter(0x0095)
	cpu.SetRegisterA(0x7B)
	cpu.SetRegisterB(0x0)
	cpu.SetRegisterC(0x12)
	cpu.SetRegisterH(0x10)
	cpu.SetRegisterL(0x80)

	err = executeNInstructions(cpu, 1)
	require.NoError(t, err)
	require.Equal(t, uint8(0x7B), cpu.GetRegisterC())
}
