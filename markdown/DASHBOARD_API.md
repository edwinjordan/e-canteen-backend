# Dashboard API Documentation

## Endpoint
**GET** `/api/dashboard/stats`

## Description
Get comprehensive dashboard statistics including:
- Total sales (all-time)
- Total transactions count
- Total customers count
- Monthly sales graph (grouped by month for a specific year)
- Top 10 best-selling products

## Authentication
Requires Bearer token authentication.

## Query Parameters
| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| year | int | No | Current year | Year for monthly sales data |

## Response Example

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
      },
      // ... months 3-12
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
      // ... top 10
    ]
  }
}
```

## Usage Examples

### Get current year dashboard
```bash
curl -X GET "http://127.0.0.1:3000/api/dashboard/stats" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Get specific year dashboard
```bash
curl -X GET "http://127.0.0.1:3000/api/dashboard/stats?year=2024" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Features

### 1. Total Sales
- Sums all transaction totals where `trans_status = 1`
- Returns 0 if no transactions found

### 2. Total Transactions
- Counts all completed transactions (`trans_status = 1`)

### 3. Total Customers
- Counts all active customers (`customer_status = 1`)

### 4. Monthly Sales Graph
- Returns 12 months of data (January to December)
- Fills missing months with 0 sales
- Groups by YEAR and MONTH
- Perfect for chart visualization

### 5. Top 10 Products
- Sorted by total quantity sold (descending)
- Includes product and variant information
- Shows both quantity and sales value
- Aggregates across all variants

## Files Created

1. **entity/entity.dashboard.go** - Dashboard data structures
2. **app/repository/repository.dashboard.go** - Repository interface
3. **repository/dashboard_repository/repository.go** - Implementation
4. **app/usecase/usecase_dashboard/interface.go** - UseCase interface
5. **app/usecase/usecase_dashboard/implement.go** - UseCase implementation
6. **router/router.dashboard.go** - Route registration

## Database Tables Used

- `transactions` - For sales and transaction counts
- `customers` - For customer count
- `transaction_details` - For top products
- `product_varians` - For variant information
- `products` - For product information

## Testing in Swagger UI

1. Open http://127.0.0.1:3000/swagger/index.html
2. Find "Dashboard" section
3. Expand "GET /dashboard/stats"
4. Click "Try it out"
5. Enter year (optional)
6. Click "Execute"
7. View response

---

**Status**: ✅ Complete and Running
**Version**: 1.0
**Created**: November 16, 2025
