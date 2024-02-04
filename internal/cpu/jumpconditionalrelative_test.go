package cpu

import (
	"testing"

	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/pkg/instructions"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/stretchr/testify/require"
)

func Test_JumpConditionalRelative_Execute(t *testing.T) {
	memory := memory.New()
	cpu := New(true, memory)

	cpu.SetProgramCounter(0x000B)
	cpu.SetFlagZero()
	cpu.WriteMemory(types.Address(0x000B), 0xfb)

	i := instructions.NewJumpConditionalRelative(types.Opcode(0x20))

	i.Execute(cpu)

	require.Equal(t, uint16(0x000C), cpu.GetProgramCounter())
}
