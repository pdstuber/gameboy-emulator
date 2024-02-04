package tests

import (
	_ "embed"
	"testing"

	"github.com/pdstuber/gameboy-emulator/internal/cpu"
	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/stretchr/testify/require"
)

var testData = []byte{
	0x31,
	0xfe,
	0xff,
	0xaf,
	0x21,
	0xff,
	0x9f,
	0x32,
	0xcb,
	0x7c,
	0x20,
	0xfb,
	0x21,
	0x26,
	0xff,
	0x0e,
	0x11,
	0x3e,
	0x80,
	0x32,
	0xe2,
	0x0c,
	0x3e,
	0xf3,
	0xe2,
	0x32,
	0x3e,
	0x77,
	0x77,
	0x3e,
	0xfc,
	0xe0,
	0x47,
	0x11,
	0x04,
	0x01,
	0x21,
	0x10,
	0x80,
	0x1a,
	0xcd,
	0x95,
	0x00,
	0xcd,
	0x96,
	0x00,
	0x13,
	0x7b,
	0xfe,
	0x34,
	0x20,
	0xf3,
	0x11,
	0xd8,
	0x00,
	0x06,
	0x08,
	0x1a,
	0x13,
	0x22,
	0x23,
	0x05,
	0x20,
	0xf9,
	0x3e,
	0x19,
	0xea,
	0x10,
	0x99,
	0x21,
	0x2f,
	0x99,
	0x0e,
	0x0c,
	0x3d,
	0x28,
	0x08,
	0x32,
	0x0d,
	0x20,
	0xf9,
	0x2e,
	0x0f,
	0x18,
	0xf3,
	0x67,
	0x3e,
	0x64,
	0x57,
	0xe0,
	0x42,
	0x3e,
	0x91,
	0xe0,
	0x40,
	0x04,
	0x1e,
	0x02,
	0x0e,
	0x0c,
	0xf0,
	0x44,
	0xfe,
	0x90,
	0x20,
	0xfa,
	0x0d,
	0x20,
	0xf7,
	0x1d,
	0x20,
	0xf2,
	0x0e,
	0x13,
	0x24,
	0x7c,
	0x1e,
	0x83,
	0xfe,
	0x62,
	0x28,
	0x06,
	0x1e,
	0xc1,
	0xfe,
	0x64,
	0x20,
	0x06,
	0x7b,
	0xe2,
	0x0c,
	0x3e,
	0x87,
	0xe2,
	0xf0,
	0x42,
	0x90,
	0xe0,
	0x42,
	0x15,
	0x20,
	0xd2,
	0x05,
	0x20,
	0x4f,
	0x16,
	0x20,
	0x18,
	0xcb,
	0x4f,
	0x06,
	0x04,
	0xc5,
	0xcb,
	0x11,
	0x17,
	0xc1,
	0xcb,
	0x11,
	0x17,
	0x05,
	0x20,
	0xf5,
	0x22,
	0x23,
	0x22,
	0x23,
	0xc9,
	0xce,
	0xed,
	0x66,
	0x66,
	0xcc,
	0x0d,
	0x00,
	0x0b,
	0x03,
	0x73,
	0x00,
	0x83,
	0x00,
	0x0c,
	0x00,
	0x0d,
	0x00,
	0x08,
	0x11,
	0x1f,
	0x88,
	0x89,
	0x00,
	0x0e,
	0xdc,
	0xcc,
	0x6e,
	0xe6,
	0xdd,
	0xdd,
	0xd9,
	0x99,
	0xbb,
	0xbb,
	0x67,
	0x63,
	0x6e,
	0x0e,
	0xec,
	0xcc,
	0xdd,
	0xdc,
	0x99,
	0x9f,
	0xbb,
	0xb9,
	0x33,
	0x3e,
	0x3c,
	0x42,
	0xb9,
	0xa5,
	0xb9,
	0xa5,
	0x42,
	0x3c,
	0x21,
	0x04,
	0x01,
	0x11,
	0xa8,
	0x00,
	0x1a,
	0x13,
	0xbe,
	0x20,
	0xfe,
	0x23,
	0x7d,
	0xfe,
	0x34,
	0x20,
	0xf5,
	0x06,
	0x19,
	0x78,
	0x86,
	0x23,
	0x05,
	0x20,
	0xfb,
	0x86,
	0x20,
	0xfe,
	0x3e,
	0x01,
	0xe0,
	0x50,
}

const (
	startAddress uint16 = 0x0055
	uint8Zero    uint8  = 0x00
	uint8One     uint8  = 0x01
	uint8Three   uint8  = 0x03
)

