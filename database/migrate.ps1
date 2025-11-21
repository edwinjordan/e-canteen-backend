# Database Migration Script for Windows
# This script runs all migration files in order

Write-Host "==========================================" -ForegroundColor Cyan
Write-Host "E-Canteen Database Migration" -ForegroundColor Cyan
Write-Host "==========================================" -ForegroundColor Cyan
Write-Host ""

# Database configuration from environment variables or defaults
$DB_HOST = if ($env:DB_HOST) { $env:DB_HOST } else { "localhost" }
$DB_PORT = if ($env:DB_PORT) { $env:DB_PORT } else { "3306" }
$DB_USER = if ($env:DB_USER) { $env:DB_USER } else { "root" }
$DB_PASS = if ($env:DB_PASS) { $env:DB_PASS } else { "" }
$DB_NAME = if ($env:DB_NAME) { $env:DB_NAME } else { "e-canteen_new" }

Write-Host "Database: $DB_NAME" -ForegroundColor Yellow
Write-Host "Host: ${DB_HOST}:${DB_PORT}" -ForegroundColor Yellow
Write-Host "User: $DB_USER" -ForegroundColor Yellow
Write-Host ""

# MySQL command path (adjust if needed)
$MYSQL_CMD = "mysql"

# Build connection string
$connStr = "-h$DB_HOST -P$DB_PORT -u$DB_USER"
if ($DB_PASS) {
    $connStr += " -p$DB_PASS"
}

# Create database if not exists
Write-Host "Creating database if not exists..." -ForegroundColor Yellow
$createDbCmd = "CREATE DATABASE IF NOT EXISTS $DB_NAME CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
echo $createDbCmd | & $MYSQL_CMD $connStr 2>&1 | Out-Null

if ($LASTEXITCODE -eq 0) {
    Write-Host "✓ Database ready" -ForegroundColor Green
} else {
    Write-Host "✗ Failed to create database" -ForegroundColor Red
    exit 1
}

Write-Host ""
Write-Host "Running migrations..." -ForegroundColor Yellow
Write-Host ""

# Get migrations directory
$SCRIPT_DIR = Split-Path -Parent $MyInvocation.MyCommand.Path
$MIGRATIONS_DIR = Join-Path $SCRIPT_DIR "migrations"

# Counter for migrations
$SUCCESS_COUNT = 0
$FAIL_COUNT = 0

# Get all migration files sorted by name
$migrationFiles = Get-ChildItem -Path $MIGRATIONS_DIR -Filter "*.sql" | Sort-Object Name

foreach ($file in $migrationFiles) {
    Write-Host "Running: $($file.Name)" -ForegroundColor Cyan
    
    Get-Content $file.FullName | & $MYSQL_CMD $connStr $DB_NAME 2>&1 | Out-Null
    
    if ($LASTEXITCODE -eq 0) {
        Write-Host "  ✓ Success" -ForegroundColor Green
        $SUCCESS_COUNT++
    } else {
        Write-Host "  ✗ Failed" -ForegroundColor Red
        $FAIL_COUNT++
    }
    Write-Host ""
}

Write-Host "==========================================" -ForegroundColor Cyan
Write-Host "Migration Summary" -ForegroundColor Cyan
Write-Host "==========================================" -ForegroundColor Cyan
Write-Host "Successful: $SUCCESS_COUNT" -ForegroundColor Green
Write-Host "Failed: $FAIL_COUNT" -ForegroundColor $(if ($FAIL_COUNT -eq 0) { "Green" } else { "Red" })
Write-Host ""

if ($FAIL_COUNT -eq 0) {
    Write-Host "All migrations completed successfully! ✓" -ForegroundColor Green
    exit 0
} else {
    Write-Host "Some migrations failed! ✗" -ForegroundColor Red
    exit 1
}
