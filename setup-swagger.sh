#!/bin/bash
# E-Canteen Cashier API - Swagger Setup Script
# This script installs swag CLI and generates Swagger documentation

echo "=========================================="
echo "E-Canteen Swagger Documentation Setup"
echo "=========================================="
echo ""

# Check if Go is installed
echo "Checking Go installation..."
if ! command -v go &> /dev/null; then
    echo "✗ Go is not installed or not in PATH"
    echo "Please install Go from https://golang.org/dl/"
    exit 1
fi
echo "✓ Go installed: $(go version)"
echo ""

# Install swag CLI tool
echo "Installing swag CLI tool..."
go install github.com/swaggo/swag/cmd/swag@latest

if [ $? -eq 0 ]; then
    echo "✓ swag CLI installed successfully"
else
    echo "✗ Failed to install swag CLI"
    exit 1
fi
echo ""

# Download dependencies
echo "Downloading Go dependencies..."
go mod download
go mod tidy

if [ $? -eq 0 ]; then
    echo "✓ Dependencies downloaded"
else
    echo "✗ Failed to download dependencies"
    exit 1
fi
echo ""

# Generate Swagger documentation
echo "Generating Swagger documentation..."
swag init -g main.go --parseDependency --parseInternal

if [ $? -eq 0 ]; then
    echo "✓ Swagger documentation generated successfully!"
else
    echo "✗ Failed to generate Swagger documentation"
    exit 1
fi
echo ""

echo "=========================================="
echo "Setup Complete!"
echo "=========================================="
echo ""
echo "To start the API server, run:"
echo "  go run main.go"
echo ""
echo "Access Swagger documentation at:"
echo "  http://127.0.0.1:3000/swagger/index.html"
echo ""
