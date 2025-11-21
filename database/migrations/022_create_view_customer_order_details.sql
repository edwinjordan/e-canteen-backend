-- +goose Up
-- Migration: Create v_tb_customer_order_detail view
-- Created: 2025-11-17
-- Description: View to join customer_order_details with customer_orders, customers, product_varians, and products

-- Drop view if exists
DROP VIEW IF EXISTS v_tb_customer_order_detail;

-- Create view v_tb_customer_order_detail
CREATE VIEW v_tb_customer_order_detail AS
SELECT
    cod.order_detail_id,
    cod.order_detail_parent_id,
    cod.order_detail_product_varian_id,
    cod.order_detail_qty,
    cod.order_detail_price,
    cod.order_detail_subtotal,
    c.customer_name,
    p.product_name,
    pv.varian_name
FROM
    customer_order_details cod
INNER JOIN
    customer_orders co ON cod.order_detail_parent_id = co.order_id
INNER JOIN
    customers c ON co.order_customer_id = c.customer_id
INNER JOIN
    product_varians pv ON cod.order_detail_product_varian_id = pv.product_varian_id
INNER JOIN
    products p ON pv.product_id = p.product_id;

-- Note: This view is used for ViewOrderDetail entity
-- Table mapping: v_tb_customer_order_detail (this view) <- customer_order_details + customer_orders + customers + product_varians + products

-- +goose Down
DROP VIEW IF EXISTS v_tb_customer_order_detail;
