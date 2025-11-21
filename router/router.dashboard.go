package router

import (
	"github.com/gorilla/mux"
	"github.com/edwinjordan/e-canteen-backend/app/usecase/usecase_dashboard"
	"github.com/edwinjordan/e-canteen-backend/repository/dashboard_repository"
	"gorm.io/gorm"
)

// DashboardRouter sets up dashboard routes
func DashboardRouter(db *gorm.DB, router *mux.Router) {
	dashboardRepository := dashboard_repository.New(db)
	dashboardController := usecase_dashboard.NewUseCase(dashboardRepository)

	router.HandleFunc("/api/dashboard/stats", dashboardController.GetDashboardStats).Methods("GET")
}
