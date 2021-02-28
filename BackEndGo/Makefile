# Grab local env values and export for local dev
include .env
export


lint:
	go fmt ./...
	golint ./...
	go vet ./...

mod:
	go mod download
	go mod verify

build:
	go build -o bin/main main.go

run:
	go run main.go

test:
	go test ./...