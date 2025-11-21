# E-Canteen Cashier API - Swagger Setup Script for Windows
# This script installs swag CLI and generates Swagger documentation

Write-Host "==========================================" -ForegroundColor Cyan
Write-Host "E-Canteen Swagger Documentation Setup" -ForegroundColor Cyan
Write-Host "==========================================" -ForegroundColor Cyan
Write-Host ""

# Check if Go is installed
Write-Host "Checking Go installation..." -ForegroundColor Yellow
$goVersion = go version 2>&1
if ($LASTEXITCODE -ne 0) {
    Write-Host "✗ Go is not installed or not in PATH" -ForegroundColor Red
    Write-Host "Please install Go from https://golang.org/dl/" -ForegroundColor Yellow
    exit 1
}
Write-Host "✓ Go installed: $goVersion" -ForegroundColor Green
Write-Host ""

# Install swag CLI tool
Write-Host "Installing swag CLI tool..." -ForegroundColor Yellow
go install github.com/swaggo/swag/cmd/swag@latest

if ($LASTEXITCODE -eq 0) {
    Write-Host "✓ swag CLI installed successfully" -ForegroundColor Green
} else {
    Write-Host "✗ Failed to install swag CLI" -ForegroundColor Red
    exit 1
}
Write-Host ""

# Download dependencies
Write-Host "Downloading Go dependencies..." -ForegroundColor Yellow
go mod download
go mod tidy

if ($LASTEXITCODE -eq 0) {
    Write-Host "✓ Dependencies downloaded" -ForegroundColor Green
} else {
    Write-Host "✗ Failed to download dependencies" -ForegroundColor Red
    exit 1
}
Write-Host ""

# Generate Swagger documentation
Write-Host "Generating Swagger documentation..." -ForegroundColor Yellow
swag init -g main.go --parseDependency --parseInternal

if ($LASTEXITCODE -eq 0) {
    Write-Host "✓ Swagger documentation generated successfully!" -ForegroundColor Green
} else {
    Write-Host "✗ Failed to generate Swagger documentation" -ForegroundColor Red
    exit 1
}
Write-Host ""

Write-Host "==========================================" -ForegroundColor Cyan
Write-Host "Setup Complete!" -ForegroundColor Green
Write-Host "==========================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "To start the API server, run:" -ForegroundColor Yellow
Write-Host "  go run main.go" -ForegroundColor White
Write-Host ""
Write-Host "Access Swagger documentation at:" -ForegroundColor Yellow
Write-Host "  http://127.0.0.1:3000/swagger/index.html" -ForegroundColor Cyan
Write-Host ""
