package cpu

import (
	"testing"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/types/instructions"
	"github.com/stretchr/testify/require"
)

const (
	h = 0x55
	l = 0xFF
)

func Test_Increment16BitRegister_Execute(t *testing.T) {

	cpu := New(true, nil)

	i := instructions.NewIncrement16BitRegister(types.Opcode(0x23))

	cpu.SetRegisterH(h)
	cpu.SetRegisterL(l)

	i.Execute(cpu)

	require.Equal(t, uint16(0x5600), cpu.GetRegisterHL())
}
