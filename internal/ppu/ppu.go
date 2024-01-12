package ppu

import (
	"context"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Color uint8

const (
	White Color = iota
	LightGrey
	DarkGrey
	Black
)

type Tile [8][8]Color

type PPU struct {
	cyclesChannel chan int
	errorChannel  chan error
	pixels        []byte
	background    [][]Tile
	screenWidth   int
	screenHeight  int
}

func New() *PPU {
	b := make([][]Tile, 20)
	for i := range b {
		b[i] = make([]Tile, 18)
	}
	return &PPU{
		cyclesChannel: make(chan int),
		errorChannel:  make(chan error),
		background:    b,
		pixels:        make([]byte, 160*144*4),
		screenWidth:   160,
		screenHeight:  144,
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
	return nil
}

// http://www.codeslinger.co.uk/pages/projects/gameboy/graphics.html
func (p *PPU) Draw(screen *ebiten.Image) {
	screen.WritePixels(p.pixels)
}

func (p *PPU) Layout(outsideWidth, outsideHeight int) (int, int) {
	return p.screenWidth, p.screenHeight
}
