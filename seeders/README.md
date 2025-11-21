# Database Seeders

This directory contains database seeders for the e-canteen application.

## Permission Seeder

Seeds the database with default permissions and assigns all permissions to the Super Admin role (role_id = 1).

### Permissions Created

The seeder creates permissions for the following resources with CRUD operations:

- **Product**: create, read, update, delete
- **Category**: create, read, update, delete
- **Order**: create, read, update, delete, process, finish, cancel
- **Customer**: create, read, update, delete
- **User**: create, read, update, delete
- **Transaction**: create, read, update, delete
- **Varian**: create, read, update, delete
- **Major**: create, read, update, delete
- **Territory**: create, read, update, delete
- **Temp Cart**: create, read, update, delete
- **Customer Address**: create, read, update, delete
- **Pegawai**: create, read, update, delete
- **Stock**: create, read, update, delete
- **Permission**: create, read, update, delete, assign, revoke
- **Report**: order, transaction

Total: **62 permissions**

### Running the Seeder

#### Prerequisites

1. Ensure the database migrations have been run:
   ```bash
   goose -dir migrations mysql "user:password@tcp(localhost:3306)/e-canteen_new" up
   ```

2. Make sure your `.env` file is properly configured with database credentials.

#### Run the Seeder

From the project root directory:

```bash
# Run the seeder
go run seeders/main.go
```

Or build and run:

```bash
# Build the seeder
go build -o bin/seeder seeders/main.go

# Run it
./bin/seeder
```

### What the Seeder Does

1. Creates all 62 permissions in the `permissions` table
2. Assigns all permissions to Super Admin (role_id = 1) in the `permission_role` table
3. Skips permissions that already exist (based on permission name)
4. Prevents duplicate assignments

### Verifying the Seeder

After running the seeder, you can verify the results:

```sql
-- Check total permissions created
SELECT COUNT(*) FROM permissions;

-- View all permissions
SELECT * FROM permissions ORDER BY permission_resource, permission_action;

-- Check Super Admin's permissions
SELECT p.* 
FROM permissions p
INNER JOIN permission_role pr ON p.permission_id = pr.permission_id
WHERE pr.role_id = 1
ORDER BY p.permission_resource, p.permission_action;

-- Count Super Admin's permissions
SELECT COUNT(*) 
FROM permission_role 
WHERE role_id = 1;
```

### Notes

- The seeder is idempotent - you can run it multiple times safely
- It will skip existing permissions and only create new ones
- Super Admin (role_id = 1) must exist in your roles table before running the seeder
- The seeder uses transactions for data integrity
