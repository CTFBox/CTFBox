.PHONY: setup
setup:
	go mod tidy
	go mod download

.PHONY: build
build:
	go build -ldflags "-s -w" .
