.PHONY: dev build test lint fmt

dev:
	go run .

build:
	go build -o server .

test:
	go test -race ./...

lint:
	golangci-lint run

fmt:
	gofmt -w .
