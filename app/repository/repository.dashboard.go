package repository

import "github.com/edwinjordan/e-canteen-backend/entity"

type DashboardRepository interface {
	GetTotalSales() (float64, error)
	GetTotalTransactions() (int, error)
	GetTotalCustomers() (int, error)
	GetMonthlySales(year int) ([]entity.MonthlySalesData, error)
	GetTopProducts(limit int) ([]entity.TopProductData, error)
}
