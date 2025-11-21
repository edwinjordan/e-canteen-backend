# Pagination Response API

## Overview
Semua endpoint GET dengan list data (FindAll) sekarang menggunakan response format baru dengan informasi pagination lengkap.

## Response Format Baru

### Structure
```json
{
  "code": 200,
  "success": true,
  "message": "Success",
  "data": [...],
  "page": {
    "currentPage": 1,
    "totalPage": 10,
    "totalRows": 100,
    "perPage": 10
  }
}
```

### Response Fields

| Field | Type | Description |
|-------|------|-------------|
| `code` | int | HTTP status code (200, 400, 404, dll) |
| `success` | boolean | Status keberhasilan request |
| `message` | string | Pesan response |
| `data` | array | Data hasil query |
| `page` | object | Informasi pagination (null jika bukan list) |

### Page Object Fields

| Field | Type | Description |
|-------|------|-------------|
| `currentPage` | int | Halaman saat ini |
| `totalPage` | int | Total jumlah halaman |
| `totalRows` | int | Total jumlah data |
| `perPage` | int | Jumlah data per halaman |

## Query Parameters

Semua endpoint FindAll mendukung parameter berikut:

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `page` | int | 1 | Nomor halaman (mulai dari 1) |
| `limit` | int | 10 | Jumlah data per halaman |
| `offset` | int | calculated | Offset data (opsional, akan di-override jika `page` diset) |
| `search` | string | "" | Keyword pencarian |

## Cara Penggunaan

### 1. Menggunakan Parameter `page` (Recommended)

```bash
# Halaman pertama (default)
GET /api/products?page=1&limit=10

# Halaman kedua
GET /api/products?page=2&limit=10

# Halaman ketiga dengan 20 data per halaman
GET /api/products?page=3&limit=20
```

**Response:**
```json
{
  "code": 200,
  "success": true,
  "message": "Berhasil mengambil data",
  "data": [
    {
      "product_id": "prod-001",
      "product_name": "Nasi Goreng",
      ...
    },
    ...
  ],
  "page": {
    "currentPage": 2,
    "totalPage": 10,
    "totalRows": 95,
    "perPage": 10
  }
}
```

### 2. Menggunakan Parameter `offset` (Legacy)

```bash
# Data dari index 0-9
GET /api/products?offset=0&limit=10

# Data dari index 10-19
GET /api/products?offset=10&limit=10
```

**Note:** Jika `page` dan `offset` sama-sama diberikan, `offset` akan diabaikan dan `page` yang digunakan.

### 3. Dengan Search

```bash
GET /api/products?page=1&limit=10&search=nasi

GET /api/customers?page=1&limit=20&search=08123
```

## Endpoints yang Sudah Diupdate

### 1. Products
**GET** `/api/products`

**Query Parameters:**
- `page` (int, default: 1)
- `limit` (int, default: 10)
- `search` (string) - Search by product name
- `category_id` (string) - Filter by category

**Example:**
```bash
curl "http://127.0.0.1:3000/api/products?page=1&limit=10&search=nasi" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Response:**
```json
{
  "code": 200,
  "success": true,
  "message": "Berhasil mengambil data",
  "data": [
    {
      "product_id": "prod-001",
      "product_code": "PRD001",
      "product_name": "Nasi Goreng",
      "product_category_id": "cat-001",
      "category_name": "Makanan",
      "varian": [
        {
          "product_varian_id": "var-001",
          "varian_name": "Reguler",
          "product_varian_price": 15000,
          ...
        }
      ]
    }
  ],
  "page": {
    "currentPage": 1,
    "totalPage": 5,
    "totalRows": 47,
    "perPage": 10
  }
}
```

### 2. Customers
**GET** `/api/customer`

**Query Parameters:**
- `page` (int, default: 1)
- `limit` (int, default: 10)
- `search` (string) - Search by name or phone number
- `customer` (string) - Specific customer ID to prioritize

**Example:**
```bash
curl "http://127.0.0.1:3000/api/customer?page=1&limit=10&search=081" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**Response:**
```json
{
  "code": 200,
  "success": true,
  "message": "Berhasil mengambil data",
  "data": [
    {
      "customer_id": "cust-001",
      "customer_code": "CST001",
      "customer_name": "John Doe",
      "customer_phonenumber": "081234567890",
      "jurusan": {
        "major_name": "Teknik Informatika"
      },
      ...
    }
  ],
  "page": {
    "currentPage": 1,
    "totalPage": 8,
    "totalRows": 78,
    "perPage": 10
  }
}
```

