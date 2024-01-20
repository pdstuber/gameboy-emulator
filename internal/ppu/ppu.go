package ppu

import (
	"context"
	"log"

	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
)

const (
	vramBegin     = 0x8000
	vramEnd       = 0x9FFF
	numberOfTiles = 32
)

type PPU struct {
	ticksChannel chan int
	errorChannel chan error
	memory       *memory.Memory
	pixels       []byte
	tiles        []types.Tile
	screenWidth  int
	screenHeight int
}

func New(memory *memory.Memory) *PPU {
	b := make([][]types.Tile, numberOfTiles)
	for i := range b {
		b[i] = make([]types.Tile, numberOfTiles)
	}
	return &PPU{
		ticksChannel: make(chan int),
		errorChannel: make(chan error),
		memory:       memory,
		pixels:       make([]byte, 256*256*8),
		screenWidth:  256,
		screenHeight: 256,
	}
}

func (p *PPU) Start(ctx context.Context) error {
	log.Println("starting ppu")

	for {
		select {
		case <-p.ticksChannel:
			err := p.tick()
			p.errorChannel <- err
		case err := <-p.errorChannel:
			return err
		case <-ctx.Done():
			return nil
		}
	}

}

func (p *PPU) tick() error {

	for x := 0; x < numberOfTiles; x++ {
		for y := 0; y < numberOfTiles; y++ {
			tilePositionAddress := (y*32 + x) + 0x9800
			tileIndex := p.memory.Read(types.Address(tilePositionAddress))

			tileDataStartAddress := 0x8000 + uint16(tileIndex*16)

			// data for one tile occupies 16 bytes
			for i := 0; i < 16; i += 2 {

				byte1 := p.memory.Read(types.Address(uint16(tileDataStartAddress) + uint16(i)))
				byte2 := p.memory.Read(types.Address(uint16(tileDataStartAddress) + uint16(i+1)))

				tile := calculateTile(i, byte1, byte2)

				p.writeToFramebuffer(tile, x, y)

			}
		}
	}

	return nil
}

func (p *PPU) NotifyTicks(ticks int) {
	go func() {
		p.ticksChannel <- ticks
	}()
}

func calculateTile(index int, firstByte, secondByte byte) types.Tile {
	var tile types.Tile

	for i := 0; i < 8; i++ {
		mask := uint8(1 << (7 - i))
		lsb := firstByte & mask
		msb := secondByte & mask

		var color types.Color = types.Color(uint8(lsb) | uint8(msb)<<8)

		tile[index][i] = color
	}
	return tile
}

func (p *PPU) writeToFramebuffer(tile types.Tile, tilePositionX, tilePositionY int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			xr := tilePositionX*8 + j
			yr := tilePositionY*8 + i
			p.pixels[yr][xr] = tile[i][j].ToStandardColor
		}
	}
}
