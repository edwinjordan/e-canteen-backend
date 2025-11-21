-- +goose Up
-- Migration: Create customer_address table
-- Created: 2025-11-14

CREATE TABLE IF NOT EXISTS customer_address (
    address_id VARCHAR(36) PRIMARY KEY,
    address_customer_id VARCHAR(36) NOT NULL,
    address_text TEXT NOT NULL,
    address_name VARCHAR(100) NOT NULL,
    address_province_id VARCHAR(10) NOT NULL,
    address_province VARCHAR(100) NOT NULL,
    address_city_id VARCHAR(10) NOT NULL,
    address_city VARCHAR(100) NOT NULL,
    address_district_id VARCHAR(10) NOT NULL,
    address_district VARCHAR(100) NOT NULL,
    address_village_id VARCHAR(10) NOT NULL,
    address_village VARCHAR(100) NOT NULL,
    address_postal_code VARCHAR(10) NULL,
    address_main INT DEFAULT 0 COMMENT '0=No, 1=Yes',
    address_create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    address_update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (address_customer_id) REFERENCES customers(customer_id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Indexes
CREATE INDEX idx_address_customer_id ON customer_address(address_customer_id);
CREATE INDEX idx_address_main ON customer_address(address_main);

-- +goose Down
DROP TABLE IF EXISTS customer_address;
