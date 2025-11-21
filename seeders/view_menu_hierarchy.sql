-- View Main Menus
SELECT 
    permission_id,
    permission_name,
    permission_resource,
    permission_description
FROM permissions
WHERE permission_status = 'main_menu'
ORDER BY permission_id;

-- View Submenus
SELECT 
    permission_id,
    permission_name,
    permission_resource,
    permission_description
FROM permissions
WHERE permission_status = 'submenu'
ORDER BY permission_resource;

-- View complete menu hierarchy
SELECT 
    CASE 
        WHEN permission_status = 'main_menu' THEN '📌 MAIN MENU'
        WHEN permission_status = 'submenu' THEN '  📁 SUBMENU'
        ELSE '    ⚡ ACTION'
    END as type,
    permission_id,
    permission_name,
    permission_resource,
    permission_action,
    permission_status
FROM permissions
ORDER BY 
    CASE permission_resource
        WHEN 'dashboard' THEN 1
        WHEN 'master_data' THEN 2
        WHEN 'category' THEN 3
        WHEN 'customer' THEN 4
        WHEN 'customer_address' THEN 5
        WHEN 'major' THEN 6
        WHEN 'product' THEN 7
        WHEN 'varian' THEN 8
        WHEN 'territory' THEN 9
        WHEN 'stock' THEN 10
        WHEN 'kasir' THEN 11
        WHEN 'pos' THEN 12
        WHEN 'transaction' THEN 13
        WHEN 'tempcart' THEN 14
        WHEN 'order' THEN 15
        WHEN 'laporan' THEN 16
        WHEN 'report' THEN 17
        ELSE 99
    END,
    FIELD(permission_status, 'main_menu', 'submenu', 'action'),
    permission_action;

-- Get menu structure for a specific role
SELECT 
    CASE permission_status
        WHEN 'main_menu' THEN CONCAT('📌 ', permission_name)
        WHEN 'submenu' THEN CONCAT('  📁 ', permission_name)
        ELSE CONCAT('    ⚡ ', permission_name)
    END as menu_item,
    permission_resource as resource,
    permission_action as action,
    permission_status as type
FROM permissions p
INNER JOIN permission_role pr ON p.permission_id = pr.permission_id
WHERE pr.role_id = 1
ORDER BY 
    CASE permission_resource
        WHEN 'dashboard' THEN 1
        WHEN 'master_data' THEN 2
        WHEN 'category' THEN 3
        WHEN 'customer' THEN 4
        WHEN 'major' THEN 6
        WHEN 'product' THEN 7
        WHEN 'kasir' THEN 11
        WHEN 'pos' THEN 12
        WHEN 'order' THEN 15
        WHEN 'laporan' THEN 16
        WHEN 'report' THEN 17
        ELSE 99
    END,
    FIELD(permission_status, 'main_menu', 'submenu', 'action'),
    permission_action;

-- Count permissions by type and resource
SELECT 
    permission_status,
    permission_resource,
    COUNT(*) as total
FROM permissions
GROUP BY permission_status, permission_resource
ORDER BY 
    FIELD(permission_status, 'main_menu', 'submenu', 'action'),
    permission_resource;
