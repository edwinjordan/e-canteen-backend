# Makefile for E-Canteen Cashier API

.PHONY: help install swagger run build test clean

help: ## Display this help message
	@echo "E-Canteen Cashier API - Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

install: ## Install dependencies
	@echo "Installing Go dependencies..."
	go mod download
	go mod tidy
	@echo "Installing swag CLI tool..."
	go install github.com/swaggo/swag/cmd/swag@latest
	@echo "Dependencies installed successfully!"

swagger: ## Generate Swagger documentation
	@echo "Generating Swagger documentation..."
	swag init -g main.go --parseDependency --parseInternal
	@echo "Swagger docs generated successfully!"
	@echo "Access docs at: http://127.0.0.1:3000/swagger/index.html"

run: ## Run the application
	@echo "Starting E-Canteen Cashier API..."
	go run main.go

dev: swagger run ## Generate swagger and run the application

build: ## Build the application
	@echo "Building E-Canteen Cashier API..."
	go build -o bin/ecanteen-api main.go
	@echo "Build complete! Binary: bin/ecanteen-api"

test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	rm -rf docs/swagger.*
	@echo "Clean complete!"

migrate: ## Run database migrations
	@echo "Running database migrations..."
	cd database && ./migrate.sh

seed: ## Run database seeders
	@echo "Running database seeders..."
	cd database && ./seed.sh

db-setup: migrate seed ## Setup database (migrate + seed)
	@echo "Database setup complete!"

# Windows specific commands
swagger-win: ## Generate Swagger documentation (Windows)
	@echo "Generating Swagger documentation..."
	swag init -g main.go --parseDependency --parseInternal
	@echo "Swagger docs generated successfully!"

migrate-win: ## Run database migrations (Windows)
	@echo "Running database migrations..."
	powershell -ExecutionPolicy Bypass -File database\migrate.ps1

seed-win: ## Run database seeders (Windows)
	@echo "Running database seeders..."
	powershell -ExecutionPolicy Bypass -File database\seed.ps1

db-setup-win: migrate-win seed-win ## Setup database (Windows)
	@echo "Database setup complete!"
