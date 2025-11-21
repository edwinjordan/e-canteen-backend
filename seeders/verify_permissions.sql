-- View all permissions grouped by resource
SELECT 
    permission_resource as resource,
    GROUP_CONCAT(permission_action ORDER BY permission_action SEPARATOR ', ') as actions,
    COUNT(*) as count
FROM permissions
GROUP BY permission_resource
ORDER BY permission_resource;

-- View Super Admin (role_id = 1) permissions
SELECT 
    p.permission_id,
    p.permission_name,
    p.permission_resource,
    p.permission_action,
    p.permission_description
FROM permissions p
INNER JOIN permission_role pr ON p.permission_id = pr.permission_id
WHERE pr.role_id = 1
ORDER BY p.permission_resource, p.permission_action;

-- Count permissions by resource
SELECT 
    permission_resource,
    COUNT(*) as total
FROM permissions
GROUP BY permission_resource
ORDER BY total DESC;

-- View all role-permission assignments
SELECT 
    pr.role_id,
    COUNT(pr.permission_id) as total_permissions
FROM permission_role pr
GROUP BY pr.role_id;
