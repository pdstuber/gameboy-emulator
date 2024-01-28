package ppu

import (
	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

const (
	vramBegin     = 0x8000
	vramEnd       = 0x9FFF
	numberOfTiles = 32
)

type LCDCReader interface {
	GetRegisterLCDC() uint8
}

type PPU struct {
	memory     *memory.Memory
	Pixels     []byte
	lcdcReader LCDCReader
}

func New(memory *memory.Memory, lcdcReader LCDCReader, screenSize int) *PPU {
	return &PPU{
		memory:     memory,
		lcdcReader: lcdcReader,
		Pixels:     make([]byte, screenSize*4),
	}
}

func (p *PPU) Tick() error {
	/*
		ppuInactive := p.lcdcReader.GetRegisterLCDC()&(1<<7) == 0
		if ppuInactive {
			return nil
		}
	*/
	for y := 0; y < numberOfTiles; y++ {
		for x := 0; x < numberOfTiles; x++ {
			currentPosition := y*32 + x
			tilePositionAddress := currentPosition + 0x9800
			tileIndex := p.memory.Read(types.Address(tilePositionAddress))

			tileDataStartAddress := 0x8000 + uint16(tileIndex*16)

			// data for one tile occupies 16 bytes
			var tileData []byte

			for i := 0; i < 16; i += 2 {
				byte1 := p.memory.Read(types.Address(uint16(tileDataStartAddress) + uint16(i)))
				byte2 := p.memory.Read(types.Address(uint16(tileDataStartAddress) + uint16(i+1)))

				tileData = append(tileData, byte1, byte2)
			}
			tile := util.CalculateTile(tileData)
			p.writeToFramebuffer(tile, currentPosition)
		}
	}

	return nil
}

func (p *PPU) writeToFramebuffer(tile types.Tile, currentPosition int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			color := tile[i][j].ToStandardColor()
			p.Pixels[currentPosition] = color.R
			p.Pixels[currentPosition+1] = color.G
			p.Pixels[currentPosition+2] = color.B
			p.Pixels[currentPosition+3] = color.A
		}
	}
}
