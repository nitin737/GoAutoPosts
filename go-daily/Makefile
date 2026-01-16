.PHONY: help build run test validate clean install

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install dependencies
	go mod download
	go mod tidy

build: ## Build the application
	go build -o bin/publisher cmd/publisher/main.go

run: ## Run the application
	go run cmd/publisher/main.go

test: ## Run tests
	go test -v ./...

test-setup: ## Run comprehensive setup validation
	@bash scripts/test.sh

validate: ## Validate libraries.json
	go run scripts/validate_data.go data/libraries.json

clean: ## Clean build artifacts
	rm -rf bin/
	rm -f data/posted.json
	echo "[]" > data/posted.json

fmt: ## Format code
	go fmt ./...

lint: ## Run linter
	golangci-lint run

dev: ## Run in development mode
	ENVIRONMENT=development go run cmd/publisher/main.go
