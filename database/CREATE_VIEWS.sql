-- ============================================
-- CREATE VIEW QUERIES FOR E-CANTEEN-CASHIER-API
-- ============================================
-- Created: 2025-11-16
-- Description: SQL queries to create views for products and product variants

-- ============================================
-- OPTION 1: If your base tables are named 'products' and 'product_varians'
-- ============================================

-- Drop existing views if they exist
DROP VIEW IF EXISTS v_ms_product;
DROP VIEW IF EXISTS v_ms_product_varian;

-- Create view: v_ms_product
-- Purpose: Join products with categories to get category_name
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

-- Create view: v_ms_product_varian
-- Purpose: Join product_varians with products to get product_name
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


-- ============================================
-- OPTION 2: If your base tables are named 'ms_product' and 'ms_product_varian'
-- ============================================

-- Drop existing views if they exist
DROP VIEW IF EXISTS v_ms_product;
DROP VIEW IF EXISTS v_ms_product_varian;

-- Create view: v_ms_product
-- Purpose: Join ms_product with ms_category to get category_name
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
    ms_product p
LEFT JOIN
    ms_category c ON p.product_category_id = c.category_id;

-- Create view: v_ms_product_varian
-- Purpose: Join ms_product_varian with ms_product to get product_name
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
    ms_product_varian pv
LEFT JOIN
    ms_product p ON pv.product_id = p.product_id;


-- ============================================
-- USAGE NOTES
-- ============================================
-- 1. Choose OPTION 1 or OPTION 2 based on your actual table names
-- 2. Run only one option, not both
-- 3. These views are used by:
--    - v_ms_product: product_repository.Product struct (for SELECT queries)
--    - v_ms_product_varian: varian_repository.Varian struct (for SELECT queries)
-- 4. For INSERT operations, the code uses:
--    - ms_product table (via ProductInsert struct)
--    - ms_product_varian table (via VarianInsert struct)

-- ============================================
-- VERIFICATION QUERIES
-- ============================================
-- Run these to verify the views were created successfully:

-- Check v_ms_product view
SELECT * FROM v_ms_product LIMIT 5;

-- Check v_ms_product_varian view
SELECT * FROM v_ms_product_varian LIMIT 5;

-- Count products with categories
SELECT
    COUNT(*) as total_products,
    COUNT(category_name) as products_with_category
FROM v_ms_product;

-- Count variants per product
SELECT
    product_id,
    product_name,
    COUNT(*) as variant_count
FROM v_ms_product_varian
GROUP BY product_id, product_name
ORDER BY variant_count DESC;
