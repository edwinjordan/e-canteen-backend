-- Seeder: Customers (Sample Customers)
-- Created: 2025-11-14
-- Default Password: customer123 (hashed with bcrypt: $2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy)

-- Get major IDs
SET @ti_major_id = (SELECT major_id FROM majors WHERE major_name = 'Teknik Informatika' LIMIT 1);
SET @si_major_id = (SELECT major_id FROM majors WHERE major_name = 'Sistem Informasi' LIMIT 1);
SET @akun_major_id = (SELECT major_id FROM majors WHERE major_name = 'Akuntansi' LIMIT 1);

INSERT INTO customers (customer_id, customer_code, customer_name, customer_gender, customer_phonenumber, customer_email, customer_dob, customer_password, customer_class, customer_major_id, customer_status) VALUES
(UUID(), 'CST001', 'Andi Wijaya', 'L', '081234560001', 'andi@student.com', '2003-05-15', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', '3A', @ti_major_id, 1),
(UUID(), 'CST002', 'Rina Susanti', 'P', '081234560002', 'rina@student.com', '2003-08-20', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', '3B', @si_major_id, 1),
(UUID(), 'CST003', 'Joko Prasetyo', 'L', '081234560003', 'joko@student.com', '2004-02-10', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', '2A', @ti_major_id, 1),
(UUID(), 'CST004', 'Lilis Handayani', 'P', '081234560004', 'lilis@student.com', '2004-11-25', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', '2B', @akun_major_id, 1),
(UUID(), 'CST005', 'Agus Setiawan', 'L', '081234560005', 'agus@student.com', '2003-07-30', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', '3C', @si_major_id, 1);
