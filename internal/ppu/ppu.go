package ppu

import (
	"context"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
)

const (
	vramBegin     = 0x8000
	vramEnd       = 0x9FFF
	numberOfTiles = 32
)

type PPU struct {
	cyclesChannel chan int
	errorChannel  chan error
	memory        *memory.Memory
	pixels        []byte
	tiles         []types.Tile
	screenWidth   int
	screenHeight  int
}

func New(memory *memory.Memory) *PPU {
	b := make([][]types.Tile, numberOfTiles)
	for i := range b {
		b[i] = make([]types.Tile, numberOfTiles)
	}
	return &PPU{
		cyclesChannel: make(chan int),
		errorChannel:  make(chan error),
		memory:        memory,
		pixels:        make([]byte, 256*256*8),
		screenWidth:   256,
		screenHeight:  256,
	}
}

func (p *PPU) Start(ctx context.Context) error {
	log.Println("starting ppu")
	go func() {
		ebiten.SetWindowSize(p.screenWidth*6, p.screenHeight*6)
		ebiten.SetWindowTitle("Gameboy Emulator")
		if err := ebiten.RunGame(p); err != nil {
			log.Fatal(err)
		}

	}()

	select {
	case err := <-p.errorChannel:
		return err
	case <-ctx.Done():
		return nil
	}
}

func (p *PPU) NotifyCycles(cycles int) {
	go func() {
		p.cyclesChannel <- cycles
	}()
}

func (p *PPU) Update() error {
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

func calculateTile(index int, firstByte, secondByte byte) {
	var tile types.Tile

	for i := 0; i < 8; i++ {
		mask := uint8(1 << (7 - i))
		lsb := firstByte & mask
		msb := secondByte & mask

		var color types.Color = types.Color(uint8(lsb) | uint8(msb)<<8)

		tile[index][i] = color
	}
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

// http://www.codeslinger.co.uk/pages/projects/gameboy/graphics.html
func (p *PPU) Draw(screen *ebiten.Image) {
	screen.WritePixels(p.pixels)
}

func (p *PPU) Layout(outsideWidth, outsideHeight int) (int, int) {
	return p.screenWidth, p.screenHeight
}
