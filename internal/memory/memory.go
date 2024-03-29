package memory

import (
	"fmt"

	"github.com/pdstuber/gameboy-emulator/pkg/types"
)

const (
	memorySize = 0xFFFF
)

type Memory struct {
	data []byte
}

func New() *Memory {
	data := make([]byte, memorySize)
	return &Memory{
		data: data,
	}
}

func (m *Memory) Load(data []byte, startAddress types.Address) error {
	size := len(data)
	if int(startAddress)+size > memorySize+1 {
		return fmt.Errorf("data to be loaded exceeds capacity, size=%d, capacity=%d", size, memorySize)
	}

	copy(m.data[startAddress:], data)

	return nil
}

func (m *Memory) Read(address types.Address) byte {
	return m.data[address]
}

func (m *Memory) Write(address types.Address, b byte) {
	m.data[address] = b
}
