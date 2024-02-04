package cpu

import (
	"testing"

	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/pkg/instructions"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/stretchr/testify/require"
)

func Test_LoadTo16BitRegister_Execute(t *testing.T) {
	memory := memory.New()
	cpu := New(true, memory)

	i := instructions.NewLoadTo16BitRegister(types.Opcode(0x21))

	///00000024: 21  !
	//00000025: 10  .
	//00000026: 80
	cpu.SetProgramCounter(uint16(0x0025))
	cpu.WriteMemory(types.Address(0x0025), 0x10)
	cpu.WriteMemory(types.Address(0x0026), 0x80)
	i.Execute(cpu)

	require.Equal(t, uint16(0x8010), cpu.GetRegisterHL())
	require.Equal(t, uint8(0x80), cpu.GetRegisterH())
	require.Equal(t, uint8(0x10), cpu.GetRegisterL())
}
