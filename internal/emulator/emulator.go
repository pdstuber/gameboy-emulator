package emulator

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pdstuber/gameboy-emulator/internal/cpu"
	"github.com/pdstuber/gameboy-emulator/internal/memory"
)

type gameboy struct {
	cpu             *cpu.CPU
	shutdownChannel chan interface{}
	errorChannel    chan error
	debug           bool
}

func New(config *Config) (*gameboy, error) {
	f, err := os.Open(config.PathToBootRomFile)
	if err != nil {
		return nil, fmt.Errorf("could not open boot rom: %w", err)
	}

	bootRomData, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("could not read boot rom data: %w", err)
	}

	f, err = os.Open(config.PathToRomFile)
	if err != nil {
		return nil, fmt.Errorf("could not open rom: %w", err)
	}

	romData, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("could not read rom data: %w", err)
	}

	memory := memory.New()
	if err := memory.Load(bootRomData, 0x0000); err != nil {
		return nil, err
	}
	if err := memory.Load(romData, 0x0100); err != nil {
		return nil, err
	}
	return &gameboy{
		cpu:             cpu.New(memory),
		shutdownChannel: make(chan interface{}),
		errorChannel:    make(chan error),
	}, nil
}

func (g *gameboy) Start(ctx context.Context) error {
	go func() {
		for {
			if err := g.tick(); err != nil {
				g.errorChannel <- err
			}
		}
	}()

	log.Println("emulator started")

	select {
	case <-g.shutdownChannel:
		return nil
	case err := <-g.errorChannel:
		return err
	case <-ctx.Done():
		return nil
	}
}

func (e *gameboy) Stop() {
	close(e.shutdownChannel)
}

func (g *gameboy) tick() error {
	return g.cpu.FetchAndExecuteNextInstruction()
}
