package repository

import "github.com/edwinjordan/e-canteen-backend/entity"

type DashboardCustomerRepository interface {
	GetTotalSales(year int, customer_id string) (float64, error)
	GetTotalTransactions(year int, customer_id string) (int, error)
	GetTotalCustomers(year int) (int, error)
	GetTotalProductCustomers(year int, customer_id string) (int, error)
	GetMonthlySales(year int, customer_id string) ([]entity.MonthlySalesData, error)
	GetTopProducts(limit int, customer_id string, year int) ([]entity.TopProductData, error)
}
