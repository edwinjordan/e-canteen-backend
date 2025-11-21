-- Seeder: Roles
-- Created: 2025-11-14

INSERT INTO roles (role_id, role_name, role_code, role_create_at) VALUES
(UUID(), 'Super Admin', 'SUPERADMIN', NOW()),
(UUID(), 'Admin', 'ADMIN', NOW()),
(UUID(), 'Kasir', 'CASHIER', NOW()),
(UUID(), 'Gudang', 'WAREHOUSE', NOW()),
(UUID(), 'Manager', 'MANAGER', NOW());
