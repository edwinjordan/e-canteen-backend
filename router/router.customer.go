package router

import (
	"github.com/edwinjordan/e-canteen-backend/app/service"
	"github.com/edwinjordan/e-canteen-backend/app/usecase/usecase_customer"
	"github.com/edwinjordan/e-canteen-backend/config"
	"github.com/edwinjordan/e-canteen-backend/repository/customer_repository"
	"github.com/edwinjordan/e-canteen-backend/repository/otp_repository"
	"github.com/edwinjordan/e-canteen-backend/repository/tempcart_repository"
	"github.com/edwinjordan/e-canteen-backend/repository/user_repository"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// CustomerRouter sets up customer routes
func CustomerRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	customerRepository := customer_repository.New(db)
	otpRepository := otp_repository.New(db)
	userLogRepository := user_repository.NewLog(db)
	tempCartRepo := tempcart_repository.New(db)

	// Initialize MinIO
	minioClient := config.NewMinioClient()
	minioService := service.NewMinioService(minioClient)

	customerController := usecase_customer.NewUseCase(customerRepository, otpRepository, tempCartRepo, userLogRepository, minioService, validate)

	// @Summary Customer login
	// @Description Authenticate customer and get JWT token
	// @Tags Customer
	// @Accept json
	// @Produce json
	// @Param request body entity.Login true "Login credentials"
	// @Success 200 {object} handler.WebResponse{data=string} "Login successful with JWT token"
	// @Failure 401 {object} handler.WebResponse "Invalid credentials"
	// @Router /customer/login [post]
	router.HandleFunc("/api/customer/login", customerController.DoLogin).Methods("POST")

	// @Summary Add customer log
	// @Description Add customer activity log
	// @Tags Customer
	// @Accept json
	// @Produce json
	// @Security BearerAuth
	// @Success 200 {object} handler.WebResponse "Log added"
	// @Router /customer/addLog [post]
	router.HandleFunc("/api/customer/addLog", customerController.AddLog).Methods("POST")

	// @Summary Customer logout
	// @Description Logout customer and invalidate token
	// @Tags Customer
	// @Accept json
	// @Produce json
	// @Security BearerAuth
	// @Success 200 {object} handler.WebResponse "Logout successful"
	// @Router /customer/logout [post]
	router.HandleFunc("/api/customer/logout", customerController.DoLogout).Methods("POST")

	// @Summary Change customer password
	// @Description Change customer password
	// @Tags Customer
	// @Accept json
	// @Produce json
	// @Security BearerAuth
	// @Success 200 {object} handler.WebResponse "Password changed"
	// @Failure 400 {object} handler.WebResponse "Invalid request"
	// @Router /customer/change_password [post]
	router.HandleFunc("/api/customer/change_password", customerController.ChangePassword).Methods("POST")

	// @Summary Get all customers
	// @Description Retrieve all customers
	// @Tags Customer
	// @Produce json
	// @Security BearerAuth
	// @Success 200 {object} handler.WebResponse{data=[]entity.Customer} "Customers list"
	// @Router /customer [get]
	router.HandleFunc("/api/customer", customerController.FindAll).Methods("GET")

	// @Summary Get customer by ID
	// @Description Retrieve customer details by ID
	// @Tags Customer
	// @Produce json
	// @Security BearerAuth
	// @Param customerId path string true "Customer ID"
	// @Success 200 {object} handler.WebResponse{data=entity.Customer} "Customer details"
	// @Failure 404 {object} handler.WebResponse "Customer not found"
	// @Router /customer/{customerId} [get]
	router.HandleFunc("/api/customer/{customerId}", customerController.FindById).Methods("GET")

	// @Summary Register customer
	// @Description Register a new customer
	// @Tags Customer
	// @Accept json
	// @Produce json
	// @Param request body entity.Customer true "Customer data"
	// @Success 201 {object} handler.WebResponse{data=entity.Customer} "Customer registered"
	// @Failure 400 {object} handler.WebResponse "Invalid request"
	// @Router /customer [post]
	router.HandleFunc("/api/customer", customerController.Register).Methods("POST")

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
	router.HandleFunc("/api/customer/{customerId}", customerController.Update).Methods("PUT")

	// @Summary Delete customer
	// @Description Delete customer by ID
	// @Tags Customer
	// @Produce json
	// @Security BearerAuth
	// @Param customerId path string true "Customer ID"
	// @Success 200 {object} handler.WebResponse "Customer deleted"
	// @Failure 404 {object} handler.WebResponse "Customer not found"
	// @Router /customer/{customerId} [delete]
	router.HandleFunc("/api/customer/{customerId}", customerController.Delete).Methods("DELETE")

	// @Summary Verify OTP
	// @Description Verify OTP for customer registration
	// @Tags Customer
	// @Accept json
	// @Produce json
	// @Success 200 {object} handler.WebResponse "OTP verified"
	// @Failure 400 {object} handler.WebResponse "Invalid OTP"
	// @Router /customer/verifyOtp [post]
	router.HandleFunc("/api/customer/verifyOtp", customerController.VerifyOtp).Methods("POST")

	// @Summary Send OTP for password reset
	// @Description Send OTP to customer for password reset
	// @Tags Customer
	// @Accept json
	// @Produce json
	// @Success 200 {object} handler.WebResponse "OTP sent"
	// @Router /customer/sentOTPResetPassword [post]
	router.HandleFunc("/api/customer/sentOTPResetPassword", customerController.SendOTPResetPassword).Methods("POST")

}
