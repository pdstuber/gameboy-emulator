package ppu

import (
	"context"
	"log"
)

type PPU struct {
	cyclesChannel chan int
	errorChannel  chan error
}

func New() *PPU {
	return &PPU{
		cyclesChannel: make(chan int),
		errorChannel:  make(chan error),
	}
}

func (ppu *PPU) Start(ctx context.Context) error {
	log.Println("starting ppu")
	go func() {
		for {
			cycles := <-ppu.cyclesChannel
			log.Printf("elapsed cycles: %d\n", cycles)
		}
	}()

	select {
	case err := <-ppu.errorChannel:
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
