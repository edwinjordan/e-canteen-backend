package docs

// Swagger annotations for API endpoints

// Login godoc
// @Summary User login
// @Description Authenticate user and get JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} WebResponse{data=LoginResponse} "Login successful"
// @Failure 400 {object} WebResponse "Invalid request"
// @Failure 401 {object} WebResponse "Invalid credentials"
// @Router /kasir/login [post]
func LoginEndpoint() {}

// Logout godoc
// @Summary User logout
// @Description Logout user and invalidate token
// @Tags Authentication
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body LogoutRequest true "Logout request with token"
// @Success 200 {object} WebResponse "Logout successful"
// @Failure 400 {object} WebResponse "Invalid request"
// @Router /kasir/logout [put]
func LogoutEndpoint() {}

// GetVersionAdmin godoc
// @Summary Get admin app version
// @Description Get current admin application version information
// @Tags Version
// @Produce json
// @Success 200 {object} WebResponse{data=VersionResponse} "Version info retrieved"
// @Router /kasir/version [get]
func GetVersionAdminEndpoint() {}

// GetVersionShop godoc
// @Summary Get shop app version
// @Description Get current customer shop application version information
// @Tags Version
// @Produce json
// @Success 200 {object} WebResponse{data=VersionResponse} "Version info retrieved"
// @Router /shop/version [get]
func GetVersionShopEndpoint() {}

// CheckMaintenanceMode godoc
// @Summary Check maintenance mode
// @Description Check if system is in maintenance mode
// @Tags Version
// @Produce json
// @Param confCode path string true "Configuration code"
// @Success 200 {object} WebResponse "Maintenance mode status"
// @Router /check_maintenance_mode/{confCode} [get]
func CheckMaintenanceModeEndpoint() {}

// CreateTransaction godoc
// @Summary Create transaction
// @Description Create a new transaction/sale
// @Tags Transactions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateTransactionRequest true "Transaction data"
// @Success 201 {object} WebResponse{data=Transaction} "Transaction created"
// @Failure 400 {object} WebResponse "Invalid request"
// @Failure 401 {object} WebResponse "Unauthorized"
// @Router /kasir/transaction [post]
func CreateTransactionEndpoint() {}

// GetAllTransactions godoc
// @Summary Get all transactions
// @Description Retrieve all transactions with pagination
// @Tags Transactions
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Param search query string false "Search by invoice or customer name"
// @Success 200 {object} WebResponse{data=[]Transaction} "Transactions retrieved"
// @Failure 401 {object} WebResponse "Unauthorized"
// @Router /transaction [get]
func GetAllTransactionsEndpoint() {}

// GetTransactionById godoc
// @Summary Get transaction by ID
// @Description Retrieve transaction details by transaction ID
// @Tags Transactions
// @Produce json
// @Security BearerAuth
// @Param transId path string true "Transaction ID"
// @Success 200 {object} WebResponse{data=Transaction} "Transaction details"
// @Failure 404 {object} WebResponse "Transaction not found"
// @Router /kasir/transaction/{transId} [get]
func GetTransactionByIdEndpoint() {}

// GetTransactionDetail godoc
// @Summary Get transaction details
// @Description Retrieve transaction detail items
// @Tags Transactions
// @Produce json
// @Security BearerAuth
// @Param trans_id query string true "Transaction ID"
// @Success 200 {object} WebResponse{data=[]TransactionDetail} "Transaction details"
// @Router /kasir/transaction_detail [get]
func GetTransactionDetailEndpoint() {}

// GetTransactionSummary godoc
// @Summary Get transaction summary
// @Description Retrieve transaction summary and statistics
// @Tags Transactions
// @Produce json
// @Security BearerAuth
// @Param date_from query string false "Start date (YYYY-MM-DD)"
// @Param date_to query string false "End date (YYYY-MM-DD)"
// @Success 200 {object} WebResponse "Transaction summary"
// @Router /kasir/transaction_summary [get]
func GetTransactionSummaryEndpoint() {}

// GetAllProducts godoc
// @Summary Get all products
// @Description Retrieve all products with their variants
// @Tags Products
// @Produce json
// @Security BearerAuth
// @Param category_id query string false "Filter by category ID"
// @Param search query string false "Search by product name"
// @Success 200 {object} WebResponse{data=[]Product} "Products retrieved"
// @Router /products [get]
func GetAllProductsEndpoint() {}

// GetAllCategories godoc
// @Summary Get all categories
// @Description Retrieve all product categories
// @Tags Categories
// @Produce json
// @Security BearerAuth
// @Success 200 {object} WebResponse{data=[]Category} "Categories retrieved"
// @Router /categories [get]
func GetAllCategoriesEndpoint() {}

