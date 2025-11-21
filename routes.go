package main

// This file contains Swagger route annotations
// These are parsed by swag to generate API documentation

// Authentication Routes

// @Summary User login (Kasir)
// @Description Authenticate kasir user and get JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body entity.Login true "Login credentials"
// @Success 200 {object} handler.WebResponse{data=string} "Login successful with JWT token"
// @Failure 401 {object} handler.WebResponse "Invalid credentials"
// @Router /kasir/login [post]
func kasirLogin() {}

// @Summary User logout (Kasir)
// @Description Logout kasir user and invalidate token
// @Tags Authentication
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} handler.WebResponse "Logout successful"
// @Router /kasir/logout [put]
func kasirLogout() {}

// Customer Routes

// @Summary Customer login
// @Description Authenticate customer and get JWT token
// @Tags Customer
// @Accept json
// @Produce json
// @Param request body entity.Login true "Login credentials"
// @Success 200 {object} handler.WebResponse{data=string} "Login successful with JWT token"
// @Failure 401 {object} handler.WebResponse "Invalid credentials"
// @Router /customer/login [post]
func customerLogin() {}

// @Summary Customer logout
// @Description Logout customer and invalidate token
// @Tags Customer
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} handler.WebResponse "Logout successful"
// @Router /customer/logout [post]
func customerLogout() {}

// @Summary Update FCM token
// @Description Update Firebase Cloud Messaging token for push notifications
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body entity.UpdateFcmRequest true "FCM token data"
// @Success 200 {object} handler.WebResponse "FCM token updated successfully"
// @Failure 400 {object} handler.WebResponse "Invalid request"
// @Failure 404 {object} handler.WebResponse "User not found"
// @Router /user/fcm [put]
func updateFcm() {}

// @Summary Get all customers
// @Description Retrieve all customers
// @Tags Customer
// @Produce json
// @Security BearerAuth
// @Success 200 {object} handler.WebResponse{data=[]entity.Customer} "Customers list"
// @Router /customer [get]
func getCustomers() {}

// @Summary Get customer by ID
// @Description Retrieve customer details by ID
// @Tags Customer
// @Produce json
// @Security BearerAuth
// @Param customerId path string true "Customer ID"
// @Success 200 {object} handler.WebResponse{data=entity.Customer} "Customer details"
// @Failure 404 {object} handler.WebResponse "Customer not found"
// @Router /customer/{customerId} [get]
func getCustomerById() {}

// @Summary Register customer
// @Description Register a new customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param request body entity.Customer true "Customer data"
// @Success 201 {object} handler.WebResponse{data=entity.Customer} "Customer registered"
// @Failure 400 {object} handler.WebResponse "Invalid request"
// @Router /customer [post]
func registerCustomer() {}

// @Summary Update customer
// @Description Update customer information
// @Tags Customer
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param customerId path string true "Customer ID"
// @Param request body entity.Customer true "Customer data"
// @Success 200 {object} handler.WebResponse{data=entity.Customer} "Customer updated"
// @Failure 404 {object} handler.WebResponse "Customer not found"
// @Router /customer/{customerId} [put]
func updateCustomer() {}

// @Summary Delete customer
// @Description Delete customer by ID
// @Tags Customer
// @Produce json
// @Security BearerAuth
// @Param customerId path string true "Customer ID"
// @Success 200 {object} handler.WebResponse "Customer deleted"
// @Failure 404 {object} handler.WebResponse "Customer not found"
// @Router /customer/{customerId} [delete]
func deleteCustomer() {}

// Product Routes

// @Summary Get all products
// @Description Retrieve all products with their variants
// @Tags Products
// @Produce json
// @Security BearerAuth
// @Success 200 {object} handler.WebResponse{data=[]entity.Product} "Products list"
// @Router /products [get]
func getProducts() {}

// Transaction Routes

// @Summary Create transaction
// @Description Create a new cashier transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body entity.Transaction true "Transaction data"
// @Success 201 {object} handler.WebResponse{data=entity.Transaction} "Transaction created"
// @Failure 400 {object} handler.WebResponse "Invalid request"
// @Router /kasir/transaction [post]
func createTransaction() {}

