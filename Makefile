.PHONY: build run

default: run

build:
	go build cmd/server.go

run:
	go run cmd/server.go

run-prod:
	APP_ENV=production go run cmd/server.go

clean:
	rm server
