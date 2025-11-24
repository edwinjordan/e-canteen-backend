package router

import (
	"github.com/edwinjordan/e-canteen-backend/app/usecase/usecase_dashboard_customer"
	"github.com/edwinjordan/e-canteen-backend/repository/dashboard_customer_repository"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// DashboardRouter sets up dashboard routes
func DashboardCustomerRouter(db *gorm.DB, router *mux.Router) {
	dashboardRepository := dashboard_customer_repository.New(db)
	dashboardController := usecase_dashboard_customer.NewUseCase(dashboardRepository)

	router.HandleFunc("/api/dashboard_customer/stats", dashboardController.GetDashboardStats).Methods("GET")
}
