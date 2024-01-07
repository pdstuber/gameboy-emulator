package cartridge

import (
	"github.com/pdstuber/gameboy-emulator/pkg/types"
)

type Cartridge struct {
	data []byte
}

func New(romData []byte) *Cartridge {
	return &Cartridge{data: romData}
}

func (c *Cartridge) Read(address types.Address) byte {
	return c.data[address]
}
