all: test build run

run: build
	@./bin/monkey

build: $(shell find cmd internal -name "*.go")
	@go build -o bin/monkey cmd/main.go

test:
	@go test ./... -v
