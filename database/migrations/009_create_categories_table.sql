-- +goose Up
-- Migration: Create categories table
-- Created: 2025-11-14

CREATE TABLE IF NOT EXISTS categories (
    category_id VARCHAR(36) PRIMARY KEY,
    category_name VARCHAR(100) NOT NULL,
    category_delete_at TIMESTAMP NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Indexes
CREATE INDEX idx_category_delete_at ON categories(category_delete_at);

-- +goose Down
DROP TABLE IF EXISTS categories;
