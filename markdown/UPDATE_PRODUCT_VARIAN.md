# Update Product dengan Varian

## Endpoint
**PUT** `/api/products/{productId}`

## Deskripsi
Update produk beserta varian-variannya. Endpoint ini mendukung:
- Update informasi produk
- Update varian yang sudah ada
- Menambah varian baru

## Authentication
Memerlukan Bearer token authentication.

## Path Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| productId | string | Yes | ID produk yang akan diupdate |

## Request Body

```json
{
  "product_code": "PRD001",
  "product_name": "Nasi Goreng Spesial",
  "product_category_id": "cat-001",
  "product_desc": "Nasi goreng dengan telur dan ayam",
  "product_photo": "http://example.com/photo.jpg",
  "varians": [
    {
      "product_varian_id": "var-001",
      "varian_id": "size-reguler",
      "varian_name": "Reguler",
      "product_varian_price": 15000,
      "product_varian_qty_booth": 50,
      "product_varian_qty_warehouse": 100
    },
    {
      "product_varian_id": "",
      "varian_id": "",
      "varian_name": "Jumbo",
      "product_varian_price": 20000,
      "product_varian_qty_booth": 30,
      "product_varian_qty_warehouse": 50
    }
  ]
}
```

## Request Body Fields

### Product Fields (Required)
- `product_code` (string, required): Kode produk
- `product_name` (string, required): Nama produk
- `product_category_id` (string, required): ID kategori produk
- `product_desc` (string, optional): Deskripsi produk
- `product_photo` (string, optional): URL foto produk

### Varians Fields (Optional)
- `varians` (array, optional): Array varian produk
  - `product_varian_id` (string, optional): ID varian (kosongkan untuk varian baru)
  - `varian_id` (string, optional): ID tipe varian (akan di-generate otomatis jika kosong)
  - `varian_name` (string, required): Nama varian (contoh: "Reguler", "Jumbo")
  - `product_varian_price` (number, required): Harga varian (min: 0)
  - `product_varian_qty_booth` (number, required): Stok di booth (min: 0)
  - `product_varian_qty_warehouse` (number, required): Stok di warehouse (min: 0)

## Cara Kerja Varian

### 1. Update Varian yang Sudah Ada
Jika `product_varian_id` terisi dan varian ditemukan di database, maka varian akan di-UPDATE.

**Contoh:**
```json
{
  "product_varian_id": "existing-var-001",
  "varian_id": "size-reguler",
  "varian_name": "Reguler",
  "product_varian_price": 16000,
  "product_varian_qty_booth": 60,
  "product_varian_qty_warehouse": 120
}
```

### 2. Tambah Varian Baru
Jika `product_varian_id` kosong atau tidak ditemukan, maka varian baru akan di-INSERT.

**Contoh:**
```json
{
  "product_varian_id": "",
  "varian_id": "",
  "varian_name": "Extra Large",
  "product_varian_price": 25000,
  "product_varian_qty_booth": 20,
  "product_varian_qty_warehouse": 40
}
```

### 3. Update Produk Saja (Tanpa Varian)
Jika `varians` tidak disertakan atau array kosong, hanya produk yang akan di-update.

**Contoh:**
```json
{
  "product_code": "PRD001",
  "product_name": "Nasi Goreng Spesial Updated",
  "product_category_id": "cat-001",
  "product_desc": "Deskripsi baru"
}
```

## Response Success

```json
{
  "error": false,
  "message": "Berhasil mengupdate data",
  "data": {
    "product_id": "prod-12345",
    "message": "Product and variants updated successfully"
  }
}
```

## Response Error

### Product Not Found
```json
{
  "error": true,
  "message": "product tidak ditemukan"
}
```

### Validation Error
```json
{
  "error": true,
  "message": "Key: 'ProductUpdateRequest.ProductName' Error:Field validation for 'ProductName' failed on the 'required' tag"
}
```

## Contoh Penggunaan

### 1. Update Product dan Varian yang Ada
```bash
curl -X PUT "http://127.0.0.1:3000/api/products/prod-12345" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "product_code": "PRD001",
    "product_name": "Nasi Goreng Spesial",
    "product_category_id": "cat-001",
    "product_desc": "Nasi goreng terbaik",
    "product_photo": "http://example.com/photo.jpg",
    "varians": [
      {
        "product_varian_id": "var-001",
        "varian_id": "size-reg",
        "varian_name": "Reguler",
        "product_varian_price": 15000,
        "product_varian_qty_booth": 50,
        "product_varian_qty_warehouse": 100
      }
    ]
  }'
```

### 2. Update Product dan Tambah Varian Baru
```bash
curl -X PUT "http://127.0.0.1:3000/api/products/prod-12345" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "product_code": "PRD001",
    "product_name": "Nasi Goreng Spesial",
    "product_category_id": "cat-001",
    "varians": [
      {
        "varian_name": "Jumbo",
        "product_varian_price": 20000,
        "product_varian_qty_booth": 30,
        "product_varian_qty_warehouse": 50
      }
    ]
  }'
```

### 3. Update Product Saja
```bash
curl -X PUT "http://127.0.0.1:3000/api/products/prod-12345" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "product_code": "PRD001",
    "product_name": "Nasi Goreng Spesial Updated",
    "product_category_id": "cat-001",
    "product_desc": "Deskripsi baru"
  }'
```

## Fitur Update

### Auto-Generate UUID
- Jika `product_varian_id` kosong, akan di-generate otomatis
- Jika `varian_id` kosong, akan di-generate otomatis

### Transaction Support
- Semua operasi (update product + update/insert varians) dilakukan dalam 1 transaksi
- Jika ada error, semua perubahan akan di-rollback

### Validasi
- Field `product_name`, `product_code`, `product_category_id` wajib diisi
- Setiap varian wajib memiliki `varian_name`, `product_varian_price`, qty booth dan warehouse
- Price dan quantity harus >= 0

## File yang Dimodifikasi

1. **app/repository/repository.product.go**
   - Update interface `Update()` untuk menerima parameter `varians`

2. **repository/product_repository/repository.go**
   - Implementasi logic update/insert varian dalam method `Update()`
   - Support check existing varian dan auto-generate UUID

3. **app/usecase/usecase_product/implement.go**
   - Update struct `ProductUpdateRequest` dengan field `Varians`
   - Update struct `VarianRequest` dengan field `ProductVarianId`
   - Logic untuk process varians dalam method `Update()`

---

**Status**: ✅ Complete and Running
**Version**: 1.0
**Created**: November 16, 2025
