# Permission Status Implementation

## ✅ Berhasil Ditambahkan

Field `permission_status` untuk membedakan **Main Menu**, **Submenu**, dan **Action** pada tabel permissions.

## Database Schema

### Field Baru: `permission_status`
```sql
permission_status ENUM('main_menu', 'submenu', 'action') NOT NULL DEFAULT 'action'
```

## Struktur Permissions

### Total: 83 Permissions

| Status      | Jumlah | Keterangan                          |
|-------------|--------|-------------------------------------|
| main_menu   | 4      | Menu utama aplikasi                 |
| submenu     | 12     | Submenu di bawah main menu          |
| action      | 67     | Aksi/permission spesifik            |

## Main Menus (4)

| ID  | Name        | Resource    |
|-----|-------------|-------------|
| 68  | Dashboard   | dashboard   |
| 69  | Master Data | master_data |
| 70  | Kasir       | kasir       |
| 71  | Laporan     | laporan     |

## Submenus (12)

### Master Data Submenus (6)
| ID  | Name      | Resource  | Parent      |
|-----|-----------|-----------|-------------|
| 72  | Category  | category  | Master Data |
| 73  | Customers | customer  | Master Data |
| 74  | Majors    | major     | Master Data |
| 75  | Produk    | product   | Master Data |
| 76  | Territory | territory | Master Data |
| 77  | Stock     | stock     | Master Data |

### Kasir Submenus (2)
| ID  | Name             | Resource | Parent |
|-----|------------------|----------|--------|
| 78  | POS              | pos      | Kasir  |
| 79  | Order Management | order    | Kasir  |

### Laporan Submenus (1)
| ID  | Name              | Resource | Parent  |
|-----|-------------------|----------|---------|
| 80  | Laporan Penjualan | report   | Laporan |

### System/Admin Submenus (3)
| ID  | Name                  | Resource   | Parent       |
|-----|-----------------------|------------|--------------|
| 81  | User Management       | user       | System/Admin |
| 82  | Employee Management   | pegawai    | System/Admin |
| 83  | Permission Management | permission | System/Admin |

## Actions (67)

Actions adalah permission spesifik untuk CRUD operations dan special actions.

### Contoh struktur per submenu:

**Category (4 actions)**
- Create Category (category.create)
- Read Category (category.read)
- Update Category (category.update)
- Delete Category (category.delete)

**POS (9 actions)**
- Access POS (pos.access)
- Transaction: create, read, update, delete
- Temp Cart: create, read, update, delete

**Order Management (7 actions)**
- Order: create, read, update, delete, process, finish, cancel

## Struktur Hierarki Lengkap

```
📌 Dashboard (main_menu)
   └─ ⚡ View Dashboard (action)

📌 Master Data (main_menu)
   ├─ 📁 Category (submenu)
   │  ├─ ⚡ Create Category
   │  ├─ ⚡ Read Category
   │  ├─ ⚡ Update Category
   │  └─ ⚡ Delete Category
   │
   ├─ 📁 Customers (submenu)
   │  ├─ ⚡ Create Customer
   │  ├─ ⚡ Read Customer
   │  ├─ ⚡ Update Customer
   │  ├─ ⚡ Delete Customer
   │  ├─ ⚡ Create Customer Address
   │  ├─ ⚡ Read Customer Address
   │  ├─ ⚡ Update Customer Address
   │  └─ ⚡ Delete Customer Address
   │
   ├─ 📁 Majors (submenu)
   ├─ 📁 Produk (submenu)
   ├─ 📁 Territory (submenu)
   └─ 📁 Stock (submenu)

📌 Kasir (main_menu)
   ├─ 📁 POS (submenu)
   │  ├─ ⚡ Access POS
   │  ├─ ⚡ Transaction (CRUD)
   │  └─ ⚡ Temp Cart (CRUD)
   │
   └─ 📁 Order Management (submenu)
      └─ ⚡ Order (CRUD + process, finish, cancel)

📌 Laporan (main_menu)
   └─ 📁 Laporan Penjualan (submenu)
      ├─ ⚡ View Sales Report
      ├─ ⚡ View Order Report
      ├─ ⚡ View Transaction Report
      └─ ⚡ Export Sales Report
```

## Penggunaan di Frontend

### 1. Render Main Menu
```javascript
// Get main menus
const mainMenus = permissions.filter(p => p.permission_status === 'main_menu');

// Example:
// - Dashboard
// - Master Data
// - Kasir
// - Laporan
```

