package cpu

import (
	"testing"

	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/pkg/instructions"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/stretchr/testify/require"
)

func Test_Return_Execute(t *testing.T) {
	memory := memory.New()
	cpu := New(true, memory)

	r := instructions.NewReturn(types.Opcode(0xC9))
	c := instructions.NewCall(types.Opcode(0xCD))

	cpu.SetRegisterSP(0x4455)

	cpu.SetProgramCounter(0x0029)

	cpu.WriteMemory(types.Address(0x0029), 0x95)
	cpu.WriteMemory(types.Address(0x002a), 0x00)

	c.Execute(cpu)

	r.Execute(cpu)

	require.Equal(t, uint16(0x002b), cpu.GetProgramCounter())
}
