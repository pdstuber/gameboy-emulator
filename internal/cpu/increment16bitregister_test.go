package cpu

import (
	"testing"

	"github.com/pdstuber/gameboy-emulator/pkg/instructions"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/stretchr/testify/require"
)

func Test_Increment16BitRegister_Execute(t *testing.T) {
	cpu := New(true, nil)

	i := instructions.NewIncrement16BitRegister(types.Opcode(0x23))

	cpu.SetRegisterH(0x80)
	cpu.SetRegisterL(0x10)

	i.Execute(cpu)

	require.Equal(t, uint16(0x8011), cpu.GetRegisterHL())
}
