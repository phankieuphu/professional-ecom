# Makefile for building and running a Go application

APP_NAME := gateway-service
CMD_PATH := ./cmd
PROTO_DIR=proto

.PHONY: all build run clean test fmt vet lint tidy

all: build

build:
	go build -o bin/$(APP_NAME) $(CMD_PATH)

run: build
	./bin/$(APP_NAME)

clean:
	rm -rf bin/

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run

tidy:
	go mod tidy
generate:
	buf generate
