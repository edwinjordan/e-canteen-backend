-- +goose Up
-- Migration: Create temp_cart table
-- Created: 2025-11-14

CREATE TABLE IF NOT EXISTS temp_cart (
    temp_cart_id VARCHAR(36) PRIMARY KEY,
    temp_cart_order_id VARCHAR(36) NULL,
    temp_cart_product_varian_id VARCHAR(36) NOT NULL,
    temp_cart_user_id VARCHAR(36) NOT NULL,
    temp_cart_qty INT NOT NULL,
    FOREIGN KEY (temp_cart_product_varian_id) REFERENCES product_varians(product_varian_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (temp_cart_user_id) REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Indexes
CREATE INDEX idx_temp_cart_user_id ON temp_cart(temp_cart_user_id);
CREATE INDEX idx_temp_cart_order_id ON temp_cart(temp_cart_order_id);
CREATE INDEX idx_temp_cart_product_varian_id ON temp_cart(temp_cart_product_varian_id);

-- +goose Down
DROP TABLE IF EXISTS temp_cart;
