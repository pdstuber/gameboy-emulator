package emulator

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pdstuber/gameboy-emulator/internal/cpu"
	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/internal/ppu"
)

const (
	screenWidth  = 256
	screenHeight = 256
)

type gameboy struct {
	cpu             *cpu.CPU
	ppu             *ppu.PPU
	shutdownChannel chan interface{}
	errorChannel    chan error
	debug           bool
}

func New(config *Config) (*gameboy, error) {
	bootRomLoaded := false

	memory := memory.New()

	if config.PathToBootRomFile != "" {
		f, err := os.Open(config.PathToBootRomFile)
		if err != nil {
			return nil, fmt.Errorf("could not open boot rom: %w", err)
		}
		bootRomData, err := io.ReadAll(f)
		if err != nil {
			return nil, fmt.Errorf("could not read boot rom data: %w", err)
		}
		if err := memory.Load(bootRomData, 0x0); err != nil {
			return nil, err
		}
		bootRomLoaded = true
	}

	if config.PathToRomFile != "" {
		f, err := os.Open(config.PathToRomFile)
		if err != nil {
			return nil, fmt.Errorf("could not open rom: %w", err)
		}
		romData, err := io.ReadAll(f)
		if err != nil {
			return nil, fmt.Errorf("could not read rom data: %w", err)
		}

		if err := memory.Load(romData, 0x100); err != nil {
			return nil, err
		}
	}

	ppu := ppu.New(memory, screenWidth*screenHeight)

	return &gameboy{
		cpu:             cpu.New(bootRomLoaded, memory, ppu),
		ppu:             ppu,
		shutdownChannel: make(chan interface{}),
		errorChannel:    make(chan error),
		debug:           config.Debug,
	}, nil
}

func (g *gameboy) Start(ctx context.Context) error {

	ebiten.SetWindowSize(screenWidth*6, screenHeight*6)
	ebiten.SetWindowTitle("Gameboy Emulator")
	if err := ebiten.RunGame(g); err != nil {
		return err
	}

	return nil
}

func (g *gameboy) Stop() {
}

func (g *gameboy) GetState() string {
	return g.cpu.GetState()
}

// http://www.codeslinger.co.uk/pages/projects/gameboy/graphics.html
func (g *gameboy) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.ppu.Pixels)
}

func (g *gameboy) Update() error {
	if err := g.cpu.Tick(); err != nil {
		return err
	}
	if err := g.ppu.Tick(); err != nil {
		return err
	}
	return nil
}

func (g *gameboy) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
