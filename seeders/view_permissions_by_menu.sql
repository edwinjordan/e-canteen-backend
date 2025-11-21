-- View all permissions organized by menu structure
SELECT 
    CASE 
        WHEN p.permission_resource = 'dashboard' THEN '1. Dashboard'
        WHEN p.permission_resource IN ('category', 'customer', 'customer_address', 'major', 'product', 'varian', 'territory', 'stock') THEN '2. Master Data'
        WHEN p.permission_resource IN ('pos', 'transaction', 'tempcart', 'order') THEN '3. Kasir'
        WHEN p.permission_resource = 'report' THEN '4. Laporan'
        ELSE '5. System/Admin'
    END AS main_menu,
    CASE 
        WHEN p.permission_resource = 'category' THEN 'Category'
        WHEN p.permission_resource IN ('customer', 'customer_address') THEN 'Customers'
        WHEN p.permission_resource = 'major' THEN 'Majors'
        WHEN p.permission_resource IN ('product', 'varian') THEN 'Produk'
        WHEN p.permission_resource IN ('pos', 'transaction', 'tempcart') THEN 'POS'
        WHEN p.permission_resource = 'order' THEN 'Order Management'
        WHEN p.permission_resource = 'report' THEN 'Laporan Penjualan'
        WHEN p.permission_resource IN ('territory', 'stock') THEN 'Additional Master Data'
        WHEN p.permission_resource = 'user' THEN 'User Management'
        WHEN p.permission_resource = 'pegawai' THEN 'Employee Management'
        WHEN p.permission_resource = 'permission' THEN 'Permission Management'
        ELSE p.permission_resource
    END AS sub_menu,
    p.permission_id,
    p.permission_name,
    p.permission_resource,
    p.permission_action
FROM permissions p
ORDER BY main_menu, sub_menu, p.permission_resource, p.permission_action;

-- Count permissions by main menu
SELECT 
    CASE 
        WHEN p.permission_resource = 'dashboard' THEN 'Dashboard'
        WHEN p.permission_resource IN ('category', 'customer', 'customer_address', 'major', 'product', 'varian', 'territory', 'stock') THEN 'Master Data'
        WHEN p.permission_resource IN ('pos', 'transaction', 'tempcart', 'order') THEN 'Kasir'
        WHEN p.permission_resource = 'report' THEN 'Laporan'
        ELSE 'System/Admin'
    END AS main_menu,
    COUNT(*) as total_permissions
FROM permissions p
GROUP BY main_menu
ORDER BY main_menu;

-- View permissions by submenu
SELECT 
    CASE 
        WHEN p.permission_resource = 'category' THEN 'Master Data > Category'
        WHEN p.permission_resource IN ('customer', 'customer_address') THEN 'Master Data > Customers'
        WHEN p.permission_resource = 'major' THEN 'Master Data > Majors'
        WHEN p.permission_resource IN ('product', 'varian') THEN 'Master Data > Produk'
        WHEN p.permission_resource IN ('pos', 'transaction', 'tempcart') THEN 'Kasir > POS'
        WHEN p.permission_resource = 'order' THEN 'Kasir > Order Management'
        WHEN p.permission_resource = 'report' THEN 'Laporan > Laporan Penjualan'
        ELSE CONCAT('System/Admin > ', p.permission_resource)
    END AS menu_path,
    COUNT(*) as total_permissions,
    GROUP_CONCAT(p.permission_action ORDER BY p.permission_action SEPARATOR ', ') as actions
FROM permissions p
GROUP BY menu_path
ORDER BY menu_path;

-- Super Admin permissions organized by menu
SELECT 
    CASE 
        WHEN p.permission_resource = 'dashboard' THEN '1. Dashboard'
        WHEN p.permission_resource IN ('category', 'customer', 'customer_address', 'major', 'product', 'varian', 'territory', 'stock') THEN '2. Master Data'
        WHEN p.permission_resource IN ('pos', 'transaction', 'tempcart', 'order') THEN '3. Kasir'
        WHEN p.permission_resource = 'report' THEN '4. Laporan'
        ELSE '5. System/Admin'
    END AS main_menu,
    p.permission_resource as resource,
    p.permission_name,
    p.permission_action
FROM permissions p
INNER JOIN permission_role pr ON p.permission_id = pr.permission_id
WHERE pr.role_id = 1
ORDER BY main_menu, p.permission_resource, p.permission_action;
