# Database Seeder Script for Windows
# This script runs all seeder files in order

Write-Host "==========================================" -ForegroundColor Cyan
Write-Host "E-Canteen Database Seeder" -ForegroundColor Cyan
Write-Host "==========================================" -ForegroundColor Cyan
Write-Host ""

# Database configuration from environment variables or defaults
$DB_HOST = if ($env:DB_HOST) { $env:DB_HOST } else { "localhost" }
$DB_PORT = if ($env:DB_PORT) { $env:DB_PORT } else { "3306" }
$DB_USER = if ($env:DB_USER) { $env:DB_USER } else { "root" }
$DB_PASS = if ($env:DB_PASS) { $env:DB_PASS } else { "" }
$DB_NAME = if ($env:DB_NAME) { $env:DB_NAME } else { "ecanteen_db" }

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

Write-Host "Running seeders..." -ForegroundColor Yellow
Write-Host ""

# Get seeders directory
$SCRIPT_DIR = Split-Path -Parent $MyInvocation.MyCommand.Path
$SEEDERS_DIR = Join-Path $SCRIPT_DIR "seeders"

# Counter for seeders
$SUCCESS_COUNT = 0
$FAIL_COUNT = 0

# Get all seeder files sorted by name
$seederFiles = Get-ChildItem -Path $SEEDERS_DIR -Filter "*.sql" | Sort-Object Name

foreach ($file in $seederFiles) {
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
Write-Host "Seeder Summary" -ForegroundColor Cyan
Write-Host "==========================================" -ForegroundColor Cyan
Write-Host "Successful: $SUCCESS_COUNT" -ForegroundColor Green
Write-Host "Failed: $FAIL_COUNT" -ForegroundColor $(if ($FAIL_COUNT -eq 0) { "Green" } else { "Red" })
Write-Host ""

if ($FAIL_COUNT -eq 0) {
    Write-Host "All seeders completed successfully! ✓" -ForegroundColor Green
    exit 0
} else {
    Write-Host "Some seeders failed! ✗" -ForegroundColor Red
    exit 1
}
