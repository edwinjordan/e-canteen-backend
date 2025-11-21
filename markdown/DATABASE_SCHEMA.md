# Database Schema Documentation

## E-Canteen Cashier API

This document describes all entities and their corresponding database schema for the E-Canteen Cashier API system.

---

## Table of Contents
- [User Management](#user-management)
- [Customer Management](#customer-management)
- [Product Management](#product-management)
- [Order Management](#order-management)
- [Transaction Management](#transaction-management)
- [Supporting Tables](#supporting-tables)

---

## User Management

### Users Table
**Table Name:** `users`

| Column Name | Data Type | Constraints | Description |
|------------|-----------|-------------|-------------|
| user_id | VARCHAR(36) | PRIMARY KEY | Unique user identifier (UUID) |
| user_name | VARCHAR(100) | NOT NULL | User's full name |
| user_email | VARCHAR(100) | NOT NULL, UNIQUE | User's email address |
| user_password | VARCHAR(255) | NOT NULL | Hashed password |
| user_pegawai_id | VARCHAR(36) | FOREIGN KEY | Reference to pegawai table |
| user_has_mobile_access | INT | DEFAULT 0 | Mobile access permission (0=No, 1=Yes) |
| user_role_id | VARCHAR(36) | FOREIGN KEY | Reference to roles table |
| user_create_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | Record creation timestamp |
| user_update_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP ON UPDATE | Record update timestamp |

**Relationships:**
- `user_pegawai_id` → `pegawai.pegawai_id`
- `user_role_id` → `roles.role_id`

---

### Roles Table
**Table Name:** `roles`

| Column Name | Data Type | Constraints | Description |
|------------|-----------|-------------|-------------|
| role_id | VARCHAR(36) | PRIMARY KEY | Unique role identifier (UUID) |
| role_name | VARCHAR(50) | NOT NULL | Role display name |
| role_code | VARCHAR(20) | NOT NULL, UNIQUE | Role code identifier |
| role_create_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | Record creation timestamp |
| role_update_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP ON UPDATE | Record update timestamp |

---

### User Logs Table
**Table Name:** `user_logs`

| Column Name | Data Type | Constraints | Description |
|------------|-----------|-------------|-------------|
| log_user_id | VARCHAR(36) | PRIMARY KEY | Unique log identifier (UUID) |
| log_user_user_id | VARCHAR(36) | FOREIGN KEY | Reference to users table |
| log_user_token | TEXT | NOT NULL | Authentication token |
| log_user_metadata | TEXT | NULL | Additional login metadata (device info, etc.) |
| log_user_login_date | TIMESTAMP | NOT NULL | Login timestamp |
| log_user_logout_date | TIMESTAMP | NULL | Logout timestamp |

**Relationships:**
- `log_user_user_id` → `users.user_id`

---

### User OTP Table
**Table Name:** `user_otp`

| Column Name | Data Type | Constraints | Description |
|------------|-----------|-------------|-------------|
| otp_id | VARCHAR(36) | PRIMARY KEY | Unique OTP identifier (UUID) |
| otp_customer_id | VARCHAR(36) | FOREIGN KEY | Reference to customers table |
| otp_number | VARCHAR(6) | NOT NULL | OTP code |
| otp_status | INT | DEFAULT 0 | OTP status (0=Active, 1=Used) |
| otp_expired | TIMESTAMP | NOT NULL | OTP expiration timestamp |

**Relationships:**
- `otp_customer_id` → `customers.customer_id`

---

## Customer Management

### Customers Table
**Table Name:** `customers`

| Column Name | Data Type | Constraints | Description |
|------------|-----------|-------------|-------------|
| customer_id | VARCHAR(36) | PRIMARY KEY | Unique customer identifier (UUID) |
| customer_code | VARCHAR(20) | NOT NULL, UNIQUE | Customer code |
| customer_name | VARCHAR(100) | NOT NULL | Customer full name |
| customer_gender | ENUM('L','P') | NOT NULL | Gender (L=Male, P=Female) |
| customer_phonenumber | VARCHAR(15) | NOT NULL | Phone number |
| customer_email | VARCHAR(100) | NOT NULL, UNIQUE | Email address |
| customer_dob | DATE | NULL | Date of birth |
| customer_password | VARCHAR(255) | NOT NULL | Hashed password |
| customer_profile_pic | VARCHAR(255) | NULL | Profile picture filename |
| customer_class | VARCHAR(20) | NULL | Student class |
| customer_major_id | VARCHAR(36) | FOREIGN KEY | Reference to majors table |
| customer_profile_pic_path | VARCHAR(255) | NULL | Full path to profile picture |
| customer_status | INT | DEFAULT 1 | Account status (0=Inactive, 1=Active) |
| customer_last_status | INT | DEFAULT 1 | Previous status |
| customer_create_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | Record creation timestamp |
| customer_update_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP ON UPDATE | Record update timestamp |

**Relationships:**
- `customer_major_id` → `majors.major_id`

---

### Customer Address Table
**Table Name:** `customer_address`

| Column Name | Data Type | Constraints | Description |
|------------|-----------|-------------|-------------|
| address_id | VARCHAR(36) | PRIMARY KEY | Unique address identifier (UUID) |
| address_customer_id | VARCHAR(36) | FOREIGN KEY | Reference to customers table |
| address_text | TEXT | NOT NULL | Full address text |
| address_name | VARCHAR(100) | NOT NULL | Address label (e.g., "Home", "School") |
| address_province_id | VARCHAR(10) | NOT NULL | Province ID |
| address_province | VARCHAR(100) | NOT NULL | Province name |
| address_city_id | VARCHAR(10) | NOT NULL | City ID |
| address_city | VARCHAR(100) | NOT NULL | City name |
| address_district_id | VARCHAR(10) | NOT NULL | District ID |
| address_district | VARCHAR(100) | NOT NULL | District name |
| address_village_id | VARCHAR(10) | NOT NULL | Village ID |
| address_village | VARCHAR(100) | NOT NULL | Village name |
| address_postal_code | VARCHAR(10) | NULL | Postal code |
| address_main | INT | DEFAULT 0 | Main address flag (0=No, 1=Yes) |
| address_create_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | Record creation timestamp |
| address_update_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP ON UPDATE | Record update timestamp |

**Relationships:**
- `address_customer_id` → `customers.customer_id`

---

### Majors Table
**Table Name:** `majors`

| Column Name | Data Type | Constraints | Description |
|------------|-----------|-------------|-------------|
| major_id | VARCHAR(36) | PRIMARY KEY | Unique major identifier (UUID) |
| major_name | VARCHAR(100) | NOT NULL | Major/Department name |

---

### Pegawai (Employee) Table
**Table Name:** `pegawai`

| Column Name | Data Type | Constraints | Description |
|------------|-----------|-------------|-------------|
| pegawai_id | VARCHAR(36) | PRIMARY KEY | Unique employee identifier (UUID) |
| pegawai_code | VARCHAR(20) | NOT NULL, UNIQUE | Employee code |
| pegawai_name | VARCHAR(100) | NOT NULL | Employee full name |
| pegawai_gender | ENUM('L','P') | NOT NULL | Gender (L=Male, P=Female) |
| pegawai_phonenumber | VARCHAR(15) | NOT NULL | Phone number |
| pegawai_create_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | Record creation timestamp |
| pegawai_update_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP ON UPDATE | Record update timestamp |
| pegawai_delete_at | TIMESTAMP | NULL | Soft delete timestamp |

---

## Product Management

### Products Table
**Table Name:** `products`

| Column Name | Data Type | Constraints | Description |
|------------|-----------|-------------|-------------|
| product_id | VARCHAR(36) | PRIMARY KEY | Unique product identifier (UUID) |
| product_code | VARCHAR(20) | NOT NULL, UNIQUE | Product code |
| product_name | VARCHAR(100) | NOT NULL | Product name |
| product_category_id | VARCHAR(36) | FOREIGN KEY | Reference to categories table |
| product_desc | TEXT | NULL | Product description |
| product_create_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | Record creation timestamp |
| product_update_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP ON UPDATE | Record update timestamp |
| product_delete_at | TIMESTAMP | NULL | Soft delete timestamp |
| product_photo | VARCHAR(255) | NULL | Product photo filename |

**Relationships:**
- `product_category_id` → `categories.category_id`

---

### Categories Table
**Table Name:** `categories`

| Column Name | Data Type | Constraints | Description |
|------------|-----------|-------------|-------------|
| category_id | VARCHAR(36) | PRIMARY KEY | Unique category identifier (UUID) |
| category_name | VARCHAR(100) | NOT NULL | Category name |
| category_delete_at | TIMESTAMP | NULL | Soft delete timestamp |

---

### Product Variants Table
**Table Name:** `product_varians`

| Column Name | Data Type | Constraints | Description |
|------------|-----------|-------------|-------------|
| product_varian_id | VARCHAR(36) | PRIMARY KEY | Unique variant identifier (UUID) |
| product_id | VARCHAR(36) | FOREIGN KEY | Reference to products table |
| varian_id | VARCHAR(36) | NOT NULL | Variant type identifier |
| varian_name | VARCHAR(100) | NOT NULL | Variant name (e.g., "Small", "Large") |
| product_varian_price | INT | NOT NULL | Variant price |
| product_varian_qty_booth | INT | DEFAULT 0 | Quantity in booth/store |
| product_varian_qty_warehouse | VARCHAR(20) | DEFAULT '0' | Quantity in warehouse |
| product_varian_qty_left | INT | DEFAULT 0 | Remaining quantity |

**Relationships:**
- `product_id` → `products.product_id`

---

### Stock Booth Table
**Table Name:** `stock_booth`

| Column Name | Data Type | Constraints | Description |
|------------|-----------|-------------|-------------|
| product_stok_id | VARCHAR(36) | PRIMARY KEY | Unique stock record identifier (UUID) |
| product_stok_product_varian_id | VARCHAR(36) | FOREIGN KEY | Reference to product_varians table |
| product_stok_first_qty | INT | NOT NULL | Initial quantity |
| product_stok_qty | INT | NOT NULL | Change quantity (can be positive or negative) |
| product_stok_last_qty | INT | NOT NULL | Final quantity after change |
| product_stok_jenis | VARCHAR(20) | NOT NULL | Stock type (e.g., "IN", "OUT", "ADJUSTMENT") |
| product_stok_datetime | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | Stock change timestamp |
| product_stok_pegawai_id | VARCHAR(36) | FOREIGN KEY | Reference to pegawai table |

**Relationships:**
- `product_stok_product_varian_id` → `product_varians.product_varian_id`
- `product_stok_pegawai_id` → `pegawai.pegawai_id`

---

## Order Management

### Customer Orders Table
**Table Name:** `customer_orders`

| Column Name | Data Type | Constraints | Description |
|------------|-----------|-------------|-------------|
| order_id | VARCHAR(36) | PRIMARY KEY | Unique order identifier (UUID) |
| order_customer_id | VARCHAR(36) | FOREIGN KEY | Reference to customers table |
| order_inv_number | VARCHAR(50) | NOT NULL, UNIQUE | Invoice number |
| order_address_id | VARCHAR(36) | FOREIGN KEY | Reference to customer_address table |
| order_delivery_type | VARCHAR(20) | NOT NULL | Delivery type (e.g., "PICKUP", "DELIVERY") |
| order_total_item | INT | NOT NULL | Total number of items |
| order_subtotal | DECIMAL(15,2) | NOT NULL | Subtotal amount |
| order_discount | DECIMAL(15,2) | DEFAULT 0 | Discount amount |
| order_total | DECIMAL(15,2) | NOT NULL | Total amount after discount |
| order_notes | TEXT | NULL | Order notes |
| order_cancel_notes | TEXT | NULL | Cancellation notes |
| order_status | INT | NOT NULL | Order status (0=Pending, 1=Processing, 2=Completed, 3=Cancelled) |
| order_processed_datetime | TIMESTAMP | NULL | Processing timestamp |
| order_processed_by | VARCHAR(36) | NULL | User who processed the order |
| order_finished_datetime | TIMESTAMP | NULL | Completion timestamp |
| order_finished_by | VARCHAR(36) | NULL | User who finished the order |
| order_create_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | Record creation timestamp |

**Relationships:**
- `order_customer_id` → `customers.customer_id`
- `order_address_id` → `customer_address.address_id`
- `order_processed_by` → `users.user_id`
- `order_finished_by` → `users.user_id`

---

### Customer Order Details Table
**Table Name:** `customer_order_details`

| Column Name | Data Type | Constraints | Description |
|------------|-----------|-------------|-------------|
| order_detail_id | VARCHAR(36) | PRIMARY KEY | Unique order detail identifier (UUID) |
| order_detail_parent_id | VARCHAR(36) | FOREIGN KEY | Reference to customer_orders table |
| order_detail_product_varian_id | VARCHAR(36) | FOREIGN KEY | Reference to product_varians table |
| order_detail_qty | INT | NOT NULL | Quantity ordered |
| order_detail_price | DECIMAL(15,2) | NOT NULL | Price per unit |
| order_detail_subtotal | DECIMAL(15,2) | NOT NULL | Subtotal (qty * price) |

**Relationships:**
- `order_detail_parent_id` → `customer_orders.order_id`
- `order_detail_product_varian_id` → `product_varians.product_varian_id`

---

## Transaction Management

### Transactions Table
**Table Name:** `transactions`

| Column Name | Data Type | Constraints | Description |
|------------|-----------|-------------|-------------|
| trans_id | VARCHAR(36) | PRIMARY KEY | Unique transaction identifier (UUID) |
| trans_user_id | VARCHAR(36) | FOREIGN KEY | Reference to users table (cashier) |
| trans_customer_id | VARCHAR(36) | FOREIGN KEY | Reference to customers table |
| trans_order_id | VARCHAR(36) | FOREIGN KEY | Reference to customer_orders table |
| trans_invoice | VARCHAR(50) | NOT NULL, UNIQUE | Transaction invoice number |
| trans_qty_total | INT | NOT NULL | Total quantity of items |
| trans_product_total | INT | NOT NULL | Total number of products |
| trans_subtotal | DECIMAL(15,2) | NOT NULL | Subtotal amount |
| trans_discount | DECIMAL(15,2) | DEFAULT 0 | Discount amount |
| trans_total | DECIMAL(15,2) | NOT NULL | Total amount |
| trans_received_total | DECIMAL(15,2) | NOT NULL | Amount received from customer |
| trans_refund_total | DECIMAL(15,2) | DEFAULT 0 | Refund/change amount |
| trans_status | INT | DEFAULT 1 | Transaction status (0=Cancelled, 1=Completed) |
| trans_create_at | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | Transaction timestamp |

**Relationships:**
- `trans_user_id` → `users.user_id`
- `trans_customer_id` → `customers.customer_id`
- `trans_order_id` → `customer_orders.order_id`

---

### Transaction Details Table
**Table Name:** `transaction_details`

| Column Name | Data Type | Constraints | Description |
|------------|-----------|-------------|-------------|
| trans_detail_id | VARCHAR(36) | PRIMARY KEY | Unique transaction detail identifier (UUID) |
| trans_detail_parent_id | VARCHAR(36) | FOREIGN KEY | Reference to transactions table |
| trans_detail_product_varian_id | VARCHAR(36) | FOREIGN KEY | Reference to product_varians table |
| trans_detail_qty | INT | NOT NULL | Quantity sold |
| trans_detail_price | DECIMAL(15,2) | NOT NULL | Price per unit |
| trans_detail_subtotal | DECIMAL(15,2) | NOT NULL | Subtotal (qty * price) |

**Relationships:**
- `trans_detail_parent_id` → `transactions.trans_id`
- `trans_detail_product_varian_id` → `product_varians.product_varian_id`

---

## Supporting Tables

### Temporary Cart Table
**Table Name:** `temp_cart`

| Column Name | Data Type | Constraints | Description |
|------------|-----------|-------------|-------------|
| temp_cart_id | VARCHAR(36) | PRIMARY KEY | Unique cart item identifier (UUID) |
| temp_cart_order_id | VARCHAR(36) | NULL | Temporary order identifier |
| temp_cart_product_varian_id | VARCHAR(36) | FOREIGN KEY | Reference to product_varians table |
| temp_cart_user_id | VARCHAR(36) | FOREIGN KEY | Reference to users table (cashier) |
| temp_cart_qty | INT | NOT NULL | Quantity in cart |

**Relationships:**
- `temp_cart_product_varian_id` → `product_varians.product_varian_id`
- `temp_cart_user_id` → `users.user_id`

**Note:** This is a temporary table for holding cart items before transaction completion.

---

### Version Admin Table
**Table Name:** `version_admin`

| Column Name | Data Type | Constraints | Description |
|------------|-----------|-------------|-------------|
| version_id | INT | PRIMARY KEY AUTO_INCREMENT | Unique version identifier |
| version_number | VARCHAR(20) | NOT NULL | Version number (e.g., "1.0.0") |
| version_code | INT | NOT NULL | Version code (numeric) |
| version_chagelog | TEXT | NULL | Changelog/release notes |
| version_datetime | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | Version release timestamp |

---

### Version Shop Table
**Table Name:** `version_shop`

| Column Name | Data Type | Constraints | Description |
|------------|-----------|-------------|-------------|
| version_id | INT | PRIMARY KEY AUTO_INCREMENT | Unique version identifier |
| version_number | VARCHAR(20) | NOT NULL | Version number (e.g., "1.0.0") |
| version_code | INT | NOT NULL | Version code (numeric) |
| version_chagelog | TEXT | NULL | Changelog/release notes |
| version_datetime | TIMESTAMP | DEFAULT CURRENT_TIMESTAMP | Version release timestamp |

---

## Entity Relationship Diagram (ERD) Summary

### Primary Relationships:

1. **User System:**
   - `users` → `pegawai` (one-to-one)
   - `users` → `roles` (many-to-one)
   - `users` → `user_logs` (one-to-many)

2. **Customer System:**
   - `customers` → `majors` (many-to-one)
   - `customers` → `customer_address` (one-to-many)
   - `customers` → `user_otp` (one-to-many)

3. **Product System:**
   - `products` → `categories` (many-to-one)
   - `products` → `product_varians` (one-to-many)
   - `product_varians` → `stock_booth` (one-to-many)

4. **Order System:**
   - `customer_orders` → `customers` (many-to-one)
   - `customer_orders` → `customer_address` (many-to-one)
   - `customer_orders` → `customer_order_details` (one-to-many)

5. **Transaction System:**
   - `transactions` → `users` (many-to-one)
   - `transactions` → `customers` (many-to-one)
   - `transactions` → `customer_orders` (many-to-one)
   - `transactions` → `transaction_details` (one-to-many)

---

## Indexes Recommendations

For optimal performance, consider creating the following indexes:

```sql
-- Users
CREATE INDEX idx_user_email ON users(user_email);
CREATE INDEX idx_user_pegawai_id ON users(user_pegawai_id);
CREATE INDEX idx_user_role_id ON users(user_role_id);

-- Customers
CREATE INDEX idx_customer_code ON customers(customer_code);
CREATE INDEX idx_customer_email ON customers(customer_email);
CREATE INDEX idx_customer_major_id ON customers(customer_major_id);
CREATE INDEX idx_customer_status ON customers(customer_status);

-- Products
CREATE INDEX idx_product_code ON products(product_code);
CREATE INDEX idx_product_category_id ON products(product_category_id);
CREATE INDEX idx_product_delete_at ON products(product_delete_at);

-- Orders
CREATE INDEX idx_order_customer_id ON customer_orders(order_customer_id);
CREATE INDEX idx_order_inv_number ON customer_orders(order_inv_number);
CREATE INDEX idx_order_status ON customer_orders(order_status);
CREATE INDEX idx_order_create_at ON customer_orders(order_create_at);

-- Transactions
CREATE INDEX idx_trans_user_id ON transactions(trans_user_id);
CREATE INDEX idx_trans_customer_id ON transactions(trans_customer_id);
CREATE INDEX idx_trans_invoice ON transactions(trans_invoice);
CREATE INDEX idx_trans_create_at ON transactions(trans_create_at);
```

---

## Notes

- All `*_id` fields use UUID (VARCHAR(36)) for primary keys
- Soft deletes are implemented using `*_delete_at` timestamp fields
- All monetary values use DECIMAL(15,2) for precision
- Timestamps are managed automatically with DEFAULT CURRENT_TIMESTAMP
- Gender fields use ENUM('L','P') where L=Laki-laki (Male), P=Perempuan (Female)

---

**Last Updated:** November 14, 2025