// @Summary Get all transactions
// @Description Retrieve all transactions
// @Tags Transactions
// @Produce json
// @Security BearerAuth
// @Success 200 {object} handler.WebResponse{data=[]entity.Transaction} "Transactions list"
// @Router /transaction [get]
func getTransactions() {}

// @Summary Get transaction by ID
// @Description Retrieve transaction details by ID
// @Tags Transactions
// @Produce json
// @Security BearerAuth
// @Param transId path string true "Transaction ID"
// @Success 200 {object} handler.WebResponse{data=entity.Transaction} "Transaction details"
// @Failure 404 {object} handler.WebResponse "Transaction not found"
// @Router /kasir/transaction/{transId} [get]
func getTransactionById() {}

// @Summary Get transaction details
// @Description Retrieve transaction detail items
// @Tags Transactions
// @Produce json
// @Security BearerAuth
// @Param trans_id query string true "Transaction ID"
// @Success 200 {object} handler.WebResponse{data=[]entity.TransactionDetail} "Transaction details"
// @Router /kasir/transaction_detail [get]
func getTransactionDetails() {}

// Cart Routes

// @Summary Get cart items by user ID
// @Description Get all temporary cart items for a specific user
// @Tags Cart
// @Produce json
// @Security BearerAuth
// @Param userId path string true "User ID"
// @Success 200 {object} handler.WebResponse{data=[]entity.TempCart} "Cart items"
// @Failure 400 {object} handler.WebResponse "Invalid request"
// @Router /tempcart/{userId} [get]
func getCartByUserId() {}

// @Summary Add item to cart
// @Description Add product variant to temporary cart
// @Tags Cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body entity.TempCart true "Cart item"
// @Success 201 {object} handler.WebResponse{data=entity.TempCart} "Item added to cart"
// @Failure 400 {object} handler.WebResponse "Invalid request"
// @Router /tempcart [post]
func addToCart() {}

// @Summary Update cart item quantity
// @Description Update quantity of item in temporary cart
// @Tags Cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param productVarianId path string true "Product Variant ID"
// @Param userId path string true "User ID"
// @Success 200 {object} handler.WebResponse "Cart item updated"
// @Failure 404 {object} handler.WebResponse "Cart item not found"
// @Router /tempcart/{productVarianId}/{userId} [put]
func updateCartItem() {}

// @Summary Remove item from cart
// @Description Remove product variant from temporary cart
// @Tags Cart
// @Produce json
// @Security BearerAuth
// @Param productVarianId path string true "Product Variant ID"
// @Param userId path string true "User ID"
// @Success 200 {object} handler.WebResponse "Item removed from cart"
// @Failure 404 {object} handler.WebResponse "Cart item not found"
// @Router /tempcart/{productVarianId}/{userId} [delete]
func removeFromCart() {}

// Order Routes

// @Summary Create a new customer order
// @Description Create a new order with order details
// @Tags Order
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body entity.CreateOrderRequest true "Order request payload"
// @Success 200 {object} handler.WebResponse{data=entity.CustomerOrder} "Order created successfully"
// @Failure 400 {object} handler.WebResponse "Invalid request or insufficient stock"
// @Router /order [post]
func createOrder() {}

// @Summary Get order report by date range
// @Description Get customer orders filtered by date range (start_date and/or end_date)
// @Tags Order
// @Produce json
// @Security BearerAuth
// @Param start_date query string false "Start date (YYYY-MM-DD format)"
// @Param end_date query string false "End date (YYYY-MM-DD format)"
// @Success 200 {object} handler.WebResponse{data=[]entity.CustomerOrder} "Order report"
// @Failure 400 {object} handler.WebResponse "Invalid request - start_date or end_date required"
// @Router /order_report [get]
func getOrderReport() {}

// Version Routes

// @Summary Get admin app version
// @Description Get current admin application version
// @Tags Version
// @Produce json
// @Success 200 {object} handler.WebResponse{data=entity.VersionAdmin} "Version info"
// @Router /kasir/version [get]
func getAdminVersion() {}