## Perhitungan Pagination

### Formula
```javascript
// Dari page number ke offset
offset = (page - 1) * limit

// Total halaman
totalPage = Math.ceil(totalRows / limit)

// Halaman saat ini
currentPage = Math.floor(offset / limit) + 1
```

### Contoh Perhitungan
Jika total data = 95, limit = 10:
- Total halaman = ceil(95/10) = 10
- Page 1: offset 0-9 (data 1-10)
- Page 2: offset 10-19 (data 11-20)
- Page 10: offset 90-94 (data 91-95)

## Migrasi dari Response Lama

### Response Format Lama
```json
{
  "error": false,
  "message": "Success",
  "data": [...]
}
```

Header:
- `offset`: Next offset value
- `Access-Control-Expose-Headers`: offset

### Response Format Baru
```json
{
  "code": 200,
  "success": true,
  "message": "Success",
  "data": [...],
  "page": {
    "currentPage": 1,
    "totalPage": 10,
    "totalRows": 100,
    "perPage": 10
  }
}
```

### Breaking Changes
1. ❌ Header `offset` sudah tidak digunakan
2. ✅ Gunakan `page.currentPage` dan `page.perPage` untuk hitung next page
3. ✅ Field `error: false` diganti `success: true`
4. ✅ Tambahan field `code` untuk HTTP status

### Migration Guide

**Frontend (JavaScript/TypeScript):**
```javascript
// Lama
const nextOffset = response.headers.get('offset');
fetch(`/api/products?offset=${nextOffset}&limit=10`);

// Baru
const currentPage = response.data.page.currentPage;
const nextPage = currentPage + 1;
if (nextPage <= response.data.page.totalPage) {
  fetch(`/api/products?page=${nextPage}&limit=10`);
}
```

## File yang Dimodifikasi

### 1. Handler (Response Structure)
- **File**: `handler/handler.go`
- **Changes**: 
  - Tambah struct `WebResponseWithPagination`
  - Tambah struct `PageInfo`

### 2. Product Module
- **Files**:
  - `app/repository/repository.product.go` - Interface update
  - `repository/product_repository/repository.go` - Add total count query
  - `app/usecase/usecase_product/implement.go` - Use new response format

### 3. Customer Module
- **Files**:
  - `app/repository/repository.customer.go` - Interface update
  - `repository/customer_repository/repository.go` - Add total count query
  - `app/usecase/usecase_customer/implement.go` - Use new response format

## Testing

### Test dengan cURL

```bash
# Products - Page 1
curl -X GET "http://127.0.0.1:3000/api/products?page=1&limit=5" \
  -H "Authorization: Bearer YOUR_TOKEN"

# Products - Page 2 dengan search
curl -X GET "http://127.0.0.1:3000/api/products?page=2&limit=5&search=nasi" \
  -H "Authorization: Bearer YOUR_TOKEN"

# Customers - Page 1
curl -X GET "http://127.0.0.1:3000/api/customer?page=1&limit=10" \
  -H "Authorization: Bearer YOUR_TOKEN"

# Customers dengan search phone
curl -X GET "http://127.0.0.1:3000/api/customer?page=1&limit=10&search=0812" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### Verify Response
```javascript
{
  "code": 200,  // ✓ Should be 200
  "success": true,  // ✓ Should be true
  "message": "Berhasil mengambil data",  // ✓ Success message
  "data": [...],  // ✓ Array of data
  "page": {
    "currentPage": 1,  // ✓ Current page number
    "totalPage": 10,  // ✓ Total pages calculated
    "totalRows": 95,  // ✓ Total rows from DB
    "perPage": 10  // ✓ Items per page
  }
}
```

## Compatibility

### Backward Compatibility
- ✅ Parameter `offset` masih bisa digunakan
- ✅ Parameter `limit` tetap sama
- ❌ Response format berbeda (breaking change)

### Forward Compatibility
- ✅ Support parameter `page` baru
- ✅ Response selalu include pagination info
- ✅ Total count selalu akurat

---

**Status**: ✅ Implemented
**Version**: 2.0
**Date**: November 16, 2025
**Breaking Changes**: Yes - Response format changed
