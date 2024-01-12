.DEFAULT_GOAL = test

BINARY = gameboy-emulator

# Build

build:
	go build -o $(BINARY) main.go
.PHONY: build

# Test
test:
	go test -v ./...
.PHONY: test