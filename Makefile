.PHONY: build

build:
	go build -o sortify cmd/main/main.go

.DEFAULT_GOAL := build