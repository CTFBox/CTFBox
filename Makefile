.PHONY: up
up:
	@docker-compose up -d --build

.PHONY: down
down:
	@docker-compose down -v

.PHONY: setup
setup:
	@go mod tidy
	@go mod download

.PHONY: build
build:
	@go build -ldflags "-s -w" .
