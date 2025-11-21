-- +goose Up
-- Migration: Create pegawai (employee) table
-- Created: 2025-11-14

CREATE TABLE IF NOT EXISTS pegawai (
    pegawai_id VARCHAR(36) PRIMARY KEY,
    pegawai_code VARCHAR(20) NOT NULL UNIQUE,
    pegawai_name VARCHAR(100) NOT NULL,
    pegawai_gender ENUM('L','P') NOT NULL COMMENT 'L=Male, P=Female',
    pegawai_phonenumber VARCHAR(15) NOT NULL,
    pegawai_create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    pegawai_update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    pegawai_delete_at TIMESTAMP NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Indexes
CREATE INDEX idx_pegawai_code ON pegawai(pegawai_code);
CREATE INDEX idx_pegawai_delete_at ON pegawai(pegawai_delete_at);

-- +goose Down
DROP TABLE IF EXISTS pegawai;
