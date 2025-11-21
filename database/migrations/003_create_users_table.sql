-- +goose Up
-- Migration: Create users table
-- Created: 2025-11-14

CREATE TABLE IF NOT EXISTS users (
    user_id VARCHAR(36) PRIMARY KEY,
    user_name VARCHAR(100) NOT NULL,
    user_email VARCHAR(100) NOT NULL UNIQUE,
    user_password VARCHAR(255) NOT NULL,
    user_pegawai_id VARCHAR(36) NOT NULL,
    user_has_mobile_access INT DEFAULT 0 COMMENT '0=No, 1=Yes',
    user_role_id VARCHAR(36) NOT NULL,
    user_create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_pegawai_id) REFERENCES pegawai(pegawai_id) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (user_role_id) REFERENCES roles(role_id) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Indexes
CREATE INDEX idx_user_email ON users(user_email);
CREATE INDEX idx_user_pegawai_id ON users(user_pegawai_id);
CREATE INDEX idx_user_role_id ON users(user_role_id);

-- +goose Down
DROP TABLE IF EXISTS users;
