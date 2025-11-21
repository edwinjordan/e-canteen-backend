-- +goose Up
-- Migration: Create v_ms_product_varian view
-- Created: 2025-11-16
-- Description: View to join product_varians with products

-- Drop view if exists
DROP VIEW IF EXISTS v_ms_product_varian;

-- Create view v_ms_product_varian
CREATE VIEW v_ms_product_varian AS
SELECT
    pv.product_varian_id,
    pv.product_id,
    p.product_name,
    pv.varian_name,
    pv.product_varian_price,
    pv.product_varian_qty_booth,
    pv.product_varian_qty_warehouse,
    pv.varian_id,
    pv.product_varian_qty_left
FROM
    product_varians pv
LEFT JOIN
    products p ON pv.product_id = p.product_id;

-- Note: This view is used in varian_repository.Varian struct
-- Table mapping: v_ms_product_varian (this view) <- product_varians (ms_product_varian) + products

-- +goose Down
DROP VIEW IF EXISTS v_ms_product_varian;
