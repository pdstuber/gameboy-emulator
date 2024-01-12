.DEFAULT_GOAL = test

BINARY = gameboy-emulator

# Build

build: $(BINARY)
.PHONY: build

# Test
test: build
	go test -v ./...
.PHONY: test

$(BINARY): 
	go build -o $@ ./cmd/golangci-lint