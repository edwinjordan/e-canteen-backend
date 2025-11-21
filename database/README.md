# Database Setup Guide

## Prerequisites

- MySQL Server 5.7+ or MariaDB 10.2+
- MySQL command-line client installed and available in PATH

## Configuration

Set the following environment variables before running the scripts:

```bash
# Linux/Mac
export DB_HOST=localhost
export DB_PORT=3306
export DB_USER=root
export DB_PASS=your_password
export DB_NAME=ecanteen_db

# Windows PowerShell
$env:DB_HOST="localhost"
$env:DB_PORT="3306"
$env:DB_USER="root"
$env:DB_PASS="your_password"
$env:DB_NAME="ecanteen_db"
```

Or modify the default values in the migration/seeder scripts.

## Running Migrations

### Linux/Mac

```bash
cd database
chmod +x migrate.sh
./migrate.sh
```

### Windows PowerShell

```powershell
cd database
.\migrate.ps1
```

## Running Seeders

### Linux/Mac

```bash
cd database
chmod +x seed.sh
./seed.sh
```

### Windows PowerShell

```powershell
cd database
.\seed.ps1
```

## Default Credentials

### Admin Users
- **Email**: admin@ecanteen.com
- **Password**: admin123
- **Role**: Super Admin

- **Email**: siti@ecanteen.com
- **Password**: admin123
- **Role**: Admin

### Cashier Users
- **Email**: budi@ecanteen.com
- **Password**: admin123
- **Role**: Cashier

- **Email**: dewi@ecanteen.com
- **Password**: admin123
- **Role**: Cashier

### Customer Accounts
All customers have default password: **customer123**

Sample customers:
- andi@student.com
- rina@student.com
- joko@student.com
- lilis@student.com
- agus@student.com

## Migration Files

The migration files are located in `database/migrations/` and are executed in numerical order:

1. `001_create_roles_table.sql` - Create roles table
2. `002_create_pegawai_table.sql` - Create employee table
3. `003_create_users_table.sql` - Create users table
4. `004_create_user_logs_table.sql` - Create user logs table
5. `005_create_majors_table.sql` - Create majors table
6. `006_create_customers_table.sql` - Create customers table
7. `007_create_user_otp_table.sql` - Create OTP table
8. `008_create_customer_address_table.sql` - Create customer address table
9. `009_create_categories_table.sql` - Create categories table
10. `010_create_products_table.sql` - Create products table
11. `011_create_product_varians_table.sql` - Create product variants table
12. `012_create_stock_booth_table.sql` - Create stock booth table
13. `013_create_customer_orders_table.sql` - Create customer orders table
14. `014_create_customer_order_details_table.sql` - Create order details table
15. `015_create_transactions_table.sql` - Create transactions table
16. `016_create_transaction_details_table.sql` - Create transaction details table
17. `017_create_temp_cart_table.sql` - Create temporary cart table
18. `018_create_version_admin_table.sql` - Create admin version table
19. `019_create_version_shop_table.sql` - Create shop version table

## Seeder Files

The seeder files are located in `database/seeders/` and populate initial data:

1. `001_seed_roles.sql` - Seed roles (Super Admin, Admin, Cashier, etc.)
2. `002_seed_majors.sql` - Seed majors/departments
3. `003_seed_categories.sql` - Seed product categories
4. `004_seed_pegawai.sql` - Seed sample employees
5. `005_seed_users.sql` - Seed sample users
6. `006_seed_customers.sql` - Seed sample customers
7. `007_seed_products.sql` - Seed sample products
8. `008_seed_product_varians.sql` - Seed product variants
9. `009_seed_version_admin.sql` - Seed admin app versions
10. `010_seed_version_shop.sql` - Seed shop app versions

## Troubleshooting

### MySQL Command Not Found

Make sure MySQL client is installed and added to your system PATH.

**Windows:**
- Add MySQL bin directory to PATH (e.g., `C:\Program Files\MySQL\MySQL Server 8.0\bin`)

**Linux/Mac:**
```bash
# Install MySQL client
sudo apt-get install mysql-client  # Ubuntu/Debian
brew install mysql-client          # Mac with Homebrew
```

### Permission Denied (Linux/Mac)

```bash
chmod +x migrate.sh seed.sh
```

### Connection Refused

- Ensure MySQL server is running
- Check host and port configuration
- Verify firewall settings

### Foreign Key Constraint Errors

Run migrations in the correct order. The migration script handles this automatically.

## Manual Execution

If you prefer to run SQL files manually:

```bash
# Run all migrations
for file in database/migrations/*.sql; do
    mysql -u root -p ecanteen_db < "$file"
done

# Run all seeders
for file in database/seeders/*.sql; do
    mysql -u root -p ecanteen_db < "$file"
done
```

## Reset Database

To completely reset the database:

```bash
mysql -u root -p -e "DROP DATABASE IF EXISTS ecanteen_db; CREATE DATABASE ecanteen_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
```

Then run migrations and seeders again.
