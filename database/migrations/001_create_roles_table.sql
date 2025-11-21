-- +goose Up
-- Migration: Create roles table
-- Created: 2025-11-14

CREATE TABLE IF NOT EXISTS roles (
    role_id VARCHAR(36) PRIMARY KEY,
    role_name VARCHAR(50) NOT NULL,
    role_code VARCHAR(20) NOT NULL UNIQUE,
    role_create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    role_update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Indexes
CREATE INDEX idx_role_code ON roles(role_code);

-- +goose Down
DROP TABLE IF EXISTS roles;
