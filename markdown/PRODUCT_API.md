# Product API Documentation

Complete API documentation for Product Management endpoints.

---

## Endpoints

### 1. Get All Products

**Endpoint**: `GET /api/products`

**Description**: Retrieve all products with their variants and pagination support

**Query Parameters**:
| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| limit | integer | No | 10 | Number of items per page |
| offset | integer | No | 0 | Offset for pagination |
| search | string | No | - | Search by product name |
| category_id | string | No | - | Filter by category ID |

**Response Headers**:
- `offset`: Next offset value for pagination
- `Access-Control-Expose-Headers`: offset

**Success Response** (200 OK):
```json
{
  "error": false,
  "message": "Success get data",
  "data": [
    {
      "product_id": "uuid-string",
      "product_code": "PRD001",
      "product_name": "Nasi Goreng",
      "product_category_id": "category-uuid",
      "product_desc": "Nasi goreng spesial",
      "category_name": "Makanan",
      "product_create_at": "2025-11-16T10:00:00Z",
      "product_update_at": "2025-11-16T10:00:00Z",
      "product_delete_at": null,
      "product_photo": "https://example.com/photo.jpg",
      "varian": [
        {
          "product_varian_id": "varian-uuid",
          "product_id": "uuid-string",
          "product_name": "Nasi Goreng",
          "varian_name": "Regular",
          "product_varian_price": 15000,
          "product_varian_qty_booth": 100,
          "product_varian_qty_warehouse": "0",
          "varian_id": "var001",
          "product_varian_qty_left": 100
        }
      ]
    }
  ]
}
```

**Example Request**:
```bash
curl -X GET "http://localhost:8080/api/products?limit=10&offset=0&search=nasi&category_id=cat123" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

### 2. Get Product by ID

**Endpoint**: `GET /api/products/{productId}`

**Description**: Retrieve a specific product by its ID with variants

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| productId | string | Yes | Product UUID |

**Success Response** (200 OK):
```json
{
  "error": false,
  "message": "Success get data",
  "data": {
    "product_id": "uuid-string",
    "product_code": "PRD001",
    "product_name": "Nasi Goreng",
    "product_category_id": "category-uuid",
    "product_desc": "Nasi goreng spesial",
    "category_name": "Makanan",
    "product_create_at": "2025-11-16T10:00:00Z",
    "product_update_at": "2025-11-16T10:00:00Z",
    "product_delete_at": null,
    "product_photo": "https://example.com/photo.jpg",
    "varian": null
  }
}
```

**Error Response** (404 Not Found):
```json
{
  "error": true,
  "message": "data tidak ditemukan"
}
```

**Example Request**:
```bash
curl -X GET "http://localhost:8080/api/products/uuid-string" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

---

### 3. Create Product with Variants

**Endpoint**: `POST /api/products`

**Description**: Create a new product along with its variants

**Request Body**:
```json
{
  "product_code": "PRD001",
  "product_name": "Nasi Goreng",
  "product_category_id": "category-uuid",
  "product_desc": "Nasi goreng spesial dengan bumbu rahasia",
  "product_photo": "https://example.com/photo.jpg",
  "varians": [
    {
      "varian_id": "var001",
      "product_varian_price": 15000,
      "product_varian_qty": 100
    },
    {
      "varian_id": "var002",
      "product_varian_price": 20000,
      "product_varian_qty": 50
    }
  ]
}
```

**Validation Rules**:
| Field | Type | Required | Rules |
|-------|------|----------|-------|
| product_code | string | Yes | - |
| product_name | string | Yes | - |
| product_category_id | string | Yes | Must be valid category ID |
| product_desc | string | No | - |
| product_photo | string | No | - |
| varians | array | Yes | Min 1 variant |
| varians[].varian_id | string | Yes | - |
| varians[].product_varian_price | integer | Yes | Min 0 |
| varians[].product_varian_qty | integer | Yes | Min 0 |

**Success Response** (200 OK):
```json
{
  "error": false,
  "message": "Success create data",
  "data": {
    "product_id": "generated-uuid",
    "message": "Product and variants created successfully"
  }
}
```

**Error Response** (400 Bad Request):
```json
{
  "error": true,
  "message": "Key: 'ProductInsertRequest.ProductName' Error:Field validation for 'ProductName' failed on the 'required' tag"
}
```

