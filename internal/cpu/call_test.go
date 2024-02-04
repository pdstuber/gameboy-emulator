package cpu

import (
	"testing"

	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/pkg/instructions"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
	"github.com/stretchr/testify/require"
)

const (
	pc = 0x0028
	sp = 0x8899
)

func Test_Call_Execute(t *testing.T) {
	memory := memory.New()
	cpu := New(true, memory)

	i := instructions.NewCall(types.Opcode(0xCD))

	cpu.SetRegisterSP(sp)
	cpu.SetProgramCounter(pc)
	cpu.WriteMemory(types.Address(pc), 0x95)
	cpu.WriteMemory(types.Address(pc+1), 0x00)
	i.Execute(cpu)

	require.Equal(t, util.GetMostSignificantBits(pc+2), cpu.ReadMemory(types.Address(sp-1)))
	require.Equal(t, util.GetLeastSignificantBits(pc+2), cpu.ReadMemory(types.Address(sp-2)))
	require.Equal(t, uint16(0x0095), cpu.GetProgramCounter())
	require.Equal(t, uint16(sp-2), cpu.GetRegisterSP())
}
