# Order & Cart API Documentation

## Database Schema

### 1. Table: tb_customer_order

Tabel untuk menyimpan data order pelanggan.

| Column Name | Data Type | Description | Constraints |
|-------------|-----------|-------------|-------------|
| order_id | VARCHAR(36) | ID unik order (UUID) | PRIMARY KEY |
| order_customer_id | VARCHAR(36) | ID customer yang melakukan order | FOREIGN KEY -> tb_customer.customer_id |
| order_inv_number | VARCHAR(50) | Nomor invoice (format: ORN-DDMMYY-XXX) | UNIQUE, NOT NULL |
| order_address_id | VARCHAR(36) | ID alamat pengiriman | FOREIGN KEY -> tb_customer_address.address_id |
| order_delivery_type | VARCHAR(20) | Tipe pengiriman (delivery/pickup) | NOT NULL |
| order_total_item | INT | Total jumlah item dalam order | DEFAULT 0 |
| order_subtotal | DECIMAL(15,2) | Subtotal harga sebelum diskon | DEFAULT 0 |
| order_discount | DECIMAL(15,2) | Total diskon | DEFAULT 0 |
| order_total | DECIMAL(15,2) | Total harga setelah diskon | DEFAULT 0 |
| order_notes | TEXT | Catatan order dari customer | NULLABLE |
| order_status | INT | Status order (1=selesai, 2=diproses, 3=dibatalkan, 4=pending) | DEFAULT 4 |
| order_cancel_notes | TEXT | Catatan pembatalan order | NULLABLE |
| order_processed_datetime | DATETIME | Waktu order diproses | NULLABLE |
| order_processed_by | VARCHAR(36) | ID user yang memproses | NULLABLE |
| order_finished_datetime | DATETIME | Waktu order selesai | NULLABLE |
| order_finished_by | VARCHAR(36) | ID user yang menyelesaikan | NULLABLE |
| order_create_at | DATETIME | Waktu order dibuat | AUTO_GENERATE |

**Order Status:**
- `1` = Order selesai
- `2` = Order sedang diproses
- `3` = Order dibatalkan
- `4` = Order pending (baru dibuat)

---

### 2. Table: tb_customer_order_detail

Tabel untuk menyimpan detail item dalam order.

| Column Name | Data Type | Description | Constraints |
|-------------|-----------|-------------|-------------|
| order_detail_id | VARCHAR(36) | ID unik order detail (UUID) | PRIMARY KEY |
| order_detail_parent_id | VARCHAR(36) | ID order induk | FOREIGN KEY -> tb_customer_order.order_id |
| order_detail_product_varian_id | VARCHAR(36) | ID varian produk yang dipesan | FOREIGN KEY -> tb_product_varian.varian_id |
| order_detail_qty | INT | Jumlah/quantity produk | NOT NULL, DEFAULT 1 |
| order_detail_price | DECIMAL(15,2) | Harga satuan produk | NOT NULL |
| order_detail_subtotal | DECIMAL(15,2) | Subtotal (qty × price) | NOT NULL |

---

### 3. Table: temp_cart

Tabel untuk menyimpan keranjang belanja sementara sebelum checkout.

| Column Name | Data Type | Description | Constraints |
|-------------|-----------|-------------|-------------|
| temp_cart_id | VARCHAR(36) | ID unik cart (UUID) | PRIMARY KEY |
| temp_cart_order_id | VARCHAR(36) | ID order (untuk tracking) | NULLABLE |
| temp_cart_product_varian_id | VARCHAR(36) | ID varian produk | FOREIGN KEY -> tb_product_varian.varian_id |
| temp_cart_user_id | VARCHAR(36) | ID user pemilik cart | NOT NULL |
| temp_cart_qty | INT | Jumlah/quantity produk | NOT NULL, DEFAULT 1 |

---

## API Routes

### Order Routes

#### 1. Create Order
```
POST /api/order
```
**Description:** Membuat order baru dari temp cart dan memproses pembayaran

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "order_customer_id": "uuid-customer",
  "order_address_id": "uuid-address",
  "order_delivery_type": "delivery",
  "order_notes": "Catatan order opsional"
}
```

**Response Success (201):**
```json
{
  "code": 201,
  "status": "Created",
  "data": {
    "order_id": "uuid",
    "order_inv_number": "ORN-151125-001",
    "order_total": 150000,
    "order_status": 4
  }
}
```

---

#### 2. Get All Orders
```
GET /api/order
```
**Description:** Mengambil semua data order

**Headers:**
```
Authorization: Bearer <token>
```

**Query Parameters:**
- `status` (optional): Filter berdasarkan status (1,2,3,4)
- `limit` (optional): Jumlah data per halaman
- `offset` (optional): Offset untuk pagination

**Response Success (200):**
```json
{
  "code": 200,
  "status": "OK",
  "data": [
    {
      "order_id": "uuid",
      "order_inv_number": "ORN-151125-001",
      "customer": {
        "customer_name": "John Doe"
      },
      "order_total": 150000,
      "order_status": 4,
      "order_create_at": "2025-11-15T10:30:00Z"
    }
  ]
}
```

---

#### 3. Get Order By ID
```
GET /api/order/{orderId}
```
**Description:** Mengambil detail order berdasarkan ID (dengan relasi customer, address, dan order detail)

**Headers:**
```
Authorization: Bearer <token>
```

**Path Parameters:**
- `orderId`: ID order (UUID)

**Response Success (200):**
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "order_id": "uuid",
    "order_inv_number": "ORN-151125-001",
    "customer": {
      "customer_id": "uuid",
      "customer_name": "John Doe"
    },
    "address": {
      "address_id": "uuid",
      "address_detail": "Jl. Example No. 123"
    },
    "order_total_item": 3,
    "order_subtotal": 150000,
    "order_discount": 0,
    "order_total": 150000,
    "order_status": 4,
    "order_detail": [
      {
        "order_detail_id": "uuid",
        "order_detail_product_varian_id": "uuid-varian",
        "order_detail_qty": 2,
        "order_detail_price": 50000,
        "order_detail_subtotal": 100000
      }
    ]
  }
}
```

