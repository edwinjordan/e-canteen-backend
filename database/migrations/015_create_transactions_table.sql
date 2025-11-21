-- +goose Up
-- Migration: Create transactions table
-- Created: 2025-11-14

CREATE TABLE IF NOT EXISTS transactions (
    trans_id VARCHAR(36) PRIMARY KEY,
    trans_user_id VARCHAR(36) NOT NULL,
    trans_customer_id VARCHAR(36) NULL,
    trans_order_id VARCHAR(36) NULL,
    trans_invoice VARCHAR(50) NOT NULL UNIQUE,
    trans_qty_total INT NOT NULL,
    trans_product_total INT NOT NULL,
    trans_subtotal DECIMAL(15,2) NOT NULL,
    trans_discount DECIMAL(15,2) DEFAULT 0,
    trans_total DECIMAL(15,2) NOT NULL,
    trans_received_total DECIMAL(15,2) NOT NULL,
    trans_refund_total DECIMAL(15,2) DEFAULT 0,
    trans_status INT DEFAULT 1 COMMENT '0=Cancelled, 1=Completed',
    trans_create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (trans_user_id) REFERENCES users(user_id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (trans_customer_id) REFERENCES customers(customer_id) ON DELETE SET NULL ON UPDATE CASCADE,
    FOREIGN KEY (trans_order_id) REFERENCES customer_orders(order_id) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Indexes
CREATE INDEX idx_trans_user_id ON transactions(trans_user_id);
CREATE INDEX idx_trans_customer_id ON transactions(trans_customer_id);
CREATE INDEX idx_trans_invoice ON transactions(trans_invoice);
CREATE INDEX idx_trans_create_at ON transactions(trans_create_at);
CREATE INDEX idx_trans_status ON transactions(trans_status);

-- +goose Down
DROP TABLE IF EXISTS transactions;