**Example Request**:
```bash
curl -X POST "http://localhost:8080/api/products" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "product_code": "PRD001",
    "product_name": "Nasi Goreng",
    "product_category_id": "category-uuid",
    "product_desc": "Nasi goreng spesial",
    "product_photo": "https://example.com/photo.jpg",
    "varians": [
      {
        "varian_id": "var001",
        "product_varian_price": 15000,
        "product_varian_qty": 100
      }
    ]
  }'
```

**Notes**:
- `product_id` and `product_varian_id` are auto-generated (UUID v4)
- Product and variants are inserted in a single transaction
- If any variant fails to insert, the entire transaction is rolled back
- XSS protection is applied to text fields (product_code, product_name, product_desc)

---

### 4. Update Product

**Endpoint**: `PUT /api/products/{productId}`

**Description**: Update a product's information

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| productId | string | Yes | Product UUID |

**Request Body**:
```json
{
  "product_code": "PRD001",
  "product_name": "Nasi Goreng Spesial",
  "product_category_id": "category-uuid",
  "product_desc": "Nasi goreng spesial dengan bumbu rahasia (updated)",
  "product_photo": "https://example.com/new-photo.jpg"
}
```

**Validation Rules**:
| Field | Type | Required | Rules |
|-------|------|----------|-------|
| product_code | string | Yes | - |
| product_name | string | Yes | - |
| product_category_id | string | Yes | Must be valid category ID |
| product_desc | string | No | - |
| product_photo | string | No | - |

**Success Response** (200 OK):
```json
{
  "error": false,
  "message": "Success update data",
  "data": {
    "product_id": "uuid-string",
    "message": "Product updated successfully"
  }
}
```

**Error Response** (404 Not Found):
```json
{
  "error": true,
  "message": "product tidak ditemukan"
}
```

**Example Request**:
```bash
curl -X PUT "http://localhost:8080/api/products/uuid-string" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "product_code": "PRD001",
    "product_name": "Nasi Goreng Spesial",
    "product_category_id": "category-uuid",
    "product_desc": "Updated description",
    "product_photo": "https://example.com/new-photo.jpg"
  }'
```

**Notes**:
- Only product information is updated, not variants
- `product_update_at` is automatically set to current timestamp
- XSS protection is applied to text fields
- To update variants, use the variant endpoints (not yet implemented)

---

### 5. Delete Product

**Endpoint**: `DELETE /api/products/{productId}`

**Description**: Soft delete a product by its ID

**Path Parameters**:
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| productId | string | Yes | Product UUID |

**Success Response** (200 OK):
```json
{
  "error": false,
  "message": "Success delete data",
  "data": {
    "product_id": "uuid-string",
    "message": "Product deleted successfully"
  }
}
```

**Error Response** (404 Not Found):
```json
{
  "error": true,
  "message": "product tidak ditemukan"
}
```

**Example Request**:
```bash
curl -X DELETE "http://localhost:8080/api/products/uuid-string" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Notes**:
- This is a **soft delete** operation
- Sets `product_delete_at` to current timestamp
- Product is not physically removed from database
- Deleted products are excluded from `GET /api/products` results
- Variants are NOT deleted (cascade delete not implemented)

---

## Error Codes

| HTTP Code | Error Type | Description |
|-----------|------------|-------------|
| 200 | Success | Request successful |
| 400 | Bad Request | Validation error or invalid input |
| 404 | Not Found | Product not found |
| 401 | Unauthorized | Missing or invalid JWT token |
| 500 | Internal Server Error | Server error |

---

## Database Tables

### Base Tables (for INSERT/UPDATE)
- `ms_product` - Product base table
- `ms_product_varian` - Product variant base table

### Views (for SELECT)
- `v_ms_product` - Product view with category join
- `v_ms_product_varian` - Variant view with product join

---

## Security Features

1. **JWT Authentication**: All endpoints require Bearer token
2. **XSS Protection**: HTML escaping for text inputs
3. **SQL Injection Protection**: Using parameterized queries
4. **Input Validation**: Using validator package
5. **Transaction Safety**: Atomic operations with rollback

---

## Related Files

- [router/router.product.go](router/router.product.go) - Route definitions
- [app/usecase/usecase_product/implement.go](app/usecase/usecase_product/implement.go) - Business logic
- [repository/product_repository/repository.go](repository/product_repository/repository.go) - Database operations
- [entity/entity.product.go](entity/entity.product.go) - Product entity
- [entity/entity.varian.go](entity/entity.varian.go) - Variant entity

---

## Swagger Documentation

All endpoints are documented with Swagger annotations. Generate Swagger docs with:

```bash
swag init
```

Access Swagger UI at: `http://localhost:8080/swagger/index.html`
