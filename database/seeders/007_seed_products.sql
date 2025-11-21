-- Seeder: Products (Sample Products)
-- Created: 2025-11-14

-- Get category IDs
SET @makanan_berat_id = (SELECT category_id FROM categories WHERE category_name = 'Makanan Berat' LIMIT 1);
SET @makanan_ringan_id = (SELECT category_id FROM categories WHERE category_name = 'Makanan Ringan' LIMIT 1);
SET @minuman_id = (SELECT category_id FROM categories WHERE category_name = 'Minuman' LIMIT 1);
SET @kue_id = (SELECT category_id FROM categories WHERE category_name = 'Kue & Roti' LIMIT 1);

INSERT INTO products (product_id, product_code, product_name, product_category_id, product_desc) VALUES
('PRD001', 'PRD-001', 'Nasi Goreng Spesial', @makanan_berat_id, 'Nasi goreng dengan telur, ayam, dan sayuran'),
('PRD002', 'PRD-002', 'Mie Ayam Bakso', @makanan_berat_id, 'Mie ayam dengan bakso dan pangsit'),
('PRD003', 'PRD-003', 'Soto Ayam', @makanan_berat_id, 'Soto ayam dengan nasi dan kerupuk'),
('PRD004', 'PRD-004', 'Teh Botol', @minuman_id, 'Teh manis dalam kemasan botol'),
('PRD005', 'PRD-005', 'Kopi Susu', @minuman_id, 'Kopi susu hangat/dingin'),
('PRD006', 'PRD-006', 'Jus Jeruk', @minuman_id, 'Jus jeruk segar'),
('PRD007', 'PRD-007', 'Roti Bakar', @kue_id, 'Roti bakar dengan berbagai topping'),
('PRD008', 'PRD-008', 'Pisang Goreng', @makanan_ringan_id, 'Pisang goreng crispy'),
('PRD009', 'PRD-009', 'Risoles', @makanan_ringan_id, 'Risoles isi sayuran dan daging'),
('PRD010', 'PRD-010', 'Air Mineral', @minuman_id, 'Air mineral kemasan 600ml');