// @Summary Get shop app version
// @Description Get current customer shop application version
// @Tags Version
// @Produce json
// @Success 200 {object} handler.WebResponse{data=entity.VersionShop} "Version info"
// @Router /shop/version [get]
func getShopVersion() {}

// Dashboard Routes

// @Summary Get dashboard statistics
// @Description Get comprehensive dashboard statistics including total sales, transactions, customers, monthly sales chart, and top 10 products
// @Tags Dashboard
// @Produce json
// @Security BearerAuth
// @Param year query int false "Year for monthly sales data (default: current year)"
// @Success 200 {object} handler.WebResponse{data=entity.DashboardStats} "Dashboard statistics"
// @Failure 500 {object} handler.WebResponse "Internal server error"
// @Router /dashboard/stats [get]
func getDashboardStats() {}

// Permission Routes (RBAC)

// @Summary Create permission
// @Description Create a new permission for RBAC
// @Tags Permission
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body entity.CreatePermissionRequest true "Permission data"
// @Success 200 {object} handler.WebResponse{data=entity.Permission} "Permission created"
// @Failure 400 {object} handler.WebResponse "Invalid request"
// @Router /permission [post]
func createPermission() {}

// @Summary Get all permissions
// @Description Retrieve all permissions
// @Tags Permission
// @Produce json
// @Security BearerAuth
// @Success 200 {object} handler.WebResponse{data=[]entity.Permission} "Permissions list"
// @Router /permission [get]
func getAllPermissions() {}

// @Summary Get permission by ID
// @Description Retrieve permission details by ID
// @Tags Permission
// @Produce json
// @Security BearerAuth
// @Param permissionId path int true "Permission ID"
// @Success 200 {object} handler.WebResponse{data=entity.Permission} "Permission details"
// @Failure 404 {object} handler.WebResponse "Permission not found"
// @Router /permission/{permissionId} [get]
func getPermissionById() {}

// @Summary Update permission
// @Description Update an existing permission
// @Tags Permission
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param permissionId path int true "Permission ID"
// @Param request body entity.UpdatePermissionRequest true "Updated permission data"
// @Success 200 {object} handler.WebResponse{data=entity.Permission} "Permission updated"
// @Failure 400 {object} handler.WebResponse "Invalid request"
// @Failure 404 {object} handler.WebResponse "Permission not found"
// @Router /permission/{permissionId} [put]
func updatePermission() {}

// @Summary Delete permission
// @Description Delete a permission
// @Tags Permission
// @Produce json
// @Security BearerAuth
// @Param permissionId path int true "Permission ID"
// @Success 200 {object} handler.WebResponse "Permission deleted"
// @Failure 404 {object} handler.WebResponse "Permission not found"
// @Router /permission/{permissionId} [delete]
func deletePermission() {}

// @Summary Get permissions by role
// @Description Retrieve all permissions assigned to a role
// @Tags Permission
// @Produce json
// @Security BearerAuth
// @Param roleId path int true "Role ID"
// @Success 200 {object} handler.WebResponse{data=[]entity.Permission} "Permissions for role"
// @Router /permission/role/{roleId} [get]
func getPermissionsByRole() {}

// @Summary Assign permissions to role
// @Description Assign multiple permissions to a role
// @Tags Permission
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body entity.AssignPermissionRequest true "Role and permission IDs"
// @Success 200 {object} handler.WebResponse "Permissions assigned"
// @Failure 400 {object} handler.WebResponse "Invalid request"
// @Router /permission/assign [post]
func assignPermissionsToRole() {}

// @Summary Revoke permission from role
// @Description Revoke a permission from a role
// @Tags Permission
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body entity.RevokePermissionRequest true "Role and permission ID"
// @Success 200 {object} handler.WebResponse "Permission revoked"
// @Failure 400 {object} handler.WebResponse "Invalid request"
// @Router /permission/revoke [delete]
func revokePermissionFromRole() {}