### 2. Render Submenus
```javascript
// Get submenus for Master Data
const masterDataSubmenus = permissions.filter(p => 
    p.permission_status === 'submenu' && 
    p.permission_resource.includes('category', 'customer', 'major', 'product', 'territory', 'stock')
);

// Or get by specific parent
const kasirSubmenus = permissions.filter(p => 
    p.permission_status === 'submenu' && 
    ['pos', 'order'].includes(p.permission_resource)
);
```

### 3. Check Actions
```javascript
// Get actions for a submenu (e.g., Category)
const categoryActions = permissions.filter(p => 
    p.permission_status === 'action' && 
    p.permission_resource === 'category'
);

// Check if user can create category
const canCreate = categoryActions.some(p => p.permission_action === 'create');
```

## API Endpoint untuk Filter berdasarkan Status

### Get Main Menus
```
GET /api/permission?status=main_menu
```

### Get Submenus
```
GET /api/permission?status=submenu
```

### Get Actions for a Resource
```
GET /api/permission?status=action&resource=category
```

## SQL Query Examples

### 1. Get all main menus
```sql
SELECT * FROM permissions 
WHERE permission_status = 'main_menu'
ORDER BY permission_id;
```

### 2. Get submenus with their parent
```sql
SELECT 
    p.permission_name as submenu,
    p.permission_resource,
    CASE 
        WHEN p.permission_resource IN ('category', 'customer', 'major', 'product', 'territory', 'stock') THEN 'Master Data'
        WHEN p.permission_resource IN ('pos', 'order') THEN 'Kasir'
        WHEN p.permission_resource = 'report' THEN 'Laporan'
        ELSE 'System/Admin'
    END as parent_menu
FROM permissions p
WHERE permission_status = 'submenu'
ORDER BY parent_menu, permission_name;
```

### 3. Get complete hierarchy for a role
```sql
SELECT 
    CASE permission_status
        WHEN 'main_menu' THEN CONCAT('📌 ', permission_name)
        WHEN 'submenu' THEN CONCAT('  📁 ', permission_name)
        ELSE CONCAT('    ⚡ ', permission_name)
    END as menu_hierarchy,
    permission_resource,
    permission_action
FROM permissions p
INNER JOIN permission_role pr ON p.permission_id = pr.permission_id
WHERE pr.role_id = ?
ORDER BY 
    CASE permission_resource
        WHEN 'dashboard' THEN 1
        WHEN 'master_data' THEN 2
        WHEN 'kasir' THEN 3
        WHEN 'laporan' THEN 4
        ELSE 99
    END,
    FIELD(permission_status, 'main_menu', 'submenu', 'action');
```

## Migration File

File: `migrations/00024_add_permission_status.sql`

```sql
-- +goose Up
ALTER TABLE `permissions` 
ADD COLUMN `permission_status` ENUM('main_menu', 'submenu', 'action') NOT NULL DEFAULT 'action' 
AFTER `permission_description`;

-- +goose Down
ALTER TABLE `permissions` 
DROP COLUMN `permission_status`;
```

## Updated Entity

File: `entity/entity.permission.go`

```go
type Permission struct {
    PermissionId          int    `json:"permission_id"`
    PermissionName        string `json:"permission_name"`
    PermissionResource    string `json:"permission_resource"`
    PermissionAction      string `json:"permission_action"`
    PermissionDescription string `json:"permission_description,omitempty"`
    PermissionStatus      string `json:"permission_status"` // NEW
}

type CreatePermissionRequest struct {
    PermissionName        string `json:"permission_name" validate:"required"`
    PermissionResource    string `json:"permission_resource" validate:"required"`
    PermissionAction      string `json:"permission_action" validate:"required"`
    PermissionDescription string `json:"permission_description,omitempty"`
    PermissionStatus      string `json:"permission_status" validate:"required,oneof=main_menu submenu action"` // NEW
}
```

## Verification

```bash
# Check status distribution
mysql -u root e-canteen_new -e "
SELECT permission_status, COUNT(*) as total 
FROM permissions 
GROUP BY permission_status;"

# Result:
# main_menu: 4
# submenu: 12
# action: 67

# View all main menus and submenus
mysql -u root e-canteen_new < seeders/view_menu_hierarchy.sql
```

## Files Modified/Created

1. ✅ `migrations/00024_add_permission_status.sql` - Migration untuk add column
2. ✅ `entity/entity.permission.go` - Update entity dengan field status
3. ✅ `repository/permission_repository/model.permission.go` - Update model GORM
4. ✅ `seeders/permission_seeder.go` - Update seeder dengan status
5. ✅ `seeders/view_menu_hierarchy.sql` - SQL queries untuk view hierarchy

## Status

✅ **COMPLETED**
- Total permissions: 83
- Main menus: 4
- Submenus: 12
- Actions: 67
- All assigned to Super Admin (role_id = 1)
