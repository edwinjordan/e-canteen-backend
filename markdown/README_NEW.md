# E-Canteen Cashier API

A comprehensive RESTful API for canteen management system built with Go, providing endpoints for product management, transactions, customer orders, and more.

## 🚀 Features

- **User Authentication** - JWT-based authentication system
- **Product Management** - CRUD operations for products, categories, and variants
- **Transaction Management** - Cashier transactions and sales tracking
- **Order Management** - Customer order processing and tracking
- **Customer Management** - Customer profiles and address management
- **Stock Management** - Inventory tracking for booth and warehouse
- **Swagger Documentation** - Interactive API documentation
- **Database Migrations** - Automated database schema setup
- **Data Seeders** - Sample data for testing

## 📋 Prerequisites

- Go 1.18 or higher
- MySQL 5.7+ or MariaDB 10.2+
- Git

## 🔧 Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/edwinjordan/e-canteen-backend.git
   cd e-canteen-cashier-api
   ```

2. **Install dependencies**
   ```bash
   go mod download
   go mod tidy
   ```

3. **Configure environment variables**
   
   Update `.env` with your configuration:
   ```env
   HOST_ADDR=127.0.0.1:3000
   DB_DRIVER=mysql
   DB_NAME=e-canteen_new
   DB_USERNAME=root
   DB_HOST=localhost
   DB_PORT=3306
   DB_PASSWORD=
   ```

4. **Setup database**
   
   **Windows:**
   ```powershell
   cd database
   .\migrate.ps1
   .\seed.ps1
   ```

   **Linux/Mac:**
   ```bash
   cd database
   chmod +x migrate.sh seed.sh
   ./migrate.sh
   ./seed.sh
   ```

   Or using Makefile:
   ```bash
   make db-setup       # Linux/Mac
   make db-setup-win   # Windows
   ```

5. **Setup Swagger documentation**
   
   **Windows:**
   ```powershell
   .\setup-swagger.ps1
   ```

   **Linux/Mac:**
   ```bash
   chmod +x setup-swagger.sh
   ./setup-swagger.sh
   ```

   Or using Makefile:
   ```bash
   make install
   make swagger
   ```

## 🏃 Running the Application

```bash
# Run the application
go run main.go

# Or using Makefile
make run

# Or build and run
make build
./bin/ecanteen-api
```

The API will be available at `http://127.0.0.1:3000`

## 📚 API Documentation

Access the interactive Swagger documentation at:

```
http://127.0.0.1:3000/swagger/index.html
```

### Default Credentials

**Admin User:**
- Email: `admin@ecanteen.com`
- Password: `admin123`
- Role: Super Admin

**Cashier User:**
- Email: `budi@ecanteen.com`
- Password: `admin123`
- Role: Cashier

**Customer:**
- Email: `andi@student.com`
- Password: `customer123`

## 🔑 Authentication

Most endpoints require authentication using JWT Bearer token.

1. **Login** via `POST /api/kasir/login`
2. **Get token** from the response
3. **Include token** in subsequent requests:
   ```
   Authorization: Bearer YOUR_JWT_TOKEN
   ```

## 📖 API Endpoints

### Authentication
- `POST /api/kasir/login` - User login
- `PUT /api/kasir/logout` - User logout

### Products
- `GET /api/products` - Get all products

### Categories
- `GET /api/categories` - Get all categories

### Transactions
- `POST /api/kasir/transaction` - Create transaction
- `GET /api/transaction` - Get all transactions
- `GET /api/kasir/transaction/{transId}` - Get transaction by ID
- `GET /api/kasir/transaction_detail` - Get transaction details
- `GET /api/kasir/transaction_summary` - Get transaction summary

### Customers
- `GET /api/customers` - Get all customers
- `GET /api/customers/{customerId}` - Get customer by ID

### Orders
- `GET /api/orders` - Get all orders
- `GET /api/orders/{orderId}` - Get order by ID

