package repository

import "github.com/edwinjordan/e-canteen-backend/entity"

type DashboardRepository interface {
	GetTotalSales(year int) (float64, error)
	GetTotalTransactions(year int) (int, error)
	GetTotalCustomers(year int) (int, error)
	GetTotalProductCustomers(year int) (int, error)
	GetMonthlySales(year int) ([]entity.MonthlySalesData, error)
	GetTopProducts(limit int, year int) ([]entity.TopProductData, error)
}
