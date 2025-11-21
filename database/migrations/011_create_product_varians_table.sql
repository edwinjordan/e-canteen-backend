-- +goose Up
-- Migration: Create product_varians table
-- Created: 2025-11-14

CREATE TABLE IF NOT EXISTS product_varians (
    product_varian_id VARCHAR(36) PRIMARY KEY,
    product_id VARCHAR(36) NOT NULL,
    varian_id VARCHAR(36) NOT NULL,
    varian_name VARCHAR(100) NOT NULL,
    product_varian_price INT NOT NULL,
    product_varian_qty_booth INT DEFAULT 0,
    product_varian_qty_warehouse VARCHAR(20) DEFAULT '0',
    product_varian_qty_left INT DEFAULT 0,
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Indexes
CREATE INDEX idx_product_varian_product_id ON product_varians(product_id);
CREATE INDEX idx_product_varian_varian_id ON product_varians(varian_id);

-- +goose Down
DROP TABLE IF EXISTS product_varians;
