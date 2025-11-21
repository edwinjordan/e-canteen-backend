# Swagger Quick Start Guide

## ✅ Setup Complete!

Your E-Canteen Cashier API now has full Swagger documentation support.

## 🚀 Accessing Swagger UI

1. **Start the server** (if not already running):
   ```bash
   go run main.go
   ```

2. **Open your browser** and navigate to:
   ```
   http://127.0.0.1:3000/swagger/index.html
   ```

3. You should see the Swagger UI with all API endpoints documented.

## 🔐 Testing with Authentication

### Step 1: Login

1. In Swagger UI, find the **Authentication** section
2. Click on `POST /api/kasir/login`
3. Click "Try it out"
4. Enter the credentials:
   ```json
   {
     "email": "admin@ecanteen.com",
     "password": "admin123"
   }
   ```
5. Click "Execute"
6. Copy the `token` from the response

### Step 2: Authorize

1. Click the **"Authorize"** button (lock icon) at the top of the page
2. In the "Value" field, enter:
   ```
   Bearer YOUR_COPIED_TOKEN_HERE
   ```
   (Replace `YOUR_COPIED_TOKEN_HERE` with the actual token)
3. Click "Authorize"
4. Click "Close"

### Step 3: Test Protected Endpoints

Now you can test any protected endpoint:

1. Find the endpoint you want to test (e.g., `GET /api/products`)
2. Click "Try it out"
3. Fill in any required parameters
4. Click "Execute"
5. View the response

## 📝 Available Test Credentials

### Admin Users
- **Email**: `admin@ecanteen.com` | **Password**: `admin123` | **Role**: Super Admin
- **Email**: `siti@ecanteen.com` | **Password**: `admin123` | **Role**: Admin

### Cashier Users
- **Email**: `budi@ecanteen.com` | **Password**: `admin123` | **Role**: Cashier
- **Email**: `dewi@ecanteen.com` | **Password**: `admin123` | **Role**: Cashier

### Customer Accounts
All have password: `customer123`
- `andi@student.com`
- `rina@student.com`
- `joko@student.com`
- `lilis@student.com`
- `agus@student.com`

## 🔄 Regenerating Documentation

If you modify API endpoints or add new ones:

```bash
# Windows
swag init -g main.go --parseDependency --parseInternal

# Linux/Mac
swag init -g main.go --parseDependency --parseInternal

# Using Makefile
make swagger
```

Then restart the server:
```bash
go run main.go
```

## 📚 Key Endpoints to Try

### Authentication
- `POST /api/kasir/login` - Login (no auth required)
- `GET /api/kasir/version` - Get version (no auth required)

### Products (requires auth)
- `GET /api/products` - Get all products

### Transactions (requires auth)
- `POST /api/kasir/transaction` - Create transaction
- `GET /api/transaction` - Get all transactions
- `GET /api/kasir/transaction_summary` - Get summary

### Customers (requires auth)
- `GET /api/customers` - Get all customers

## 🐛 Troubleshooting

### Swagger UI shows blank page
- Clear browser cache
- Check browser console for errors
- Verify server is running on `127.0.0.1:3000`

### "401 Unauthorized" errors
- Make sure you've clicked "Authorize" and entered a valid token
- Token must start with `Bearer ` (with a space)
- Token expires after some time - get a new one by logging in again

### Documentation not updating
1. Stop the server (Ctrl+C)
2. Delete generated files:
   ```bash
   rm docs/swagger.json docs/swagger.yaml
   ```
3. Regenerate:
   ```bash
   swag init -g main.go --parseDependency --parseInternal
   ```
4. Restart server

### Server panics on startup
- Check if port 3000 is already in use
- Verify database connection in `.env` file
- Ensure all dependencies are installed: `go mod download`

## 📖 Swagger Features

### Try It Out
Click "Try it out" on any endpoint to test it directly from the browser.

### Schemas
Click on "Schemas" at the bottom to see all data models and their structure.

### Export
You can download the OpenAPI specification:
- JSON: `http://127.0.0.1:3000/swagger/doc.json`
- YAML: `http://127.0.0.1:3000/swagger/doc.yaml`

### Use in Postman
Import the JSON/YAML file into Postman for API testing.

## ✨ Tips

1. **Use realistic data** when testing to see actual responses
2. **Check response schemas** to understand the data structure
3. **Explore all endpoints** to understand the API capabilities
4. **Use the search** feature to quickly find endpoints
5. **Bookmark** the Swagger URL for easy access

## 🎉 Success!

Your API documentation is now live and accessible. Share this URL with your team:

```
http://127.0.0.1:3000/swagger/index.html
```

For production, update the host in `main.go` to your production domain.

---

**Need Help?**
- Check the main [README.md](../README.md)
- See [docs/README.md](README.md) for detailed documentation
- View [DATABASE_SCHEMA.md](../DATABASE_SCHEMA.md) for database info
