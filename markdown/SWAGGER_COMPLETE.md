# Swagger API Documentation - Complete Setup

## ✅ Status: SUCCESSFULLY CONFIGURED

Your Swagger documentation is now fully functional with all API endpoints visible!

## 📍 Access Swagger UI

Open your browser to: **http://127.0.0.1:3000/swagger/index.html**

## 🎯 Available API Endpoints

### Authentication
- `POST /api/kasir/login` - Kasir login
- `PUT /api/kasir/logout` - Kasir logout  
- `POST /api/customer/login` - Customer login
- `POST /api/customer/logout` - Customer logout

### Customer Management
- `GET /api/customer` - Get all customers
- `GET /api/customer/{customerId}` - Get customer by ID
- `POST /api/customer` - Register new customer
- `PUT /api/customer/{customerId}` - Update customer
- `DELETE /api/customer/{customerId}` - Delete customer

### Products
- `GET /api/products` - Get all products with variants

### Transactions
- `POST /api/kasir/transaction` - Create new transaction
- `GET /api/transaction` - Get all transactions
- `GET /api/kasir/transaction/{transId}` - Get transaction by ID
- `GET /api/kasir/transaction_detail` - Get transaction details

### Shopping Cart
- `POST /api/tempcart` - Add item to cart
- `PUT /api/tempcart/{productVarianId}/{userId}` - Update cart item
- `DELETE /api/tempcart/{productVarianId}/{userId}` - Remove from cart

### Version Info
- `GET /api/kasir/version` - Get admin app version
- `GET /api/shop/version` - Get shop app version

## 🔐 Authentication Setup

1. **Login First**: Use the `/api/kasir/login` or `/api/customer/login` endpoint
   
   Example request body:
   ```json
   {
     "username": "your_username",
     "password": "your_password"
   }
   ```

2. **Get JWT Token**: Copy the token from the response

3. **Authorize**: Click the 🔒 "Authorize" button at the top of Swagger UI

4. **Enter Token**: Type `Bearer YOUR_TOKEN_HERE` (include the word "Bearer" with a space)

5. **Test Protected Endpoints**: Now you can test endpoints marked with 🔒

## 📝 How to Use Swagger UI

1. **Expand Endpoint**: Click on any endpoint to expand it
2. **Try it out**: Click the "Try it out" button
3. **Fill Parameters**: Enter required parameters and request body
4. **Execute**: Click "Execute" to send the request
5. **View Response**: See the response code, body, and headers below

## 🔄 Regenerating Swagger Docs

After adding new endpoints or modifying existing ones:

1. Add route annotations in `routes.go` file following this pattern:
   ```go
   // @Summary Short description
   // @Description Detailed description
   // @Tags TagName
   // @Accept json
   // @Produce json
   // @Security BearerAuth
   // @Param paramName path/query/body string true "Description"
   // @Success 200 {object} handler.WebResponse "Success message"
   // @Failure 400 {object} handler.WebResponse "Error message"
   // @Router /endpoint/path [get/post/put/delete]
   func functionName() {}
   ```

2. Run the regeneration command:
   ```powershell
   swag init -g main.go --parseDependency --parseInternal
   ```

3. Fix the generated `docs/docs.go` by removing `LeftDelim` and `RightDelim` lines:
   ```powershell
   # Find lines around 1400-1402 and remove these two lines:
   # LeftDelim:        "{{",
   # RightDelim:       "}}",
   ```

4. Restart the server:
   ```powershell
   go run main.go
   ```

## 🛠️ Files Modified

- `main.go` - Added Swagger imports and configuration
- `routes.go` - Contains all route annotations (NEW FILE)
- `docs/docs.go` - Auto-generated Swagger specification
- `docs/swagger.json` - JSON format API spec
- `docs/swagger.yaml` - YAML format API spec
- `middleware/middleware.auth.go` - Added bypass for `/swagger/*` routes
- `router/*.go` - Added inline documentation comments

## 📚 Swagger Annotations Reference

### Common Tags
- `@Summary` - Short one-line description
- `@Description` - Detailed description
- `@Tags` - Group endpoints (e.g., "Customer", "Products")
- `@Accept` - Input format (json, xml, etc.)
- `@Produce` - Output format (json, xml, etc.)
- `@Security` - Security scheme (BearerAuth)

### Parameters
- `@Param name path type required "description"`
- `@Param name query type required "description"`
- `@Param name body type required "description"`

### Responses
- `@Success code {type} description`
- `@Failure code {type} description`
- `@Router /path [method]`

## 🎨 Response Types

All responses follow this structure:
```go
type WebResponse struct {
    Code   int         `json:"code"`
    Status string      `json:"status"`
    Data   interface{} `json:"data,omitempty"`
}
```

## 🚀 Quick Test Example

1. Test login endpoint:
   - Endpoint: `POST /api/kasir/login`
   - Body: `{"username": "admin", "password": "password"}`
   - Response: JWT token

2. Use token to test protected endpoint:
   - Click 🔒 Authorize
   - Enter: `Bearer <your-token>`
   - Test: `GET /api/products`

## 📖 Additional Resources

- Swagger Official Docs: https://swagger.io/docs/
- Swaggo GitHub: https://github.com/swaggo/swag
- Swaggo Declarative Comments: https://github.com/swaggo/swag#declarative-comments-format

---

**Last Updated**: November 14, 2025  
**API Version**: 1.2.0  
**Server**: http://127.0.0.1:3000
