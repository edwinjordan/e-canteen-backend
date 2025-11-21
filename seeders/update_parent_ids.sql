-- Update parent_id for existing permissions

-- Dashboard actions -> parent: Dashboard main menu
UPDATE permissions SET permission_parent_id = (SELECT permission_id FROM (SELECT * FROM permissions) AS p WHERE p.permission_resource = 'dashboard' AND p.permission_status = 'main_menu')
WHERE permission_resource = 'dashboard' AND permission_status = 'action';

-- Master Data submenus -> parent: Master Data main menu
UPDATE permissions SET permission_parent_id = (SELECT permission_id FROM (SELECT * FROM permissions) AS p WHERE p.permission_resource = 'master_data' AND p.permission_status = 'main_menu')
WHERE permission_resource IN ('category', 'customer', 'major', 'product', 'territory', 'stock') AND permission_status = 'submenu';

-- Category actions -> parent: Category submenu
UPDATE permissions SET permission_parent_id = (SELECT permission_id FROM (SELECT * FROM permissions) AS p WHERE p.permission_resource = 'category' AND p.permission_status = 'submenu')
WHERE permission_resource = 'category' AND permission_status = 'action';

-- Customer actions -> parent: Customers submenu
UPDATE permissions SET permission_parent_id = (SELECT permission_id FROM (SELECT * FROM permissions) AS p WHERE p.permission_resource = 'customer' AND p.permission_status = 'submenu')
WHERE permission_resource IN ('customer', 'customer_address') AND permission_status = 'action';

-- Major actions -> parent: Majors submenu
UPDATE permissions SET permission_parent_id = (SELECT permission_id FROM (SELECT * FROM permissions) AS p WHERE p.permission_resource = 'major' AND p.permission_status = 'submenu')
WHERE permission_resource = 'major' AND permission_status = 'action';

-- Product and Varian actions -> parent: Produk submenu
UPDATE permissions SET permission_parent_id = (SELECT permission_id FROM (SELECT * FROM permissions) AS p WHERE p.permission_resource = 'product' AND p.permission_status = 'submenu')
WHERE permission_resource IN ('product', 'varian') AND permission_status = 'action';

-- Territory actions -> parent: Territory submenu
UPDATE permissions SET permission_parent_id = (SELECT permission_id FROM (SELECT * FROM permissions) AS p WHERE p.permission_resource = 'territory' AND p.permission_status = 'submenu')
WHERE permission_resource = 'territory' AND permission_status = 'action';

-- Stock actions -> parent: Stock submenu
UPDATE permissions SET permission_parent_id = (SELECT permission_id FROM (SELECT * FROM permissions) AS p WHERE p.permission_resource = 'stock' AND p.permission_status = 'submenu')
WHERE permission_resource = 'stock' AND permission_status = 'action';

-- Kasir submenus -> parent: Kasir main menu
UPDATE permissions SET permission_parent_id = (SELECT permission_id FROM (SELECT * FROM permissions) AS p WHERE p.permission_resource = 'kasir' AND p.permission_status = 'main_menu')
WHERE permission_resource IN ('pos', 'order') AND permission_status = 'submenu';

-- POS and Transaction actions -> parent: POS submenu
UPDATE permissions SET permission_parent_id = (SELECT permission_id FROM (SELECT * FROM permissions) AS p WHERE p.permission_resource = 'pos' AND p.permission_status = 'submenu')
WHERE permission_resource IN ('pos', 'transaction', 'tempcart') AND permission_status = 'action';

-- Order actions -> parent: Order Management submenu
UPDATE permissions SET permission_parent_id = (SELECT permission_id FROM (SELECT * FROM permissions) AS p WHERE p.permission_resource = 'order' AND p.permission_status = 'submenu')
WHERE permission_resource = 'order' AND permission_status = 'action';

-- Laporan submenu -> parent: Laporan main menu
UPDATE permissions SET permission_parent_id = (SELECT permission_id FROM (SELECT * FROM permissions) AS p WHERE p.permission_resource = 'laporan' AND p.permission_status = 'main_menu')
WHERE permission_resource = 'report' AND permission_status = 'submenu';

-- Report actions -> parent: Laporan Penjualan submenu
UPDATE permissions SET permission_parent_id = (SELECT permission_id FROM (SELECT * FROM permissions) AS p WHERE p.permission_resource = 'report' AND p.permission_status = 'submenu')
WHERE permission_resource = 'report' AND permission_status = 'action';

-- User Management actions -> parent: User Management submenu
UPDATE permissions SET permission_parent_id = (SELECT permission_id FROM (SELECT * FROM permissions) AS p WHERE p.permission_resource = 'user' AND p.permission_status = 'submenu')
WHERE permission_resource = 'user' AND permission_status = 'action';

-- Employee Management actions -> parent: Employee Management submenu
UPDATE permissions SET permission_parent_id = (SELECT permission_id FROM (SELECT * FROM permissions) AS p WHERE p.permission_resource = 'pegawai' AND p.permission_status = 'submenu')
WHERE permission_resource = 'pegawai' AND permission_status = 'action';

-- Permission Management actions -> parent: Permission Management submenu
UPDATE permissions SET permission_parent_id = (SELECT permission_id FROM (SELECT * FROM permissions) AS p WHERE p.permission_resource = 'permission' AND p.permission_status = 'submenu')
WHERE permission_resource = 'permission' AND permission_status = 'action';
