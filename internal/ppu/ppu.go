package ppu

import (
	"context"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/pkg/types"
)

type Color uint8

const (
	White Color = iota
	LightGrey
	DarkGrey
	Black
)

type Tile [8][8]Color

func (c Color) ToStandardColor() color.Color {

	switch c {
	case White:
		return color.White
	case Black:
		return color.Black
	case LightGrey:
		return color.Gray16{0x0000}
	case DarkGrey:
		return color.Gray16{}
	}

	return color.Opaque
}

const (
	vramBegin = 0x8000
	vramEnd   = 0x9FFF
)

type PPU struct {
	cyclesChannel chan int
	errorChannel  chan error
	memory        *memory.Memory
	pixels        []byte
	tiles         []Tile
	screenWidth   int
	screenHeight  int
}

func New(memory *memory.Memory) *PPU {
	b := make([][]Tile, 32)
	for i := range b {
		b[i] = make([]Tile, 32)
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
	for i := vramBegin; i < vramEnd; i += 2 {
		addressFirstTileByte := types.Address(i)
		addressSecondTileByte := types.Address(i + 1)
		byte1 := p.memory.Read(addressFirstTileByte)
		byte2 := p.memory.Read(addressSecondTileByte)

		tileIndex := int((i - vramBegin) / 16)
		rowIndex := int(((i - vramBegin) % 16) / 2)

		for j := 0; j < 8; j++ {

			mask := uint8(1 << (7 - j))
			lsb := byte1 & mask
			msb := byte2 & mask

			var value Color

			if lsb != 0 {
				if msb != 0 {
					value = Black
				} else {
					value = LightGrey
				}
			} else {
				if msb != 0 {
					value = DarkGrey
				} else {
					value = White
				}

			}
			p.tiles[tileIndex][rowIndex][j] = value
		}

	}
	fmt.Println(p.tiles)
	return nil
}

func tilesToPixels(tiles []Tile) []byte {
	for i := range tiles {
		tile := tiles[i]
		zs
	}
}

// http://www.codeslinger.co.uk/pages/projects/gameboy/graphics.html
func (p *PPU) Draw(screen *ebiten.Image) {
	screen.WritePixels(p.pixels)
}

func (p *PPU) Layout(outsideWidth, outsideHeight int) (int, int) {
	return p.screenWidth, p.screenHeight
}
