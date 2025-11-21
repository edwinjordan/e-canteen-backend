-- Seeder: Users (Sample Users)
-- Created: 2025-11-14
-- Default Password: admin123 (hashed with bcrypt: $2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy)

-- First, get the role IDs (adjust these based on your actual role_id values)
SET @superadmin_role_id = (SELECT role_id FROM roles WHERE role_code = 'SUPERADMIN' LIMIT 1);
SET @admin_role_id = (SELECT role_id FROM roles WHERE role_code = 'ADMIN' LIMIT 1);
SET @cashier_role_id = (SELECT role_id FROM roles WHERE role_code = 'CASHIER' LIMIT 1);

INSERT INTO users (user_id, user_name, user_email, user_password, user_pegawai_id, user_has_mobile_access, user_role_id) VALUES
(UUID(), 'Ahmad Fauzi', 'admin@ecanteen.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', '550e8400-e29b-41d4-a716-446655440001', 1, @superadmin_role_id),
(UUID(), 'Siti Nurhaliza', 'siti@ecanteen.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', '550e8400-e29b-41d4-a716-446655440002', 1, @admin_role_id),
(UUID(), 'Budi Santoso', 'budi@ecanteen.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', '550e8400-e29b-41d4-a716-446655440003', 0, @cashier_role_id),
(UUID(), 'Dewi Lestari', 'dewi@ecanteen.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', '550e8400-e29b-41d4-a716-446655440004', 0, @cashier_role_id);
