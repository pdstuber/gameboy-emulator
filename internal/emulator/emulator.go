package emulator

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/pdstuber/gameboy-emulator/internal/cpu"
	"github.com/pdstuber/gameboy-emulator/internal/memory"
	"github.com/pdstuber/gameboy-emulator/internal/ppu"
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

	ppu := ppu.New(memory)

	return &gameboy{
		cpu:             cpu.New(bootRomLoaded, memory, ppu),
		ppu:             ppu,
		shutdownChannel: make(chan interface{}),
		errorChannel:    make(chan error),
		debug:           config.Debug,
	}, nil
}

func (g *gameboy) Start(ctx context.Context) error {

	// Add ebiten Update function and add channels to notify cpu and ppu about ticks they should perform.
	go func() {
		ebiten.SetWindowSize(p.screenWidth*6, p.screenHeight*6)
		ebiten.SetWindowTitle("Gameboy Emulator")
		if err := ebiten.RunGame(p); err != nil {
			log.Fatal(err)
		}

	}()

	go func() {
		if err := g.cpu.Start(ctx); err != nil {
			g.errorChannel <- err
		}
	}()

	go func() {
		if err := g.ppu.Start(ctx); err != nil {
			g.errorChannel <- err
		}
	}()

	select {
	case err := <-g.errorChannel:
		if g.debug {
			log.Println(g.GetState())
		}
		return err
	case <-ctx.Done():
		return nil
	}
}

func (e *gameboy) Stop() {
}

func (g *gameboy) GetState() string {
	return g.cpu.GetState()
}

// http://www.codeslinger.co.uk/pages/projects/gameboy/graphics.html
func (p *PPU) Draw(screen *ebiten.Image) {
	screen.WritePixels(p.pixels)
}

func (p *PPU) Layout(outsideWidth, outsideHeight int) (int, int) {
	return p.screenWidth, p.screenHeight
}
