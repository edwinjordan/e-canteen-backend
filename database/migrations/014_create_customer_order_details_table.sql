-- +goose Up
-- Migration: Create customer_order_details table
-- Created: 2025-11-14

CREATE TABLE IF NOT EXISTS customer_order_details (
    order_detail_id VARCHAR(36) PRIMARY KEY,
    order_detail_parent_id VARCHAR(36) NOT NULL,
    order_detail_product_varian_id VARCHAR(36) NOT NULL,
    order_detail_qty INT NOT NULL,
    order_detail_price DECIMAL(15,2) NOT NULL,
    order_detail_subtotal DECIMAL(15,2) NOT NULL,
    FOREIGN KEY (order_detail_parent_id) REFERENCES customer_orders(order_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (order_detail_product_varian_id) REFERENCES product_varians(product_varian_id) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Indexes
CREATE INDEX idx_order_detail_parent_id ON customer_order_details(order_detail_parent_id);
CREATE INDEX idx_order_detail_product_varian_id ON customer_order_details(order_detail_product_varian_id);

-- +goose Down
DROP TABLE IF EXISTS customer_order_details;
