-- +goose Up
-- Migration: Create version_admin table
-- Created: 2025-11-14

CREATE TABLE IF NOT EXISTS version_admin (
    version_id INT PRIMARY KEY AUTO_INCREMENT,
    version_number VARCHAR(20) NOT NULL,
    version_code INT NOT NULL,
    version_chagelog TEXT NULL,
    version_datetime TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Indexes
CREATE INDEX idx_version_admin_code ON version_admin(version_code);

-- +goose Down
DROP TABLE IF EXISTS version_admin;
