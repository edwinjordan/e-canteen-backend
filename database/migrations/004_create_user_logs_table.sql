-- +goose Up
-- Migration: Create user_logs table
-- Created: 2025-11-14

CREATE TABLE IF NOT EXISTS user_logs (
    log_user_id VARCHAR(36) PRIMARY KEY,
    log_user_user_id VARCHAR(36) NOT NULL,
    log_user_token TEXT NOT NULL,
    log_user_metadata TEXT NULL,
    log_user_login_date TIMESTAMP NOT NULL,
    log_user_logout_date TIMESTAMP NULL,
    FOREIGN KEY (log_user_user_id) REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Indexes
CREATE INDEX idx_log_user_user_id ON user_logs(log_user_user_id);
CREATE INDEX idx_log_user_login_date ON user_logs(log_user_login_date);

-- +goose Down
DROP TABLE IF EXISTS user_logs;
