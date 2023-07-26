.PHONY: deps test clean build

deps:
	go get && go mod tidy

build:
	go build -o cacooclip cmd/cacooclip/main.go
