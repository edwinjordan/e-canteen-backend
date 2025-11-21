-- +goose Up
-- Migration: Create customers table
-- Created: 2025-11-14

CREATE TABLE IF NOT EXISTS customers (
    customer_id VARCHAR(36) PRIMARY KEY,
    customer_code VARCHAR(20) NOT NULL UNIQUE,
    customer_name VARCHAR(100) NOT NULL,
    customer_gender ENUM('L','P') NOT NULL COMMENT 'L=Male, P=Female',
    customer_phonenumber VARCHAR(15) NOT NULL,
    customer_email VARCHAR(100) NOT NULL UNIQUE,
    customer_dob DATE NULL,
    customer_password VARCHAR(255) NOT NULL,
    customer_profile_pic VARCHAR(255) NULL,
    customer_class VARCHAR(20) NULL,
    customer_major_id VARCHAR(36) NULL,
    customer_profile_pic_path VARCHAR(255) NULL,
    customer_status INT DEFAULT 1 COMMENT '0=Inactive, 1=Active',
    customer_last_status INT DEFAULT 1,
    customer_create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    customer_update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_major_id) REFERENCES majors(major_id) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Indexes
CREATE INDEX idx_customer_code ON customers(customer_code);
CREATE INDEX idx_customer_email ON customers(customer_email);
CREATE INDEX idx_customer_major_id ON customers(customer_major_id);
CREATE INDEX idx_customer_status ON customers(customer_status);

-- +goose Down
DROP TABLE IF EXISTS customers;
