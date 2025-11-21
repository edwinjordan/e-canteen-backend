-- +goose Up
-- Migration: Create stock_booth table
-- Created: 2025-11-14

CREATE TABLE IF NOT EXISTS stock_booth (
    product_stok_id VARCHAR(36) PRIMARY KEY,
    product_stok_product_varian_id VARCHAR(36) NOT NULL,
    product_stok_first_qty INT NOT NULL,
    product_stok_qty INT NOT NULL,
    product_stok_last_qty INT NOT NULL,
    product_stok_jenis VARCHAR(20) NOT NULL COMMENT 'IN, OUT, ADJUSTMENT',
    product_stok_datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    product_stok_pegawai_id VARCHAR(36) NOT NULL,
    FOREIGN KEY (product_stok_product_varian_id) REFERENCES product_varians(product_varian_id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (product_stok_pegawai_id) REFERENCES pegawai(pegawai_id) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Indexes
CREATE INDEX idx_stock_booth_product_varian_id ON stock_booth(product_stok_product_varian_id);
CREATE INDEX idx_stock_booth_pegawai_id ON stock_booth(product_stok_pegawai_id);
CREATE INDEX idx_stock_booth_datetime ON stock_booth(product_stok_datetime);
CREATE INDEX idx_stock_booth_jenis ON stock_booth(product_stok_jenis);

-- +goose Down
DROP TABLE IF EXISTS stock_booth;
