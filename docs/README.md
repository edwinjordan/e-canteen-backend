# Swagger API Documentation

This directory contains Swagger/OpenAPI documentation for the E-Canteen Cashier API.

## Quick Start

### Prerequisites

- Go 1.18 or higher
- `swag` CLI tool

### Setup (Windows)

```powershell
# Run the setup script
.\setup-swagger.ps1

# Or manually:
go install github.com/swaggo/swag/cmd/swag@latest
go mod download
swag init -g main.go --parseDependency --parseInternal
```

### Setup (Linux/Mac)

```bash
# Run the setup script
chmod +x setup-swagger.sh
./setup-swagger.sh

# Or manually:
go install github.com/swaggo/swag/cmd/swag@latest
go mod download
swag init -g main.go --parseDependency --parseInternal
```

### Using Makefile

```bash
# Install dependencies and swag CLI
make install

# Generate Swagger documentation
make swagger

# Generate swagger and run the app
make dev

# Run the application
make run
```

## Accessing Swagger UI

After running the application:

1. Start the API server:
   ```bash
   go run main.go
   ```

2. Open your browser and navigate to:
   ```
   http://127.0.0.1:3000/swagger/index.html
   ```

## Documentation Structure

```
docs/
├── docs.go          # Main Swagger configuration
├── models.go        # Data model definitions
├── endpoints.go     # API endpoint annotations
└── swagger.json     # Generated Swagger specification (auto-generated)
└── swagger.yaml     # Generated Swagger specification (auto-generated)
```

## Regenerating Documentation

Whenever you update API endpoints or models, regenerate the documentation:

```bash
# Windows
swag init -g main.go --parseDependency --parseInternal

# Linux/Mac
swag init -g main.go --parseDependency --parseInternal

# Using Makefile
make swagger
```

## API Documentation Overview

### Available Endpoints

#### Authentication
- `POST /api/kasir/login` - User login
- `PUT /api/kasir/logout` - User logout

#### Transactions
- `POST /api/kasir/transaction` - Create transaction
- `GET /api/transaction` - Get all transactions
- `GET /api/kasir/transaction/{transId}` - Get transaction by ID
- `GET /api/kasir/transaction_detail` - Get transaction details
- `GET /api/kasir/transaction_summary` - Get transaction summary

#### Products
- `GET /api/products` - Get all products

#### Categories
- `GET /api/categories` - Get all categories

#### Variants
- `GET /api/variants` - Get all variants

#### Customers
- `GET /api/customers` - Get all customers
- `GET /api/customers/{customerId}` - Get customer by ID

#### Addresses
- `GET /api/customer_address` - Get customer addresses

#### Cart
- `GET /api/temp_cart` - Get cart items
- `POST /api/temp_cart` - Add to cart
- `PUT /api/temp_cart/{cartId}` - Update cart item
- `DELETE /api/temp_cart/{cartId}` - Delete cart item
- `DELETE /api/temp_cart/clear` - Clear cart

#### Orders
- `GET /api/orders` - Get all orders
- `GET /api/orders/{orderId}` - Get order by ID

#### Majors
- `GET /api/majors` - Get all majors

#### Version
- `GET /api/kasir/version` - Get admin app version
- `GET /api/shop/version` - Get shop app version
- `GET /api/check_maintenance_mode/{confCode}` - Check maintenance mode

## Authentication

Most endpoints require authentication using JWT Bearer token.

### How to authenticate in Swagger UI:

1. Click the "Authorize" button (lock icon) at the top
2. Enter: `Bearer YOUR_JWT_TOKEN`
3. Click "Authorize"

### Getting a token:

1. Use the `/api/kasir/login` endpoint
2. Use credentials:
   - Email: `admin@ecanteen.com`
   - Password: `admin123`
3. Copy the token from the response
4. Use it in the Authorization header

## Response Format

All endpoints return responses in the following format:

```json
{
  "error": false,
  "message": "Success",
  "data": {}
}
```

### Success Response
```json
{
  "error": false,
  "message": "Data retrieved successfully",
  "data": { ... }
}
```

### Error Response
```json
{
  "error": true,
  "message": "Error description",
  "data": null
}
```

## Adding New Endpoints

To add documentation for a new endpoint:

1. Add the endpoint function in `docs/endpoints.go`
2. Add Swagger annotations above the function
3. Regenerate documentation: `swag init -g main.go`

Example:

```go
// GetExample godoc
// @Summary Get example data
// @Description Get example data with parameters
// @Tags Example
// @Produce json
// @Security BearerAuth
// @Param id path string true "Example ID"
// @Success 200 {object} WebResponse{data=ExampleModel}
// @Failure 404 {object} WebResponse
// @Router /example/{id} [get]
func GetExampleEndpoint() {}
```

## Swagger Annotations Reference

### Common Annotations

- `@Summary` - Short description
- `@Description` - Detailed description
- `@Tags` - Group endpoints by tags
- `@Accept` - Request content type (json, xml, etc.)
- `@Produce` - Response content type
- `@Param` - Parameter definition
- `@Success` - Success response
- `@Failure` - Error response
- `@Router` - Endpoint path and method
- `@Security` - Security requirement

### Parameter Types

- `path` - URL path parameter
- `query` - Query string parameter
- `header` - HTTP header
- `body` - Request body
- `formData` - Form data

## Troubleshooting

### Swagger docs not updating

1. Delete generated files:
   ```bash
   rm -rf docs/swagger.*
   ```
2. Regenerate:
   ```bash
   swag init -g main.go --parseDependency --parseInternal
   ```

### "swag command not found"

Install swag CLI:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Make sure `$GOPATH/bin` is in your PATH.

### Documentation not showing in browser

1. Ensure the server is running
2. Check the correct URL: `http://127.0.0.1:3000/swagger/index.html`
3. Clear browser cache
4. Check console for errors

## Additional Resources

- [Swag Documentation](https://github.com/swaggo/swag)
- [Swagger/OpenAPI Specification](https://swagger.io/specification/)
- [Go Swagger Tutorial](https://github.com/swaggo/swag/blob/master/README.md)

## Support

For issues or questions about the API documentation, please contact:
- Email: support@ecanteen.com
- Repository: [GitHub Issues](https://github.com/edwinjordan/e-canteen-backend/issues)