// GetAllVariants godoc
// @Summary Get all variants
// @Description Retrieve all product variants
// @Tags Variants
// @Produce json
// @Security BearerAuth
// @Param product_id query string false "Filter by product ID"
// @Success 200 {object} WebResponse{data=[]ProductVariant} "Variants retrieved"
// @Router /variants [get]
func GetAllVariantsEndpoint() {}

// GetAllCustomers godoc
// @Summary Get all customers
// @Description Retrieve all customers
// @Tags Customers
// @Produce json
// @Security BearerAuth
// @Param search query string false "Search by name or email"
// @Param major_id query string false "Filter by major ID"
// @Success 200 {object} WebResponse{data=[]Customer} "Customers retrieved"
// @Router /customers [get]
func GetAllCustomersEndpoint() {}

// GetCustomerById godoc
// @Summary Get customer by ID
// @Description Retrieve customer details by customer ID
// @Tags Customers
// @Produce json
// @Security BearerAuth
// @Param customerId path string true "Customer ID"
// @Success 200 {object} WebResponse{data=Customer} "Customer details"
// @Failure 404 {object} WebResponse "Customer not found"
// @Router /customers/{customerId} [get]
func GetCustomerByIdEndpoint() {}

// GetCustomerAddresses godoc
// @Summary Get customer addresses
// @Description Retrieve all addresses for a customer
// @Tags Addresses
// @Produce json
// @Security BearerAuth
// @Param customer_id query string true "Customer ID"
// @Success 200 {object} WebResponse{data=[]CustomerAddress} "Addresses retrieved"
// @Router /customer_address [get]
func GetCustomerAddressesEndpoint() {}

// GetAllMajors godoc
// @Summary Get all majors
// @Description Retrieve all majors/departments
// @Tags Majors
// @Produce json
// @Security BearerAuth
// @Success 200 {object} WebResponse{data=[]Major} "Majors retrieved"
// @Router /majors [get]
func GetAllMajorsEndpoint() {}

// GetAllOrders godoc
// @Summary Get all orders
// @Description Retrieve all customer orders
// @Tags Orders
// @Produce json
// @Security BearerAuth
// @Param status query int false "Filter by order status"
// @Param customer_id query string false "Filter by customer ID"
// @Param date_from query string false "Start date (YYYY-MM-DD)"
// @Param date_to query string false "End date (YYYY-MM-DD)"
// @Success 200 {object} WebResponse{data=[]CustomerOrder} "Orders retrieved"
// @Router /orders [get]
func GetAllOrdersEndpoint() {}

// GetOrderById godoc
// @Summary Get order by ID
// @Description Retrieve order details by order ID
// @Tags Orders
// @Produce json
// @Security BearerAuth
// @Param orderId path string true "Order ID"
// @Success 200 {object} WebResponse{data=CustomerOrder} "Order details"
// @Failure 404 {object} WebResponse "Order not found"
// @Router /orders/{orderId} [get]
func GetOrderByIdEndpoint() {}

// GetTempCart godoc
// @Summary Get temporary cart
// @Description Retrieve temporary cart items for current user
// @Tags Cart
// @Produce json
// @Security BearerAuth
// @Param user_id query string true "User ID"
// @Success 200 {object} WebResponse{data=[]TempCart} "Cart items retrieved"
// @Router /temp_cart [get]
func GetTempCartEndpoint() {}

// AddToTempCart godoc
// @Summary Add to cart
// @Description Add item to temporary cart
// @Tags Cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body TempCart true "Cart item"
// @Success 201 {object} WebResponse{data=TempCart} "Item added to cart"
// @Failure 400 {object} WebResponse "Invalid request"
// @Router /temp_cart [post]
func AddToTempCartEndpoint() {}

// UpdateTempCart godoc
// @Summary Update cart item
// @Description Update quantity of cart item
// @Tags Cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param cartId path string true "Cart item ID"
// @Param request body TempCart true "Updated cart item"
// @Success 200 {object} WebResponse{data=TempCart} "Cart item updated"
// @Failure 400 {object} WebResponse "Invalid request"
// @Router /temp_cart/{cartId} [put]
func UpdateTempCartEndpoint() {}

// DeleteTempCart godoc
// @Summary Delete cart item
// @Description Remove item from temporary cart
// @Tags Cart
// @Produce json
// @Security BearerAuth
// @Param cartId path string true "Cart item ID"
// @Success 200 {object} WebResponse "Cart item deleted"
// @Router /temp_cart/{cartId} [delete]
func DeleteTempCartEndpoint() {}

// ClearTempCart godoc
// @Summary Clear cart
// @Description Clear all items from temporary cart
// @Tags Cart
// @Produce json
// @Security BearerAuth
// @Param user_id query string true "User ID"
// @Success 200 {object} WebResponse "Cart cleared"
// @Router /temp_cart/clear [delete]
func ClearTempCartEndpoint() {}
