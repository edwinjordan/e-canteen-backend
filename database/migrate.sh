#!/bin/bash
# Database Migration Script
# This script runs all migration files in order

echo "=========================================="
echo "E-Canteen Database Migration"
echo "=========================================="
echo ""

# Database configuration
DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-3306}"
DB_USER="${DB_USER:-root}"
DB_PASS="${DB_PASS:-}"
DB_NAME="${DB_NAME:-e-canteen_new}"

echo "Database: $DB_NAME"
echo "Host: $DB_HOST:$DB_PORT"
echo "User: $DB_USER"
echo ""

# Create database if not exists
echo "Creating database if not exists..."
mysql -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -p"$DB_PASS" -e "CREATE DATABASE IF NOT EXISTS $DB_NAME CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

if [ $? -eq 0 ]; then
    echo "✓ Database ready"
else
    echo "✗ Failed to create database"
    exit 1
fi

echo ""
echo "Running migrations..."
echo ""

# Get the directory where the script is located
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
MIGRATIONS_DIR="$SCRIPT_DIR/migrations"

# Counter for migrations
SUCCESS_COUNT=0
FAIL_COUNT=0

# Run each migration file in order
for migration_file in "$MIGRATIONS_DIR"/*.sql; do
    if [ -f "$migration_file" ]; then
        filename=$(basename "$migration_file")
        echo "Running: $filename"
        
        mysql -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -p"$DB_PASS" "$DB_NAME" < "$migration_file"
        
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
echo "Migration Summary"
echo "=========================================="
echo "Successful: $SUCCESS_COUNT"
echo "Failed: $FAIL_COUNT"
echo ""

if [ $FAIL_COUNT -eq 0 ]; then
    echo "All migrations completed successfully! ✓"
    exit 0
else
    echo "Some migrations failed! ✗"
    exit 1
fi
