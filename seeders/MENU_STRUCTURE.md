# Permission Menu Structure

This document outlines the hierarchical menu structure and associated permissions for the e-Canteen application.

## Menu Hierarchy

```
├── Dashboard
│   └── View Dashboard (dashboard.view)
│
├── Master Data
│   ├── Category
│   │   ├── Create Category (category.create)
│   │   ├── Read Category (category.read)
│   │   ├── Update Category (category.update)
│   │   └── Delete Category (category.delete)
│   │
│   ├── Customers
│   │   ├── Create Customer (customer.create)
│   │   ├── Read Customer (customer.read)
│   │   ├── Update Customer (customer.update)
│   │   ├── Delete Customer (customer.delete)
│   │   ├── Create Customer Address (customer_address.create)
│   │   ├── Read Customer Address (customer_address.read)
│   │   ├── Update Customer Address (customer_address.update)
│   │   └── Delete Customer Address (customer_address.delete)
│   │
│   ├── Majors
│   │   ├── Create Major (major.create)
│   │   ├── Read Major (major.read)
│   │   ├── Update Major (major.update)
│   │   └── Delete Major (major.delete)
│   │
│   └── Produk
│       ├── Create Product (product.create)
│       ├── Read Product (product.read)
│       ├── Update Product (product.update)
│       ├── Delete Product (product.delete)
│       ├── Create Varian (varian.create)
│       ├── Read Varian (varian.read)
│       ├── Update Varian (varian.update)
│       └── Delete Varian (varian.delete)
│
├── Kasir
│   ├── POS
│   │   ├── Access POS (pos.access)
│   │   ├── Create Transaction (transaction.create)
│   │   ├── Read Transaction (transaction.read)
│   │   ├── Update Transaction (transaction.update)
│   │   ├── Delete Transaction (transaction.delete)
│   │   ├── Create Temp Cart (tempcart.create)
│   │   ├── Read Temp Cart (tempcart.read)
│   │   ├── Update Temp Cart (tempcart.update)
│   │   └── Delete Temp Cart (tempcart.delete)
│   │
│   └── Order Management
│       ├── Create Order (order.create)
│       ├── Read Order (order.read)
│       ├── Update Order (order.update)
│       ├── Delete Order (order.delete)
│       ├── Process Order (order.process)
│       ├── Finish Order (order.finish)
│       └── Cancel Order (order.cancel)
│
└── Laporan
    └── Laporan Penjualan
        ├── View Sales Report (report.sales)
        ├── View Order Report (report.order)
        ├── View Transaction Report (report.transaction)
        └── Export Sales Report (report.export)
```

## Additional System Permissions

These are typically for Super Admin only:

```
System / Admin
├── User Management
│   ├── Create User (user.create)
│   ├── Read User (user.read)
│   ├── Update User (user.update)
│   └── Delete User (user.delete)
│
├── Employee Management
│   ├── Create Pegawai (pegawai.create)
│   ├── Read Pegawai (pegawai.read)
│   ├── Update Pegawai (pegawai.update)
│   └── Delete Pegawai (pegawai.delete)
│
├── Permission Management
│   ├── Create Permission (permission.create)
│   ├── Read Permission (permission.read)
│   ├── Update Permission (permission.update)
│   ├── Delete Permission (permission.delete)
│   ├── Assign Permission (permission.assign)
│   └── Revoke Permission (permission.revoke)
│
└── Master Data (Additional)
    ├── Territory Management
    │   ├── Create Territory (territory.create)
    │   ├── Read Territory (territory.read)
    │   ├── Update Territory (territory.update)
    │   └── Delete Territory (territory.delete)
    │
    └── Stock Management
        ├── Create Stock (stock.create)
        ├── Read Stock (stock.read)
        ├── Update Stock (stock.update)
        └── Delete Stock (stock.delete)
```

## Total Permissions: 71

### Breakdown by Menu:
- **Dashboard**: 1 permission
- **Master Data**: 32 permissions
  - Category: 4
  - Customers: 8 (including addresses)
  - Majors: 4
  - Produk: 8 (including varian)
  - Territory: 4
  - Stock: 4
- **Kasir**: 16 permissions
  - POS: 9 (including transactions and temp cart)
  - Order Management: 7
- **Laporan**: 4 permissions
  - Laporan Penjualan: 4
- **System/Admin**: 18 permissions
  - User Management: 4
  - Employee Management: 4
  - Permission Management: 6
  - Additional: 4

## Permission Naming Convention

Format: `{Resource}.{Action}`

Examples:
- `dashboard.view`
- `category.create`
- `product.read`
- `order.process`
- `report.sales`

## Checking Permissions in Code

```go
// Example: Check if user can create products
hasPermission := permissionRoleRepo.CheckPermission(ctx, roleId, "product", "create")

// Example: Check if user can access POS
hasPermission := permissionRoleRepo.CheckPermission(ctx, roleId, "pos", "access")

// Example: Check if user can view sales report
hasPermission := permissionRoleRepo.CheckPermission(ctx, roleId, "report", "sales")
```

## Role-Based Permission Suggestions

### Super Admin (role_id = 1)
- **All permissions** (71 total)

### Kasir (role_id = 2) - Suggested Permissions
- Dashboard: view
- POS: all permissions (access, transaction CRUD, temp cart CRUD)
- Order Management: all permissions
- Product: read (view products for selling)
- Customer: read, create (register new customers)
- Stock: read (check stock availability)

### Manager (role_id = 3) - Suggested Permissions
- Dashboard: view
- Master Data: all permissions (category, customer, major, product, varian, territory, stock)
- Kasir: all permissions (POS, order management)
- Laporan: all permissions (sales, order, transaction reports)
- User: read (view users)
- Pegawai: all permissions

### Customer (role_id = 4) - Suggested Permissions
- Order: create, read (own orders only)
- Product: read
- Customer Address: CRUD (own addresses only)

## Implementation Notes

1. **Menu Visibility**: Use permissions to control which menu items are visible to users
2. **Route Protection**: Apply permission checks in middleware before allowing access to endpoints
3. **Feature Flags**: Use permissions to enable/disable specific features per role
4. **Audit Trail**: Log permission checks for security auditing

## SQL Queries for Menu-Based Permissions

```sql
-- Get all Master Data permissions
SELECT * FROM permissions 
WHERE permission_resource IN ('category', 'customer', 'customer_address', 'major', 'product', 'varian', 'territory', 'stock')
ORDER BY permission_resource, permission_action;

-- Get all Kasir permissions
SELECT * FROM permissions 
WHERE permission_resource IN ('pos', 'transaction', 'tempcart', 'order')
ORDER BY permission_resource, permission_action;

-- Get all Laporan permissions
SELECT * FROM permissions 
WHERE permission_resource = 'report'
ORDER BY permission_action;

-- Get permissions for a specific role grouped by menu
SELECT 
    CASE 
        WHEN p.permission_resource IN ('category', 'customer', 'customer_address', 'major', 'product', 'varian', 'territory', 'stock') THEN 'Master Data'
        WHEN p.permission_resource IN ('pos', 'transaction', 'tempcart', 'order') THEN 'Kasir'
        WHEN p.permission_resource = 'report' THEN 'Laporan'
        WHEN p.permission_resource = 'dashboard' THEN 'Dashboard'
        ELSE 'System/Admin'
    END AS menu,
    p.permission_resource as resource,
    p.permission_action as action,
    p.permission_name as name
FROM permissions p
INNER JOIN permission_role pr ON p.permission_id = pr.permission_id
WHERE pr.role_id = ?
ORDER BY menu, resource, action;
```