func Test_ScrollLogoPlaySound(t *testing.T) {
	memory := memory.New()
	err := memory.Load(testData, types.Address(0x0000))
	require.NoError(t, err)

	var cpu types.CPU = cpu.New(true, memory)

	cpu.SetProgramCounter(startAddress)

	err = executeNInstructions(cpu, 3)
	require.NoError(t, err)

	require.Equal(t, uint8Zero, cpu.GetRegisterH())
	require.Equal(t, uint8(0x64), cpu.GetRegisterA())
	require.Equal(t, uint8(0x64), cpu.GetRegisterD())
	require.Equal(t, uint16(0x0059), cpu.GetProgramCounter())

	err = executeNInstructions(cpu, 4)
	require.NoError(t, err)
	require.Equal(t, uint8One, cpu.GetRegisterB())
	require.Equal(t, uint16(0x0060), cpu.GetProgramCounter())

	cpu.WriteMemory(types.Address(0xFF00+0x44), 0x90)
	err = executeNInstructions(cpu, 129)
	require.NoError(t, err)
	require.Equal(t, uint8One, cpu.GetRegisterH())
	require.Equal(t, uint16(0x0073), cpu.GetProgramCounter())

	err = executeNInstructions(cpu, 10)
	require.NoError(t, err)
	require.Equal(t, uint8(0x63), cpu.GetRegisterA())
	require.Equal(t, uint8(0x63), cpu.ReadMemory(types.Address(0xFF00+0x42)))
	require.Equal(t, uint16(0x008B), cpu.GetProgramCounter())

	err = executeNInstructions(cpu, 1)
	require.NoError(t, err)
	require.Equal(t, uint8(0x63), cpu.GetRegisterD())
	require.Equal(t, uint16(0x008C), cpu.GetProgramCounter())

	err = executeNInstructions(cpu, 1)
	require.NoError(t, err)
	require.Equal(t, uint16(0x0060), cpu.GetProgramCounter())

	err = executeNInstructions(cpu, 4934)
	require.NoError(t, err)
	require.Equal(t, uint8(0x40), cpu.GetRegisterD())
	require.Equal(t, uint16(0x008C), cpu.GetProgramCounter())

	err = executeNInstructions(cpu, 8885)
	require.NoError(t, err)
	require.Equal(t, uint8(0x01), cpu.GetRegisterD())
	require.Equal(t, uint16(0x008C), cpu.GetProgramCounter())

	err = executeNInstructions(cpu, 1)
	require.NoError(t, err)
	require.Equal(t, uint8(0x01), cpu.GetRegisterD())
	require.Equal(t, uint16(0x0060), cpu.GetProgramCounter())

	err = executeNInstructions(cpu, 146)
	require.NoError(t, err)
	require.Equal(t, uint8One, cpu.GetRegisterB())
	require.Equal(t, uint8Zero, cpu.GetRegisterD())
	require.Equal(t, true, cpu.GetFlagZero())
	require.Equal(t, uint16(0x008E), cpu.GetProgramCounter())

	err = executeNInstructions(cpu, 1)
	require.NoError(t, err)
	require.Equal(t, uint8Zero, cpu.GetRegisterB())
	require.Equal(t, uint8Zero, cpu.GetRegisterD())
	require.Equal(t, true, cpu.GetFlagZero())
	require.Equal(t, uint16(0x008F), cpu.GetProgramCounter())

	err = executeNInstructions(cpu, 3)
	require.NoError(t, err)
	require.Equal(t, uint8(0x20), cpu.GetRegisterD())
	require.Equal(t, uint16(0x0060), cpu.GetProgramCounter())

	err = executeNInstructions(cpu, 4510)
	require.NoError(t, err)
	require.Equal(t, uint8(0x01), cpu.GetRegisterD())
	require.Equal(t, true, cpu.GetFlagZero())
	require.Equal(t, uint16(0x008B), cpu.GetProgramCounter())

	err = executeNInstructions(cpu, 1)
	require.NoError(t, err)
	require.Equal(t, uint8(0x00), cpu.GetRegisterD())
	require.Equal(t, true, cpu.GetFlagZero())
	require.Equal(t, uint16(0x008C), cpu.GetProgramCounter())

	err = executeNInstructions(cpu, 1)
	require.NoError(t, err)
	require.Equal(t, uint8(0x00), cpu.GetRegisterB())
	require.Equal(t, uint16(0x008E), cpu.GetProgramCounter())

	err = executeNInstructions(cpu, 1)
	require.NoError(t, err)
	require.Equal(t, false, cpu.GetFlagZero())
	require.Equal(t, uint16(0x008F), cpu.GetProgramCounter())

	err = executeNInstructions(cpu, 1)
	require.NoError(t, err)
	require.Equal(t, uint16(0x00E0), cpu.GetProgramCounter())

	err = executeNInstructions(cpu, 1)
	require.NoError(t, err)
	require.Equal(t, uint16(0x00E3), cpu.GetProgramCounter())

	err = executeNInstructions(cpu, 3)
	require.NoError(t, err)
	require.Equal(t, uint16(0x00E8), cpu.GetProgramCounter())
}

func executeNInstructions(cpu types.CPU, numberOfInstructions int) error {
	for i := 0; i < numberOfInstructions; i++ {
		_, err := cpu.DecodeAndExecuteNextInstruction()
		if err != nil {
			return err
		}
	}
	return nil
}
