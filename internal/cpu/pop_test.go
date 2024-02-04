package cpu

import (
	"testing"

	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/pkg/instructions"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/stretchr/testify/require"
)

const (
	b uint8 = 0x57
	c uint8 = 0x58
)

func Test_Pop_Execute(t *testing.T) {
	memory := memory.New()
	cpu := New(true, memory)

	cpu.SetRegisterSP(0x9000)
	sp_prior := cpu.GetRegisterSP()

	cpu.WriteMemory(types.Address(sp_prior+1), b)
	cpu.WriteMemory(types.Address(sp_prior+2), c)

	i := instructions.NewPop(types.Opcode(0xC1))

	i.Execute(cpu)

	sp_post := cpu.GetRegisterSP()

	require.Equal(t, uint16(2), sp_post-sp_prior)
	require.Equal(t, b, cpu.GetRegisterB())
	require.Equal(t, c, cpu.GetRegisterC())
}
