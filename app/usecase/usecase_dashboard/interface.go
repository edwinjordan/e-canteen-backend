package usecase_dashboard

import "net/http"

type DashboardUseCase interface {
	GetDashboardStats(w http.ResponseWriter, r *http.Request)
}
