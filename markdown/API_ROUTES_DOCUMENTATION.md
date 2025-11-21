# API Routes Documentation
## E-Canteen Cashier API

Dokumentasi lengkap untuk route API Login, Dashboard, dan POS (Point of Sale/Kasir) beserta struktur tabel database.

---

## 📋 Table of Contents

1. [Authentication Routes](#authentication-routes)
2. [Dashboard Routes](#dashboard-routes)
3. [POS/Kasir Routes](#poskasir-routes)
4. [Database Schema](#database-schema)

---

## 🔐 Authentication Routes

### 1. Kasir Login

**Endpoint:** `POST /api/kasir/login`

**Description:** Authenticate kasir user dan mendapatkan JWT token

**Request Body:**
```json
{
  "user_email": "kasir@example.com",
  "user_password": "password123",
  "user_fcmtoken": "fcm_token_here",
  "user_device_metadata": "Android 12, Samsung Galaxy S21"
}
```

**Response Success (200):**
```json
{
  "error": false,
  "message": "Login berhasil",
  "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Response Error (401):**
```json
{
  "error": true,
  "message": "Email atau password salah",
  "data": null
}
```

**cURL Example:**
```bash
curl -X POST "http://127.0.0.1:3000/api/kasir/login" \
  -H "Content-Type: application/json" \
  -d '{
    "user_email": "kasir@example.com",
    "user_password": "password123"
  }'
```

---

### 2. Kasir Logout

**Endpoint:** `PUT /api/kasir/logout`

**Description:** Logout kasir user dan invalidate token

**Headers:**
```
Authorization: Bearer <JWT_TOKEN>
```

**Response Success (200):**
```json
{
  "error": false,
  "message": "Logout berhasil",
  "data": null
}
```

**cURL Example:**
```bash
curl -X PUT "http://127.0.0.1:3000/api/kasir/logout" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

### 3. Verify Token

**Endpoint:** `GET /api/verify-token`

**Description:** Validasi bearer token dan return user data

**Headers:**
```
Authorization: Bearer <JWT_TOKEN>
```

**Response Success (200):**
```json
{
  "error": false,
  "message": "Token valid",
  "data": {
    "user_id": "usr-001",
    "user_name": "John Doe",
    "user_email": "kasir@example.com",
    "user_role_id": "role-001"
  }
}
```

---

## 📊 Dashboard Routes

### 1. Get Dashboard Statistics

**Endpoint:** `GET /api/dashboard/stats`

**Description:** Mendapatkan statistik dashboard komprehensif

**Headers:**
```
Authorization: Bearer <JWT_TOKEN>
```

**Query Parameters:**

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| year | int | No | Current year | Tahun untuk data penjualan bulanan |

**Response Success (200):**
```json
{
  "error": false,
  "message": "Berhasil mengambil data dashboard",
  "data": {
    "total_sales": 15750000.50,
    "total_transactions": 1250,
    "total_customers": 450,
    "monthly_sales": [
      {
        "month": 1,
        "month_name": "January",
        "year": 2025,
        "total_sales": 1200000.00
      },
      {
        "month": 2,
        "month_name": "February",
        "year": 2025,
        "total_sales": 1350000.50
      }
    ],
    "top_products": [
      {
        "product_id": "prod-001",
        "product_name": "Nasi Goreng",
        "variant_name": "Reguler",
        "total_quantity": 580,
        "total_sales": 2900000.00
      },
      {
        "product_id": "prod-002",
        "product_name": "Mie Ayam",
        "variant_name": "Jumbo",
        "total_quantity": 450,
        "total_sales": 2475000.00
      }
    ]
  }
}
```

**Features:**
- ✅ Total penjualan (all-time)
- ✅ Total transaksi
- ✅ Total customer
- ✅ Grafik penjualan bulanan (12 bulan)
- ✅ Top 10 produk terlaris

**cURL Example:**
```bash
# Current year
curl -X GET "http://127.0.0.1:3000/api/dashboard/stats" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"

# Specific year
curl -X GET "http://127.0.0.1:3000/api/dashboard/stats?year=2024" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

## 🛒 POS/Kasir Routes

### 1. Create Transaction

**Endpoint:** `POST /api/kasir/transaction`

**Description:** Membuat transaksi kasir baru

**Headers:**
```
Authorization: Bearer <JWT_TOKEN>
```

**Request Body:**
```json
{
  "trans_user_id": "usr-001",
  "trans_customer_id": "cust-001",
  "trans_order_id": "order-001",
  "trans_invoice": "INV-2025-001",
  "trans_qty_total": 5,
  "trans_product_total": 3,
  "trans_subtotal": 150000.00,
  "trans_discount": 10000.00,
  "trans_total": 140000.00,
  "trans_received_total": 150000.00,
  "trans_refund_total": 10000.00,
  "trans_status": 1,
  "trans_detail": [
    {
      "trans_detail_product_varian_id": "var-001",
      "trans_detail_qty": 2,
      "trans_detail_price": 25000.00,
      "trans_detail_subtotal": 50000.00
    },
    {
      "trans_detail_product_varian_id": "var-002",
      "trans_detail_qty": 3,
      "trans_detail_price": 30000.00,
      "trans_detail_subtotal": 90000.00
    }
  ]
}
```

**Response Success (201):**
```json
{
  "error": false,
  "message": "Transaksi berhasil dibuat",
  "data": {
    "trans_id": "trans-001",
    "trans_invoice": "INV-2025-001",
    "trans_total": 140000.00,
    "trans_create_at": "2025-11-20T19:16:02+07:00"
  }
}
```

**cURL Example:**
```bash
curl -X POST "http://127.0.0.1:3000/api/kasir/transaction" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "trans_user_id": "usr-001",
    "trans_invoice": "INV-2025-001",
    "trans_qty_total": 5,
    "trans_total": 140000.00,
    "trans_received_total": 150000.00
  }'
```

---

### 2. Get All Transactions

**Endpoint:** `GET /api/transaction`

**Description:** Mendapatkan semua transaksi

**Headers:**
```
Authorization: Bearer <JWT_TOKEN>
```

**Response Success (200):**
```json
{
  "error": false,
  "message": "Berhasil mengambil data transaksi",
  "data": [
    {
      "trans_id": "trans-001",
      "trans_user_id": "usr-001",
      "trans_customer_id": "cust-001",
      "trans_invoice": "INV-2025-001",
      "trans_qty_total": 5,
      "trans_product_total": 3,
      "trans_subtotal": 150000.00,
      "trans_discount": 10000.00,
      "trans_total": 140000.00,
      "trans_received_total": 150000.00,
      "trans_refund_total": 10000.00,
      "trans_status": 1,
      "trans_create_at": "2025-11-20T19:16:02+07:00",
      "customer": {
        "customer_id": "cust-001",
        "customer_name": "John Doe"
      },
      "user": {
        "user_id": "usr-001",
        "user_name": "Kasir 1"
      }
    }
  ]
}
```

---

### 3. Get Transaction by ID

**Endpoint:** `GET /api/kasir/transaction/{transId}`

**Description:** Mendapatkan detail transaksi berdasarkan ID

**Headers:**
```
Authorization: Bearer <JWT_TOKEN>
```

**Path Parameters:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| transId | string | Yes | Transaction ID |

**Response Success (200):**
```json
{
  "error": false,
  "message": "Berhasil mengambil data transaksi",
  "data": {
    "trans_id": "trans-001",
    "trans_invoice": "INV-2025-001",
    "trans_total": 140000.00,
    "trans_detail": [
      {
        "trans_detail_id": "detail-001",
        "trans_detail_product_varian_id": "var-001",
        "trans_detail_qty": 2,
        "trans_detail_price": 25000.00,
        "trans_detail_subtotal": 50000.00
      }
    ]
  }
}
```

**cURL Example:**
```bash
curl -X GET "http://127.0.0.1:3000/api/kasir/transaction/trans-001" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

### 4. Get Transaction Details

**Endpoint:** `GET /api/kasir/transaction_detail`

**Description:** Mendapatkan detail item transaksi

**Headers:**
```
Authorization: Bearer <JWT_TOKEN>
```

**Query Parameters:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| trans_id | string | Yes | Transaction ID |

**Response Success (200):**
```json
{
  "error": false,
  "message": "Berhasil mengambil detail transaksi",
  "data": [
    {
      "trans_detail_id": "detail-001",
      "trans_detail_parent_id": "trans-001",
      "trans_detail_product_varian_id": "var-001",
      "trans_detail_qty": 2,
      "trans_detail_price": 25000.00,
      "trans_detail_subtotal": 50000.00
    }
  ]
}
```

**cURL Example:**
```bash
curl -X GET "http://127.0.0.1:3000/api/kasir/transaction_detail?trans_id=trans-001" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

### 5. Get Transaction Summary

**Endpoint:** `GET /api/kasir/transaction_summary`

**Description:** Mendapatkan ringkasan dan statistik transaksi

**Headers:**
```
Authorization: Bearer <JWT_TOKEN>
```

**Response Success (200):**
```json
{
  "error": false,
  "message": "Berhasil mengambil ringkasan transaksi",
  "data": {
    "total_transactions": 150,
    "total_sales": 15750000.00,
    "total_discount": 500000.00,
    "average_transaction": 105000.00
  }
}
```

---

## 🗄️ Database Schema

### Table: users

**Description:** Tabel untuk menyimpan data user/kasir

| Column | Type | Nullable | Default | Description |
|--------|------|----------|---------|-------------|
| user_id | VARCHAR(36) | NO | - | Primary Key (UUID) |
| user_name | VARCHAR(100) | NO | - | Nama user |
| user_email | VARCHAR(100) | NO | - | Email user (unique) |
| user_password | VARCHAR(255) | NO | - | Password (hashed) |
| user_pegawai_id | VARCHAR(36) | NO | - | Foreign Key ke pegawai |
| user_has_mobile_access | INT | NO | 0 | 0=No, 1=Yes |
| user_role_id | VARCHAR(36) | NO | - | Foreign Key ke roles |
| user_create_at | TIMESTAMP | NO | CURRENT_TIMESTAMP | Waktu dibuat |
| user_update_at | TIMESTAMP | NO | CURRENT_TIMESTAMP | Waktu diupdate |

**Indexes:**
- `idx_user_email` on `user_email`
- `idx_user_pegawai_id` on `user_pegawai_id`
- `idx_user_role_id` on `user_role_id`

**Foreign Keys:**
- `user_pegawai_id` → `pegawai(pegawai_id)`
- `user_role_id` → `roles(role_id)`

---

### Table: transactions

**Description:** Tabel untuk menyimpan data transaksi kasir

| Column | Type | Nullable | Default | Description |
|--------|------|----------|---------|-------------|
| trans_id | VARCHAR(36) | NO | - | Primary Key (UUID) |
| trans_user_id | VARCHAR(36) | NO | - | Foreign Key ke users |
| trans_customer_id | VARCHAR(36) | YES | NULL | Foreign Key ke customers |
| trans_order_id | VARCHAR(36) | YES | NULL | Foreign Key ke customer_orders |
| trans_invoice | VARCHAR(50) | NO | - | Nomor invoice (unique) |
| trans_qty_total | INT | NO | - | Total quantity item |
| trans_product_total | INT | NO | - | Total jenis produk |
| trans_subtotal | DECIMAL(15,2) | NO | - | Subtotal sebelum diskon |
| trans_discount | DECIMAL(15,2) | NO | 0 | Total diskon |
| trans_total | DECIMAL(15,2) | NO | - | Total setelah diskon |
| trans_received_total | DECIMAL(15,2) | NO | - | Uang yang diterima |
| trans_refund_total | DECIMAL(15,2) | NO | 0 | Uang kembalian |
| trans_status | INT | NO | 1 | 0=Cancelled, 1=Completed |
| trans_create_at | TIMESTAMP | NO | CURRENT_TIMESTAMP | Waktu transaksi |

**Indexes:**
- `idx_trans_user_id` on `trans_user_id`
- `idx_trans_customer_id` on `trans_customer_id`
- `idx_trans_invoice` on `trans_invoice`
- `idx_trans_create_at` on `trans_create_at`
- `idx_trans_status` on `trans_status`

**Foreign Keys:**
- `trans_user_id` → `users(user_id)`
- `trans_customer_id` → `customers(customer_id)`
- `trans_order_id` → `customer_orders(order_id)`

---

### Table: transaction_details

**Description:** Tabel untuk menyimpan detail item transaksi

| Column | Type | Nullable | Default | Description |
|--------|------|----------|---------|-------------|
| trans_detail_id | VARCHAR(36) | NO | - | Primary Key (UUID) |
| trans_detail_parent_id | VARCHAR(36) | NO | - | Foreign Key ke transactions |
| trans_detail_product_varian_id | VARCHAR(36) | NO | - | Foreign Key ke product_varians |
| trans_detail_qty | INT | NO | - | Quantity item |
| trans_detail_price | DECIMAL(15,2) | NO | - | Harga per item |
| trans_detail_subtotal | DECIMAL(15,2) | NO | - | Subtotal (qty × price) |

**Indexes:**
- `idx_trans_detail_parent_id` on `trans_detail_parent_id`
- `idx_trans_detail_product_varian_id` on `trans_detail_product_varian_id`

**Foreign Keys:**
- `trans_detail_parent_id` → `transactions(trans_id)` ON DELETE CASCADE
- `trans_detail_product_varian_id` → `product_varians(product_varian_id)`

---

## 📝 Entity Structures

### Login Entity

```go
type Login struct {
    UserId             string `json:"user_id"`
    UserEmail          string `json:"user_email" validate:"required"`
    UserPassword       string `json:"user_password" validate:"required"`
    UserFcmToken       string `json:"user_fcmtoken"`
    UserDeviceMetadata string `json:"user_device_metadata"`
}
```

---

### Dashboard Entity

```go
type DashboardStats struct {
    TotalSales        float64            `json:"total_sales"`
    TotalTransactions int                `json:"total_transactions"`
    TotalCustomers    int                `json:"total_customers"`
    MonthlySales      []MonthlySalesData `json:"monthly_sales"`
    TopProducts       []TopProductData   `json:"top_products"`
}

type MonthlySalesData struct {
    Month      int     `json:"month"`
    MonthName  string  `json:"month_name"`
    Year       int     `json:"year"`
    TotalSales float64 `json:"total_sales"`
}

type TopProductData struct {
    ProductId     string  `json:"product_id"`
    ProductName   string  `json:"product_name"`
    VariantName   string  `json:"variant_name"`
    TotalQuantity int     `json:"total_quantity"`
    TotalSales    float64 `json:"total_sales"`
}
```

---

### Transaction Entity

```go
type Transaction struct {
    TransId            string               `json:"trans_id"`
    TransUserId        string               `json:"trans_user_id"`
    TransCustomerId    string               `json:"trans_customer_id"`
    TransOrderId       string               `json:"trans_order_id"`
    TransInvoice       string               `json:"trans_invoice"`
    TransQtyTotal      int                  `json:"trans_qty_total"`
    TransProductTotal  int                  `json:"trans_product_total"`
    TransSubtotal      float64              `json:"trans_subtotal"`
    TransDiscount      float64              `json:"trans_discount"`
    TransTotal         float64              `json:"trans_total"`
    TransReceivedTotal float64              `json:"trans_received_total"`
    TransRefundTotal   float64              `json:"trans_refund_total"`
    TransStatus        int                  `json:"trans_status"`
    TransCreateAt      time.Time            `json:"trans_create_at"`
    TransDetail        *[]TransactionDetail `json:"trans_detail"`
    Customer           *CustomerResponse    `json:"customer"`
    User               *User                `json:"user"`
}
```

---

## 🔒 Authentication

Semua endpoint (kecuali login) memerlukan Bearer token authentication:

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

Token didapatkan dari endpoint login dan harus disertakan di header setiap request.

---

## 📌 Status Codes

| Code | Description |
|------|-------------|
| 200 | OK - Request berhasil |
| 201 | Created - Resource berhasil dibuat |
| 400 | Bad Request - Request tidak valid |
| 401 | Unauthorized - Token tidak valid/expired |
| 404 | Not Found - Resource tidak ditemukan |
| 500 | Internal Server Error - Error di server |

---

## 🚀 Quick Start

1. **Login sebagai kasir:**
```bash
curl -X POST "http://127.0.0.1:3000/api/kasir/login" \
  -H "Content-Type: application/json" \
  -d '{"user_email":"kasir@example.com","user_password":"password123"}'
```

2. **Simpan JWT token dari response**

3. **Gunakan token untuk request lainnya:**
```bash
curl -X GET "http://127.0.0.1:3000/api/dashboard/stats" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

**Status:** ✅ Complete  
**Version:** 1.0  
**Last Updated:** November 20, 2025  
**Author:** E-Canteen Development Team
