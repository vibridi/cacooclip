.PHONY: deps test clean build

deps:
	go get && go mod tidy

install:
	go install github.com/vibridi/cacooclip/cmd/cacooclip@latest
