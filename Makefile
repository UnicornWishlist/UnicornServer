.PHONY: build run

default: run

build:
	go build cmd/server.go

run:
	go run cmd/server.go