---

#### 4. Get Order Detail
```
GET /api/order_detail
```
**Description:** Mengambil data order detail

**Headers:**
```
Authorization: Bearer <token>
```

**Response Success (200):**
```json
{
  "code": 200,
  "status": "OK",
  "data": [
    {
      "order_detail_id": "uuid",
      "order_detail_parent_id": "uuid-order",
      "order_detail_product_varian_id": "uuid-varian",
      "order_detail_qty": 2,
      "order_detail_price": 50000,
      "order_detail_subtotal": 100000
    }
  ]
}
```

---

#### 5. Cancel Order
```
PUT /api/order_canceled/{orderId}
```
**Description:** Membatalkan order (mengubah status menjadi 3)

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Path Parameters:**
- `orderId`: ID order yang akan dibatalkan

**Request Body:**
```json
{
  "order_cancel_notes": "Alasan pembatalan"
}
```

**Response Success (200):**
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "order_id": "uuid",
    "order_status": 3,
    "order_cancel_notes": "Alasan pembatalan"
  }
}
```

---

#### 6. Process Order (Kasir Only)
```
PUT /api/kasir/order_processed/{orderId}
```
**Description:** Memproses order oleh kasir (mengubah status dari pending ke diproses)

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Path Parameters:**
- `orderId`: ID order yang akan diproses

**Response Success (200):**
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "order_id": "uuid",
    "order_status": 2,
    "order_processed_datetime": "2025-11-15T10:35:00Z",
    "order_processed_by": "uuid-kasir"
  }
}
```

---

### Temp Cart Routes

#### 1. Add to Cart
```
POST /api/tempcart
```
**Description:** Menambahkan produk ke keranjang sementara

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "temp_cart_product_varian_id": "uuid-varian",
  "temp_cart_user_id": "uuid-user",
  "temp_cart_qty": 2
}
```

**Response Success (201):**
```json
{
  "code": 201,
  "status": "Created",
  "data": {
    "temp_cart_id": "uuid",
    "temp_cart_product_varian_id": "uuid-varian",
    "temp_cart_user_id": "uuid-user",
    "temp_cart_qty": 2
  }
}
```

---

#### 2. Update Cart Item
```
PUT /api/tempcart/{productVarianId}/{userId}
```
**Description:** Mengupdate quantity item di keranjang

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Path Parameters:**
- `productVarianId`: ID varian produk
- `userId`: ID user

**Request Body:**
```json
{
  "temp_cart_qty": 5
}
```

**Response Success (200):**
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "temp_cart_id": "uuid",
    "temp_cart_qty": 5
  }
}
```

---

#### 3. Remove from Cart
```
DELETE /api/tempcart/{productVarianId}/{userId}
```
**Description:** Menghapus item dari keranjang

**Headers:**
```
Authorization: Bearer <token>
```

**Path Parameters:**
- `productVarianId`: ID varian produk yang akan dihapus
- `userId`: ID user

**Response Success (200):**
```json
{
  "code": 200,
  "status": "OK",
  "message": "Item berhasil dihapus dari keranjang"
}
```

---

## Business Logic Flow

### 1. Shopping Cart Flow
```
1. Customer menambahkan produk ke temp_cart (POST /api/tempcart)
2. Customer dapat update qty atau hapus item (PUT/DELETE /api/tempcart)
3. Customer melakukan checkout, membuat order dari temp_cart (POST /api/order)
4. Setelah order berhasil, temp_cart dikosongkan
```

### 2. Order Processing Flow
```
1. Order dibuat dengan status 4 (pending)
2. Kasir memproses order (PUT /api/kasir/order_processed/{orderId})
   - Status berubah menjadi 2 (diproses)
   - order_processed_datetime dan order_processed_by diisi
3. Order dapat dibatalkan dengan status 3 (PUT /api/order_canceled/{orderId})
4. Order selesai dengan status 1
   - order_finished_datetime dan order_finished_by diisi
```

### 3. Invoice Generation
Format invoice: `ORN-DDMMYY-XXX`
- `ORN` = Order Number prefix
- `DDMMYY` = Tanggal (15 November 2025 = 151125)
- `XXX` = Sequential number per hari (001, 002, dst.)

---

## Notes

1. **Auto-generated Fields:**
   - `order_id`, `order_detail_id`, `temp_cart_id` → UUID otomatis
   - `order_create_at` → Timestamp saat order dibuat
   - `order_inv_number` → Format ORN-DDMMYY-XXX

2. **Relationships:**
   - Order memiliki relasi dengan Customer, Address, dan OrderDetail
   - OrderDetail memiliki relasi dengan Order (parent) dan ProductVarian
   - TempCart memiliki relasi dengan User dan ProductVarian

3. **Authorization:**
   - Semua endpoint memerlukan Bearer token
   - Endpoint `/api/kasir/*` hanya dapat diakses oleh role kasir

4. **Validation:**
   - Quantity harus > 0
   - Product varian harus tersedia dan aktif
   - Customer harus memiliki address untuk delivery type
