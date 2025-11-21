#!/bin/bash
# Database Seeder Script
# This script runs all seeder files in order

echo "=========================================="
echo "E-Canteen Database Seeder"
echo "=========================================="
echo ""

# Database configuration
DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-3306}"
DB_USER="${DB_USER:-root}"
DB_PASS="${DB_PASS:-}"
DB_NAME="${DB_NAME:-ecanteen_db}"

echo "Database: $DB_NAME"
echo "Host: $DB_HOST:$DB_PORT"
echo "User: $DB_USER"
echo ""

echo "Running seeders..."
echo ""

# Get the directory where the script is located
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
SEEDERS_DIR="$SCRIPT_DIR/seeders"

# Counter for seeders
SUCCESS_COUNT=0
FAIL_COUNT=0

# Run each seeder file in order
for seeder_file in "$SEEDERS_DIR"/*.sql; do
    if [ -f "$seeder_file" ]; then
        filename=$(basename "$seeder_file")
        echo "Running: $filename"
        
        mysql -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -p"$DB_PASS" "$DB_NAME" < "$seeder_file"
        
        if [ $? -eq 0 ]; then
            echo "  ✓ Success"
            ((SUCCESS_COUNT++))
        else
            echo "  ✗ Failed"
            ((FAIL_COUNT++))
        fi
        echo ""
    fi
done

echo "=========================================="
echo "Seeder Summary"
echo "=========================================="
echo "Successful: $SUCCESS_COUNT"
echo "Failed: $FAIL_COUNT"
echo ""

if [ $FAIL_COUNT -eq 0 ]; then
    echo "All seeders completed successfully! ✓"
    exit 0
else
    echo "Some seeders failed! ✗"
    exit 1
fi
