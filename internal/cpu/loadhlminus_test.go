package cpu

import (
	"testing"

	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/types/instructions"
	"github.com/stretchr/testify/require"
)

func Test_LoadHLMinus_Execute(t *testing.T) {

	memory := memory.New()
	cpu := New(true, memory)

	i := instructions.NewLoadHLMinus(types.Opcode(0x32))

	cpu.SetRegisterH(0x90)
	cpu.SetRegisterL(0x00)
	i.Execute(cpu)

	require.Equal(t, uint16(0x8FFF), cpu.GetRegisterHL())
	require.Equal(t, uint8(0x8F), cpu.GetRegisterH())
	require.Equal(t, uint8(0xFF), cpu.GetRegisterL())
}
