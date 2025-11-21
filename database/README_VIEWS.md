# Database Views Documentation

## Overview
This document explains the database views used in the e-canteen-cashier-api project.

## Views Structure

### 1. v_ms_product

**Purpose**: Combines product data with category information

**Base Tables**:
- `products` (or `ms_product`)
- `categories` (or `ms_category`)

**Columns**:
| Column | Type | Source | Description |
|--------|------|--------|-------------|
| product_id | VARCHAR(36) | products | Primary key |
| product_code | VARCHAR(20) | products | Unique product code |
| product_name | VARCHAR(100) | products | Product name |
| product_category_id | VARCHAR(36) | products | Foreign key to categories |
| product_desc | TEXT | products | Product description |
| category_name | VARCHAR(100) | categories | Category name (joined) |
| product_create_at | TIMESTAMP | products | Creation timestamp |
| product_update_at | TIMESTAMP | products | Last update timestamp |
| product_delete_at | TIMESTAMP | products | Soft delete timestamp |
| product_photo | VARCHAR(255) | products | Product photo URL |

**Join Logic**:
```sql
LEFT JOIN categories c ON p.product_category_id = c.category_id
```

**Used By**: `product_repository.Product` struct (for SELECT/READ operations)

---

### 2. v_ms_product_varian

**Purpose**: Combines product variant data with product information

**Base Tables**:
- `product_varians` (or `ms_product_varian`)
- `products` (or `ms_product`)

**Columns**:
| Column | Type | Source | Description |
|--------|------|--------|-------------|
| product_varian_id | VARCHAR(36) | product_varians | Primary key |
| product_id | VARCHAR(36) | product_varians | Foreign key to products |
| product_name | VARCHAR(100) | products | Product name (joined) |
| varian_name | VARCHAR(100) | product_varians | Variant name |
| product_varian_price | INT | product_varians | Variant price |
| product_varian_qty_booth | INT | product_varians | Quantity in booth |
| product_varian_qty_warehouse | VARCHAR(20) | product_varians | Quantity in warehouse |
| varian_id | VARCHAR(36) | product_varians | Variant identifier |
| product_varian_qty_left | INT | product_varians | Remaining quantity |

**Join Logic**:
```sql
LEFT JOIN products p ON pv.product_id = p.product_id
```

**Used By**: `varian_repository.Varian` struct (for SELECT/READ operations)

---

## Table Naming Convention

The project uses two naming conventions:

### View Names (for SELECT operations)
- `v_ms_product` - View for reading product data
- `v_ms_product_varian` - View for reading variant data

### Base Table Names (for INSERT operations)
Choose one of the following based on your database setup:

**Option A**: Standard naming
- `products` - Base product table
- `product_varians` - Base variant table
- `categories` - Base category table

**Option B**: Prefixed naming
- `ms_product` - Base product table
- `ms_product_varian` - Base variant table
- `ms_category` - Base category table

---

## Setup Instructions

### 1. Check Your Table Names

Run this query to check which naming convention your database uses:

```sql
SHOW TABLES LIKE '%product%';
```

### 2. Create the Views

Open `database/CREATE_VIEWS.sql` and:
1. Choose **OPTION 1** if tables are named `products`, `product_varians`, `categories`
2. Choose **OPTION 2** if tables are named `ms_product`, `ms_product_varian`, `ms_category`
3. Run only the chosen option

### 3. Verify the Views

After creating the views, run:

```sql
-- Verify v_ms_product
SELECT * FROM v_ms_product LIMIT 5;

-- Verify v_ms_product_varian
SELECT * FROM v_ms_product_varian LIMIT 5;
```

---

## Usage in Code

### Reading Products (uses view)
```go
// repository/product_repository/repository.go
func (repo *ProductRepositoryImpl) FindAll(ctx context.Context, where entity.Product, config map[string]interface{}) []entity.Product {
    product := []Product{}
    // This queries v_ms_product view
    tx.WithContext(ctx).Where(whereProduct).Find(&product)
    // ...
}
```

### Inserting Products (uses base table)
```go
// repository/product_repository/repository.go
func (repo *ProductRepositoryImpl) Insert(ctx context.Context, data entity.Product, varians []entity.Varian) error {
    // This inserts into ms_product table directly
    productData := &ProductInsert{
        ProductId:   data.ProductId,
        ProductName: data.ProductName,
        // ...
    }
    tx.WithContext(ctx).Create(&productData)
    // ...
}
```

---

## Benefits of Using Views

1. **Cleaner Code**: Repositories don't need to write JOIN queries manually
2. **Consistent Data**: All reads get the same joined data structure
3. **Performance**: Database can optimize view queries
4. **Maintainability**: Change JOIN logic in one place (the view definition)
5. **Separation**: Read operations use views, write operations use base tables

---

## Troubleshooting

### View doesn't exist error
```
Error: Table 'database.v_ms_product' doesn't exist
```

**Solution**: Run the CREATE VIEW queries from `CREATE_VIEWS.sql`

### Wrong table name in view
```
Error: Table 'database.products' doesn't exist
```

**Solution**: You're using the wrong OPTION. Check your actual table names and use the correct OPTION in `CREATE_VIEWS.sql`

### Column mismatch error
```
Error: Unknown column 'category_name'
```

**Solution**: The view might not be created yet. Run the CREATE VIEW queries.

---

## Related Files

- [CREATE_VIEWS.sql](./CREATE_VIEWS.sql) - SQL queries to create the views
- [repository/product_repository/model.go](../repository/product_repository/model.go) - Product model using v_ms_product
- [repository/varian_repository/model.go](../repository/varian_repository/model.go) - Variant model using v_ms_product_varian
- [migrations/010_create_products_table.sql](./migrations/010_create_products_table.sql) - Product table schema
- [migrations/011_create_product_varians_table.sql](./migrations/011_create_product_varians_table.sql) - Variant table schema

---

## Example Queries

### Get all products with categories
```sql
SELECT * FROM v_ms_product
WHERE product_delete_at IS NULL
ORDER BY product_create_at DESC;
```

### Get all variants for a specific product
```sql
SELECT * FROM v_ms_product_varian
WHERE product_id = 'your-product-id-here';
```

### Get products with variant count
```sql
SELECT
    p.product_id,
    p.product_name,
    p.category_name,
    COUNT(v.product_varian_id) as variant_count,
    SUM(v.product_varian_qty_booth) as total_stock
FROM v_ms_product p
LEFT JOIN v_ms_product_varian v ON p.product_id = v.product_id
WHERE p.product_delete_at IS NULL
GROUP BY p.product_id, p.product_name, p.category_name
ORDER BY variant_count DESC;
```
