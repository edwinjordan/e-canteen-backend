package usecase_dashboard

import (
	"net/http"
	"strconv"
	"time"

	"github.com/edwinjordan/e-canteen-backend/app/repository"
	"github.com/edwinjordan/e-canteen-backend/entity"
	"github.com/edwinjordan/e-canteen-backend/handler"
	"github.com/edwinjordan/e-canteen-backend/pkg/helpers"
)

type UseCaseImpl struct {
	DashboardRepository repository.DashboardRepository
}

func NewUseCase(dashboardRepository repository.DashboardRepository) DashboardUseCase {
	return &UseCaseImpl{
		DashboardRepository: dashboardRepository,
	}
}

func (controller *UseCaseImpl) GetDashboardStats(w http.ResponseWriter, r *http.Request) {
	// Get year from query parameter, default to current year
	yearStr := r.URL.Query().Get("year")
	year := time.Now().Year()
	if yearStr != "" {
		if y, err := strconv.Atoi(yearStr); err == nil {
			year = y
		}
	}

	// Get total sales
	totalSales, err := controller.DashboardRepository.GetTotalSales()
	helpers.PanicIfError(err)

	// Get total transactions
	totalTransactions, err := controller.DashboardRepository.GetTotalTransactions()
	helpers.PanicIfError(err)

	// Get total customers
	totalCustomers, err := controller.DashboardRepository.GetTotalCustomers()
	helpers.PanicIfError(err)

	// Get monthly sales for the year
	monthlySales, err := controller.DashboardRepository.GetMonthlySales(year)
	helpers.PanicIfError(err)

	// Get top 10 products
	topProducts, err := controller.DashboardRepository.GetTopProducts(10)
	helpers.PanicIfError(err)

	// Build dashboard stats response
	dashboardStats := entity.DashboardStats{
		TotalSales:        totalSales,
		TotalTransactions: totalTransactions,
		TotalCustomers:    totalCustomers,
		MonthlySales:      monthlySales,
		TopProducts:       topProducts,
	}

	webResponse := handler.WebResponse{
		Error:   false,
		Message: "Berhasil mengambil data dashboard",
		Data:    dashboardStats,
	}
	helpers.WriteToResponseBody(w, webResponse)
}