### Cart
- `GET /api/temp_cart` - Get cart items
- `POST /api/temp_cart` - Add to cart
- `PUT /api/temp_cart/{cartId}` - Update cart item
- `DELETE /api/temp_cart/{cartId}` - Delete cart item

### Version
- `GET /api/kasir/version` - Get admin app version
- `GET /api/shop/version` - Get shop app version

For complete API documentation, visit the Swagger UI.

## 🗄️ Database Schema

The database schema includes the following main tables:

- **Users & Authentication** - users, roles, user_logs, user_otp
- **Customers** - customers, customer_address, majors
- **Products** - products, categories, product_varians, stock_booth
- **Transactions** - transactions, transaction_details
- **Orders** - customer_orders, customer_order_details
- **Supporting** - temp_cart, pegawai, version_admin, version_shop

For detailed schema documentation, see [DATABASE_SCHEMA.md](DATABASE_SCHEMA.md)

## 📁 Project Structure

```
e-canteen-cashier-api/
├── app/
│   ├── repository/          # Data access layer
│   └── usecase/             # Business logic layer
├── config/                  # Configuration files
├── database/
│   ├── migrations/          # Database migration files
│   ├── seeders/            # Database seeder files
│   ├── migrate.ps1         # Migration script (Windows)
│   ├── migrate.sh          # Migration script (Linux/Mac)
│   ├── seed.ps1            # Seeder script (Windows)
│   └── seed.sh             # Seeder script (Linux/Mac)
├── docs/                    # Swagger documentation
├── entity/                  # Data models
├── handler/                 # Response handlers
├── middleware/              # HTTP middlewares
├── pkg/                     # Utilities and helpers
│   ├── exceptions/         # Custom exceptions
│   ├── helpers/            # Helper functions
│   ├── mysql/              # Database connection
│   └── validations/        # Validation utilities
├── repository/              # Repository implementations
├── router/                  # Route definitions
├── .env                     # Environment variables
├── main.go                  # Application entry point
├── Makefile                 # Build automation
└── README.md               # This file
```

## 🛠️ Development

### Generate Swagger Documentation

```bash
# Install swag CLI
go install github.com/swaggo/swag/cmd/swag@latest

# Generate docs
swag init -g main.go --parseDependency --parseInternal

# Or using Makefile
make swagger
```

### Run Tests

```bash
go test -v ./...

# Or using Makefile
make test
```

### Build for Production

```bash
go build -o bin/ecanteen-api main.go

# Or using Makefile
make build
```

## 🔄 Database Management

### Run Migrations

```bash
# Windows
cd database
.\migrate.ps1

# Linux/Mac
cd database
./migrate.sh

# Using Makefile
make migrate        # Linux/Mac
make migrate-win    # Windows
```

### Run Seeders

```bash
# Windows
cd database
.\seed.ps1

# Linux/Mac
cd database
./seed.sh

# Using Makefile
make seed          # Linux/Mac
make seed-win      # Windows
```

### Reset Database

```bash
# Drop and recreate database
mysql -u root -p -e "DROP DATABASE IF EXISTS e-canteen_new; CREATE DATABASE e-canteen_new CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

# Run migrations and seeders
make db-setup      # Linux/Mac
make db-setup-win  # Windows
```

## 📝 Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `HOST_ADDR` | API server address | `127.0.0.1:3000` |
| `DB_DRIVER` | Database driver | `mysql` |
| `DB_NAME` | Database name | `e-canteen_new` |
| `DB_USERNAME` | Database username | `root` |
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `3306` |
| `DB_PASSWORD` | Database password | - |
| `SECRET_KEY` | JWT secret key | - |
| `BASE_ASSETS` | Base URL for assets | - |
| `FCM_SERVER_KEY` | Firebase Cloud Messaging key | - |

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License.

## 👥 Authors

- **jolebo** - [GitHub](https://github.com/jolebo)

## 📞 Support

For support, email support@ecanteen.com or create an issue in the repository.

---

Made with ❤️ using Go
