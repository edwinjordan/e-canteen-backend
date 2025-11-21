-- +goose Up
-- Migration: Create majors table
-- Created: 2025-11-14

CREATE TABLE IF NOT EXISTS majors (
    major_id VARCHAR(36) PRIMARY KEY,
    major_name VARCHAR(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +goose Down
DROP TABLE IF EXISTS majors;
