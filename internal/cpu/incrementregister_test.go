package cpu

import (
	"testing"

	"github.com/pdstuber/gameboy-emulator/pkg/instructions"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/stretchr/testify/require"
)

func Test_IncrementRegister_Execute(t *testing.T) {
	cpu := New(true, nil)

	cpu.SetRegisterC(0x45)
	i := instructions.NewIncrementRegister(types.Opcode(0x0C))

	i.Execute(cpu)

	require.Equal(t, uint8(0x46), cpu.GetRegisterC())
}
