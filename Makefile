include .env

build:
	go build -o ${BINARY} ./cmd/api