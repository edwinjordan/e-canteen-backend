-- +goose Up
-- Migration: Create transaction_details table
-- Created: 2025-11-14

CREATE TABLE IF NOT EXISTS transaction_details (
    trans_detail_id VARCHAR(36) PRIMARY KEY,
    trans_detail_parent_id VARCHAR(36) NOT NULL,
    trans_detail_product_varian_id VARCHAR(36) NOT NULL,
    trans_detail_qty INT NOT NULL,
    trans_detail_price DECIMAL(15,2) NOT NULL,
    trans_detail_subtotal DECIMAL(15,2) NOT NULL,
    FOREIGN KEY (trans_detail_parent_id) REFERENCES transactions(trans_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (trans_detail_product_varian_id) REFERENCES product_varians(product_varian_id) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Indexes
CREATE INDEX idx_trans_detail_parent_id ON transaction_details(trans_detail_parent_id);
CREATE INDEX idx_trans_detail_product_varian_id ON transaction_details(trans_detail_product_varian_id);

-- +goose Down
DROP TABLE IF EXISTS transaction_details;
