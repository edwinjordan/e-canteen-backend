package main

import (
	"context"
	"fmt"
	"log"

	"github.com/edwinjordan/e-canteen-backend/entity"
	"github.com/edwinjordan/e-canteen-backend/repository/permission_repository"
	"gorm.io/gorm"
)

type PermissionSeeder struct {
	DB *gorm.DB
}

func NewPermissionSeeder(db *gorm.DB) *PermissionSeeder {
	return &PermissionSeeder{DB: db}
}

func (s *PermissionSeeder) SeedPermissions() error {
	ctx := context.Background()
	permissionRepo := permission_repository.New(s.DB)
	permissionRoleRepo := permission_repository.NewPermissionRole(s.DB)

	// Map to store main menu and submenu IDs
	menuIds := make(map[string]int)
	submenuIds := make(map[string]int)

	// Define all permissions for the system
	// Organized by menu structure: Dashboard, Master Data, Kasir, Laporan
	permissions := []struct {
		Name        string
		Resource    string
		Action      string
		Description string
		Status      string // main_menu, submenu, or action
		ParentKey   string // key to find parent ID from map
	}{
		// ========== MAIN MENUS ==========
		{"Dashboard", "dashboard", "menu", "Dashboard main menu", "main_menu", ""},
		{"Master Data", "master_data", "menu", "Master Data main menu", "main_menu", ""},
		{"Kasir", "kasir", "menu", "Kasir main menu", "main_menu", ""},
		{"Laporan", "laporan", "menu", "Laporan main menu", "main_menu", ""},

		// ========== DASHBOARD ==========
		{"View Dashboard", "dashboard", "view", "Permission to view dashboard", "action", "dashboard"},

		// ========== MASTER DATA MENU ==========
		// Category submenu
		{"Category", "category", "submenu", "Category submenu", "submenu", "master_data"},
		{"Create Category", "category", "create", "Permission to create new categories", "action", "category"},
		{"Read Category", "category", "read", "Permission to view categories", "action", "category"},
		{"Update Category", "category", "update", "Permission to update categories", "action", "category"},
		{"Delete Category", "category", "delete", "Permission to delete categories", "action", "category"},

		// Customers submenu
		{"Customers", "customer", "submenu", "Customers submenu", "submenu", "master_data"},
		{"Create Customer", "customer", "create", "Permission to create new customers", "action", "customer"},
		{"Read Customer", "customer", "read", "Permission to view customers", "action", "customer"},
		{"Update Customer", "customer", "update", "Permission to update customers", "action", "customer"},
		{"Delete Customer", "customer", "delete", "Permission to delete customers", "action", "customer"},

		// Customer Address (part of customers submenu)
		{"Create Customer Address", "customer_address", "create", "Permission to create customer addresses", "action", "customer"},
		{"Read Customer Address", "customer_address", "read", "Permission to view customer addresses", "action", "customer"},
		{"Update Customer Address", "customer_address", "update", "Permission to update customer addresses", "action", "customer"},
		{"Delete Customer Address", "customer_address", "delete", "Permission to delete customer addresses", "action", "customer"},

		// Majors submenu
		{"Majors", "major", "submenu", "Majors submenu", "submenu", "master_data"},
		{"Create Major", "major", "create", "Permission to create new majors", "action", "major"},
		{"Read Major", "major", "read", "Permission to view majors", "action", "major"},
		{"Update Major", "major", "update", "Permission to update majors", "action", "major"},
		{"Delete Major", "major", "delete", "Permission to delete majors", "action", "major"},

		// Produk submenu
		{"Produk", "product", "submenu", "Produk submenu", "submenu", "master_data"},
		{"Create Product", "product", "create", "Permission to create new products", "action", "product"},
		{"Read Product", "product", "read", "Permission to view products", "action", "product"},
		{"Update Product", "product", "update", "Permission to update products", "action", "product"},
		{"Delete Product", "product", "delete", "Permission to delete products", "action", "product"},

		// Varian (part of product submenu)
		{"Create Varian", "varian", "create", "Permission to create new variants", "action", "product"},
		{"Read Varian", "varian", "read", "Permission to view variants", "action", "product"},
		{"Update Varian", "varian", "update", "Permission to update variants", "action", "product"},
		{"Delete Varian", "varian", "delete", "Permission to delete variants", "action", "product"},

		// Territory submenu
		{"Territory", "territory", "submenu", "Territory submenu", "submenu", "master_data"},
		{"Create Territory", "territory", "create", "Permission to create new territories", "action", "territory"},
		{"Read Territory", "territory", "read", "Permission to view territories", "action", "territory"},
		{"Update Territory", "territory", "update", "Permission to update territories", "action", "territory"},
		{"Delete Territory", "territory", "delete", "Permission to delete territories", "action", "territory"},

		// Stock submenu
		{"Stock", "stock", "submenu", "Stock submenu", "submenu", "master_data"},
		{"Create Stock", "stock", "create", "Permission to create stock entries", "action", "stock"},
		{"Read Stock", "stock", "read", "Permission to view stock", "action", "stock"},
		{"Update Stock", "stock", "update", "Permission to update stock", "action", "stock"},
		{"Delete Stock", "stock", "delete", "Permission to delete stock entries", "action", "stock"},

		// ========== KASIR MENU ==========
		// POS submenu
		{"POS", "pos", "submenu", "POS submenu", "submenu", "kasir"},
		{"Access POS", "pos", "access", "Permission to access POS system", "action", "pos"},
		{"Create Transaction", "transaction", "create", "Permission to create new transactions", "action", "pos"},
		{"Read Transaction", "transaction", "read", "Permission to view transactions", "action", "pos"},
		{"Update Transaction", "transaction", "update", "Permission to update transactions", "action", "pos"},
		{"Delete Transaction", "transaction", "delete", "Permission to delete transactions", "action", "pos"},

		// Temp Cart (for POS)
		{"Create Temp Cart", "tempcart", "create", "Permission to create temp cart items", "action", "pos"},
		{"Read Temp Cart", "tempcart", "read", "Permission to view temp cart", "action", "pos"},
		{"Update Temp Cart", "tempcart", "update", "Permission to update temp cart", "action", "pos"},
		{"Delete Temp Cart", "tempcart", "delete", "Permission to delete temp cart items", "action", "pos"},

		// Order Management submenu
		{"Order Management", "order", "submenu", "Order Management submenu", "submenu", "kasir"},
		{"Create Order", "order", "create", "Permission to create new orders", "action", "order"},
		{"Read Order", "order", "read", "Permission to view orders", "action", "order"},
		{"Update Order", "order", "update", "Permission to update orders", "action", "order"},
		{"Delete Order", "order", "delete", "Permission to delete orders", "action", "order"},
		{"Process Order", "order", "process", "Permission to process orders", "action", "order"},
		{"Finish Order", "order", "finish", "Permission to finish orders", "action", "order"},
		{"Cancel Order", "order", "cancel", "Permission to cancel orders", "action", "order"},

		// ========== LAPORAN MENU ==========
		// Laporan Penjualan submenu
		{"Laporan Penjualan", "report", "submenu", "Laporan Penjualan submenu", "submenu", "laporan"},
		{"View Sales Report", "report", "sales", "Permission to view sales reports", "action", "report"},
		{"View Order Report", "report", "order", "Permission to view order reports", "action", "report"},
		{"View Transaction Report", "report", "transaction", "Permission to view transaction reports", "action", "report"},
		{"Export Sales Report", "report", "export", "Permission to export sales reports", "action", "report"},

		// ========== SYSTEM / ADMIN ==========
		// User Management submenu (no main menu, standalone in sidebar)
		{"User Management", "user", "submenu", "User Management submenu", "submenu", ""},
		{"Create User", "user", "create", "Permission to create new users", "action", "user"},
		{"Read User", "user", "read", "Permission to view users", "action", "user"},
		{"Update User", "user", "update", "Permission to update users", "action", "user"},
		{"Delete User", "user", "delete", "Permission to delete users", "action", "user"},

		// Employee Management submenu
		{"Employee Management", "pegawai", "submenu", "Employee Management submenu", "submenu", ""},
		{"Create Pegawai", "pegawai", "create", "Permission to create new employees", "action", "pegawai"},
		{"Read Pegawai", "pegawai", "read", "Permission to view employees", "action", "pegawai"},
		{"Update Pegawai", "pegawai", "update", "Permission to update employees", "action", "pegawai"},
		{"Delete Pegawai", "pegawai", "delete", "Permission to delete employees", "action", "pegawai"},

		// Permission Management submenu
		{"Permission Management", "permission", "submenu", "Permission Management submenu", "submenu", ""},
		{"Create Permission", "permission", "create", "Permission to create new permissions", "action", "permission"},
		{"Read Permission", "permission", "read", "Permission to view permissions", "action", "permission"},
		{"Update Permission", "permission", "update", "Permission to update permissions", "action", "permission"},
		{"Delete Permission", "permission", "delete", "Permission to delete permissions", "action", "permission"},
		{"Assign Permission", "permission", "assign", "Permission to assign permissions to roles", "action", "permission"},
		{"Revoke Permission", "permission", "revoke", "Permission to revoke permissions from roles", "action", "permission"},
	}

	// Create permissions and collect their IDs
	var permissionIds []int
	log.Println("Starting permission seeding...")

	for _, perm := range permissions {
		// Determine parent ID
		var parentId *int
		if perm.ParentKey != "" {
			if id, exists := menuIds[perm.ParentKey]; exists {
				parentId = &id
			} else if id, exists := submenuIds[perm.ParentKey]; exists {
				parentId = &id
			}
		}

		model := entity.Permission{
			PermissionName:        perm.Name,
			PermissionResource:    perm.Resource,
			PermissionAction:      perm.Action,
			PermissionDescription: perm.Description,
			PermissionStatus:      perm.Status,
			PermissionParentId:    parentId,
		}

		// Check if permission already exists
		existing := permissionRepo.FindByName(ctx, perm.Name)
		if existing.PermissionId != 0 {
			log.Printf("Permission '%s' already exists, skipping...\n", perm.Name)
			permissionIds = append(permissionIds, existing.PermissionId)

			// Store ID for reference
			if perm.Status == "main_menu" {
				menuIds[perm.Resource] = existing.PermissionId
			} else if perm.Status == "submenu" {
				submenuIds[perm.Resource] = existing.PermissionId
			}
			continue
		}

		// Create new permission
		created := permissionRepo.Create(ctx, model)
		permissionIds = append(permissionIds, created.PermissionId)
		log.Printf("Created permission: %s (ID: %d, Parent: %v)\n", perm.Name, created.PermissionId, parentId)

		// Store ID for reference
		if perm.Status == "main_menu" {
			menuIds[perm.Resource] = created.PermissionId
		} else if perm.Status == "submenu" {
			submenuIds[perm.Resource] = created.PermissionId
		}
	}

	log.Printf("Total permissions created/found: %d\n", len(permissionIds))

	// Assign all permissions to Super Admin (role_id = 1)
	superAdminRoleId := 1
	log.Printf("Assigning all permissions to Super Admin (role_id: %d)...\n", superAdminRoleId)

	err := permissionRoleRepo.Assign(ctx, superAdminRoleId, permissionIds)
	if err != nil {
		return fmt.Errorf("failed to assign permissions to Super Admin: %v", err)
	}

	log.Printf("Successfully assigned %d permissions to Super Admin\n", len(permissionIds))
	return nil
}
