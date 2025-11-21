-- +goose Up
-- Migration: Create user_otp table
-- Created: 2025-11-14

CREATE TABLE IF NOT EXISTS user_otp (
    otp_id VARCHAR(36) PRIMARY KEY,
    otp_customer_id VARCHAR(36) NOT NULL,
    otp_number VARCHAR(6) NOT NULL,
    otp_status INT DEFAULT 0 COMMENT '0=Active, 1=Used',
    otp_expired TIMESTAMP NOT NULL,
    FOREIGN KEY (otp_customer_id) REFERENCES customers(customer_id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Indexes
CREATE INDEX idx_otp_customer_id ON user_otp(otp_customer_id);
CREATE INDEX idx_otp_status ON user_otp(otp_status);
CREATE INDEX idx_otp_expired ON user_otp(otp_expired);

-- +goose Down
DROP TABLE IF EXISTS user_otp;
