-- +goose Up
-- Migration: Create v_ms_product view
-- Created: 2025-11-16
-- Description: View to join products with categories

-- Drop view if exists
DROP VIEW IF EXISTS v_ms_product;

-- Create view v_ms_product
CREATE VIEW v_ms_product AS
SELECT
    p.product_id,
    p.product_code,
    p.product_name,
    p.product_category_id,
    p.product_desc,
    c.category_name,
    p.product_create_at,
    p.product_update_at,
    p.product_delete_at,
    p.product_photo
FROM
    products p
LEFT JOIN
    categories c ON p.product_category_id = c.category_id;

-- Note: This view is used in product_repository.Product struct
-- Table mapping: v_ms_product (this view) <- products (ms_product) + categories

-- +goose Down
DROP VIEW IF EXISTS v_ms_product;
