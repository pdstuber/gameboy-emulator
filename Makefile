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

lint:
	docker run -t --rm -v $(PWD):/app -w /app golangci/golangci-lint:latest golangci-lint run --enable-all --fix