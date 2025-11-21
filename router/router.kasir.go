package router

import (
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/edwinjordan/e-canteen-backend/app/usecase/usecase_transaction"
	"github.com/edwinjordan/e-canteen-backend/app/usecase/usecase_user"
	"github.com/edwinjordan/e-canteen-backend/repository/order_repository"
	"github.com/edwinjordan/e-canteen-backend/repository/stock_repository"
	"github.com/edwinjordan/e-canteen-backend/repository/tempcart_repository"
	"github.com/edwinjordan/e-canteen-backend/repository/transaction_repository"
	"github.com/edwinjordan/e-canteen-backend/repository/user_repository"
	"github.com/edwinjordan/e-canteen-backend/repository/varian_repository"
	"github.com/edwinjordan/e-canteen-backend/repository/version_repository"
	"gorm.io/gorm"
)

// KasirRouter sets up kasir and transaction routes
func KasirRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	tempCartRepo := tempcart_repository.New(db)
	userRepository := user_repository.New(db)
	userLogRepository := user_repository.NewLog(db)
	versionRepository := version_repository.New(db)
	userController := usecase_user.NewUseCase(userRepository, userLogRepository, tempCartRepo, versionRepository, validate)

	// @Summary User login
	// @Description Authenticate user and get JWT token
	// @Tags Authentication
	// @Accept json
	// @Produce json
	// @Param request body entity.Login true "Login credentials"
	// @Success 200 {object} handler.WebResponse{data=string} "Login successful with JWT token"
	// @Failure 401 {object} handler.WebResponse "Invalid credentials"
	// @Router /kasir/login [post]
	router.HandleFunc("/api/kasir/login", userController.DoLogin).Methods("POST")

	// @Summary User logout
	// @Description Logout user and invalidate token
	// @Tags Authentication
	// @Accept json
	// @Produce json
	// @Security BearerAuth
	// @Success 200 {object} handler.WebResponse "Logout successful"
	// @Router /kasir/logout [put]
	router.HandleFunc("/api/kasir/logout", userController.DoLogout).Methods("PUT")

	// @Summary Verify token
	// @Description Validate bearer token and return user data
	// @Tags Authentication
	// @Produce json
	// @Security BearerAuth
	// @Success 200 {object} handler.WebResponse "Token valid with user data"
	// @Failure 401 {object} handler.WebResponse "Unauthorized"
	// @Router /verify-token [get]
	router.HandleFunc("/api/verify-token", userController.VerifyToken).Methods("GET")

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
	router.HandleFunc("/api/user/fcm", userController.UpdateFcm).Methods("PUT")

	// @Summary Get admin app version
	// @Description Get current admin application version
	// @Tags Version
	// @Produce json
	// @Success 200 {object} handler.WebResponse{data=entity.VersionAdmin} "Version info"
	// @Router /kasir/version [get]
	router.HandleFunc("/api/kasir/version", userController.GetVersionAdmin).Methods("GET")

	// @Summary Get shop app version
	// @Description Get current customer shop application version
	// @Tags Version
	// @Produce json
	// @Success 200 {object} handler.WebResponse{data=entity.VersionShop} "Version info"
	// @Router /shop/version [get]
	router.HandleFunc("/api/shop/version", userController.GetVersionShop).Methods("GET")

	// @Summary Check maintenance mode
	// @Description Check if system is in maintenance mode
	// @Tags Version
	// @Produce json
	// @Param confCode path string true "Configuration code"
	// @Success 200 {object} handler.WebResponse "Maintenance status"
	// @Router /check_maintenance_mode/{confCode} [get]
	router.HandleFunc("/api/check_maintenance_mode/{confCode}", userController.CheckMaintenanceMode).Methods("GET")

	/* transaction */
	transRepo := transaction_repository.NewTrans(db)
	transDetailRepo := transaction_repository.NewTransDetail(db)
	stockBoothRepo := stock_repository.NewBooth(db)
	varianRepo := varian_repository.New(db)
	orderRepository := order_repository.NewOrder(db)
	orderDetailRepository := order_repository.NewOrderDetail(db)
	transController := usecase_transaction.NewUseCase(transRepo, transDetailRepo, tempCartRepo, stockBoothRepo, varianRepo, orderRepository, orderDetailRepository, validate)

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
	router.HandleFunc("/api/kasir/transaction", transController.Create).Methods("POST")

	// @Summary Get all transactions
	// @Description Retrieve all transactions
	// @Tags Transactions
	// @Produce json
	// @Security BearerAuth
	// @Success 200 {object} handler.WebResponse{data=[]entity.Transaction} "Transactions list"
	// @Router /transaction [get]
	router.HandleFunc("/api/transaction", transController.FindAll).Methods("GET")

	// @Summary Get transaction by ID
	// @Description Retrieve transaction details by ID
	// @Tags Transactions
	// @Produce json
	// @Security BearerAuth
	// @Param transId path string true "Transaction ID"
	// @Success 200 {object} handler.WebResponse{data=entity.Transaction} "Transaction details"
	// @Failure 404 {object} handler.WebResponse "Transaction not found"
	// @Router /kasir/transaction/{transId} [get]
	router.HandleFunc("/api/kasir/transaction/{transId}", transController.FindById).Methods("GET")

	// @Summary Get transaction details
	// @Description Retrieve transaction detail items
	// @Tags Transactions
	// @Produce json
	// @Security BearerAuth
	// @Param trans_id query string true "Transaction ID"
	// @Success 200 {object} handler.WebResponse{data=[]entity.TransactionDetail} "Transaction details"
	// @Router /kasir/transaction_detail [get]
	router.HandleFunc("/api/kasir/transaction_detail", transController.GetTransDetail).Methods("GET")

	// @Summary Get transaction summary
	// @Description Get transaction summary and statistics
	// @Tags Transactions
	// @Produce json
	// @Security BearerAuth
	// @Success 200 {object} handler.WebResponse "Transaction summary"
	// @Router /kasir/transaction_summary [get]
	router.HandleFunc("/api/kasir/transaction_summary", transController.GetTransactionSummary).Methods("GET")
}
