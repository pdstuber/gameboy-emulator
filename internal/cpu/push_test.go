package cpu

import (
	"testing"

	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/types/instructions"
	"github.com/stretchr/testify/require"
)

func Test_Push_Execute(t *testing.T) {
	memory := memory.New()
	cpu := New(true, memory)

	cpu.SetRegisterSP(0x9000)

	cpu.SetRegisterB(b)
	cpu.SetRegisterC(c)

	i := instructions.NewPush(types.Opcode(0xC5))

	i.Execute(cpu)

	sp_post := cpu.GetRegisterSP()

	require.Equal(t, uint16(2), 0x9000-sp_post)
	require.Equal(t, c, cpu.ReadMemory(types.Address(0x9000)))
	require.Equal(t, b, cpu.ReadMemory(types.Address(0x9000-1)))
}
