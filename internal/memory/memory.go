package memory

import (
	"errors"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
)

const (
	memorySize = 0xFFFF
)

type Memory struct {
	data []byte
}

func New() *Memory {
	return &Memory{
		data: make([]byte, memorySize),
	}
}

func (m *Memory) Load(data []byte, startAddress types.Address) error {
	if int(startAddress)+len(data) > memorySize {
		return errors.New("data to be loaded exceeds capacity")
	}

	copy(m.data[startAddress:], data[:])

	return nil
}

func (m *Memory) Read(address types.Address) byte {
	return m.data[address]
}
