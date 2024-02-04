package ppu

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

const (
	vramBegin     = 0x8000
	vramEnd       = 0x9FFF
	ldcd          = 0xFF40
	scx           = 0xFF43
	scy           = 0xFF42
	numberOfTiles = 32
)

type PPU struct {
	memory  *memory.Memory
	Pixels  []byte
	counter int
}

func New(memory *memory.Memory, screenSize int) *PPU {
	return &PPU{
		memory:  memory,
		Pixels:  make([]byte, screenSize*4),
		counter: 0,
	}
}

func (p *PPU) Tick() error {
	/*
		if ppuInactive := p.memory.Read(types.Address(0xFF40))&(1<<7) == 0; ppuInactive {
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
	upLeft := image.Point{0, 0}
	lowRight := image.Point{256, 256}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	img.Pix = p.Pixels
	// Encode as PNG.
	f, _ := os.Create(fmt.Sprintf("images/image%d.png", p.counter))
	png.Encode(f, img)
	p.counter = p.counter + 1
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
