-- +goose Up
-- Migration: Create products table
-- Created: 2025-11-14

CREATE TABLE IF NOT EXISTS products (
    product_id VARCHAR(36) PRIMARY KEY,
    product_code VARCHAR(20) NOT NULL UNIQUE,
    product_name VARCHAR(100) NOT NULL,
    product_category_id VARCHAR(36) NOT NULL,
    product_desc TEXT NULL,
    product_create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    product_update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    product_delete_at TIMESTAMP NULL DEFAULT NULL,
    product_photo VARCHAR(255) NULL,
    FOREIGN KEY (product_category_id) REFERENCES categories(category_id) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Indexes
CREATE INDEX idx_product_code ON products(product_code);
CREATE INDEX idx_product_category_id ON products(product_category_id);
CREATE INDEX idx_product_delete_at ON products(product_delete_at);

-- +goose Down
DROP TABLE IF EXISTS products;
