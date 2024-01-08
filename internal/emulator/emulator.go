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
	f, err := os.Open(config.PathToRomFile)
	if err != nil {
		return nil, fmt.Errorf("could not open rom: %w", err)
	}

	romData, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("could not read rom data: %w", err)
	}

	memory := memory.New()

	if err := memory.Load(romData, 0x100); err != nil {
		return nil, err
	}

	return &gameboy{
		cpu:             cpu.New(memory),
		shutdownChannel: make(chan interface{}),
		errorChannel:    make(chan error),
		debug:           config.Debug,
	}, nil
}

func (g *gameboy) Start(ctx context.Context) error {
	go func() {
		for {
			if err := g.tick(); err != nil {
				g.errorChannel <- err
				break
			}
		}
	}()

	log.Println("emulator started")

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

func (g *gameboy) tick() error {
	if err := g.cpu.FetchAndExecuteNextInstruction(); err != nil {
		return err
	}
	return nil
}

func (g *gameboy) GetState() string {
	return g.cpu.GetState()
}
