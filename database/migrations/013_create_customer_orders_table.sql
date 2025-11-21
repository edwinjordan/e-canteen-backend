-- +goose Up
-- Migration: Create customer_orders table
-- Created: 2025-11-14

CREATE TABLE IF NOT EXISTS customer_orders (
    order_id VARCHAR(36) PRIMARY KEY,
    order_customer_id VARCHAR(36) NOT NULL,
    order_inv_number VARCHAR(50) NOT NULL UNIQUE,
    order_address_id VARCHAR(36) NULL,
    order_delivery_type VARCHAR(20) NOT NULL COMMENT 'PICKUP, DELIVERY',
    order_total_item INT NOT NULL,
    order_subtotal DECIMAL(15,2) NOT NULL,
    order_discount DECIMAL(15,2) DEFAULT 0,
    order_total DECIMAL(15,2) NOT NULL,
    order_notes TEXT NULL,
    order_cancel_notes TEXT NULL,
    order_status INT NOT NULL COMMENT '0=Pending, 1=Processing, 2=Completed, 3=Cancelled',
    order_processed_datetime TIMESTAMP NULL,
    order_processed_by VARCHAR(36) NULL,
    order_finished_datetime TIMESTAMP NULL,
    order_finished_by VARCHAR(36) NULL,
    order_create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_customer_id) REFERENCES customers(customer_id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (order_address_id) REFERENCES customer_address(address_id) ON DELETE SET NULL ON UPDATE CASCADE,
    FOREIGN KEY (order_processed_by) REFERENCES users(user_id) ON DELETE SET NULL ON UPDATE CASCADE,
    FOREIGN KEY (order_finished_by) REFERENCES users(user_id) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Indexes
CREATE INDEX idx_order_customer_id ON customer_orders(order_customer_id);
CREATE INDEX idx_order_inv_number ON customer_orders(order_inv_number);
CREATE INDEX idx_order_status ON customer_orders(order_status);
CREATE INDEX idx_order_create_at ON customer_orders(order_create_at);

-- +goose Down
DROP TABLE IF EXISTS customer_orders;
