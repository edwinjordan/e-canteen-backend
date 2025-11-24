package usecase_dashboard_customer

import "net/http"

type DashboardCustomerUseCase interface {
	GetDashboardStats(w http.ResponseWriter, r *http.Request)
}
