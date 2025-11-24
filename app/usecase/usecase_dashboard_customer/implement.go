package usecase_dashboard_customer

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
	DashboardCustomerRepository repository.DashboardCustomerRepository
}

func NewUseCase(dashboardCustomerRepository repository.DashboardCustomerRepository) DashboardCustomerUseCase {
	return &UseCaseImpl{
		DashboardCustomerRepository: dashboardCustomerRepository,
	}
}

func (controller *UseCaseImpl) GetDashboardStats(w http.ResponseWriter, r *http.Request) {
	// Get year from query parameter, default to current year
	yearStr := r.URL.Query().Get("year")
	year := time.Now().Year()
	customer_id := r.URL.Query().Get("customer_id")
	if yearStr != "" {
		if y, err := strconv.Atoi(yearStr); err == nil {
			year = y
		}
	}

	// Get total sales
	totalSales, err := controller.DashboardCustomerRepository.GetTotalSales(year, customer_id)
	helpers.PanicIfError(err)

	// Get total transactions
	totalTransactions, err := controller.DashboardCustomerRepository.GetTotalTransactions(year, customer_id)
	helpers.PanicIfError(err)

	// Get total customers
	totalCustomers, err := controller.DashboardCustomerRepository.GetTotalCustomers(year)
	helpers.PanicIfError(err)

	// Get total product
	totalProductCustomers, err := controller.DashboardCustomerRepository.GetTotalProductCustomers(year, customer_id)
	helpers.PanicIfError(err)

	// Get monthly sales for the year
	monthlySales, err := controller.DashboardCustomerRepository.GetMonthlySales(year, customer_id)
	helpers.PanicIfError(err)

	// Get top 10 products
	topProducts, err := controller.DashboardCustomerRepository.GetTopProducts(10, customer_id, year)
	helpers.PanicIfError(err)

	// Build dashboard stats response
	dashboardStats := entity.DashboardStats{
		TotalSales:            totalSales,
		TotalTransactions:     totalTransactions,
		TotalCustomers:        totalCustomers,
		TotalProductCustomers: totalProductCustomers,
		MonthlySales:          monthlySales,
		TopProducts:           topProducts,
	}

	webResponse := handler.WebResponse{
		Error:   false,
		Message: "Berhasil mengambil data dashboard",
		Data:    dashboardStats,
	}
	helpers.WriteToResponseBody(w, webResponse)
}
