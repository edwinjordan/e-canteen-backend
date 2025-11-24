package dashboard_customer_repository

import (
	"github.com/edwinjordan/e-canteen-backend/entity"
	"gorm.io/gorm"
)

type dashboardCustomerRepository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *dashboardCustomerRepository {
	return &dashboardCustomerRepository{DB: db}
}

func (repository *dashboardCustomerRepository) GetTotalSales(year int, customer_id string) (float64, error) {
	var totalSales float64
	err := repository.DB.Table("customer_orders").
		Select("COALESCE(SUM(order_total), 0) as total_sales").
		Where("order_status = ? AND YEAR(order_create_at) = ? AND order_customer_id = ?", 1, year, customer_id).
		Scan(&totalSales).Error
	return totalSales, err
}

func (repository *dashboardCustomerRepository) GetTotalTransactions(year int, customer_id string) (int, error) {
	var count int64
	err := repository.DB.Table("customer_orders").
		Where("order_status = ? AND YEAR(order_create_at) = ? AND order_customer_id = ?", 1, year, customer_id).
		Count(&count).Error
	return int(count), err
}

func (repository *dashboardCustomerRepository) GetTotalCustomers(year int) (int, error) {
	var count int64
	err := repository.DB.Table("customers").
		Where("customer_status = ? AND YEAR(customer_create_at) = ?", 1, year).
		Count(&count).Error
	return int(count), err
}

func (repository *dashboardCustomerRepository) GetTotalProductCustomers(year int, customer_id string) (int, error) {
	var count int64
	err := repository.DB.Table("products").
		Joins("LEFT JOIN product_varians ON products.product_id = product_varians.product_id").
		Joins("LEFT JOIN customer_order_details ON product_varians.product_varian_id = customer_order_details.order_detail_product_varian_id").
		Joins("LEFT JOIN customer_orders ON customer_order_details.order_detail_parent_id = customer_orders.order_id").
		Where("products.product_delete_at IS NULL AND customer_orders.order_status = ? AND YEAR(customer_orders.order_create_at) = ? AND customer_orders.order_customer_id = ?", 1, year, customer_id).
		Distinct("products.product_id").
		Count(&count).Error
	return int(count), err
}

func (repository *dashboardCustomerRepository) GetMonthlySales(year int, customer_id string) ([]entity.MonthlySalesData, error) {
	var results []entity.MonthlySalesData

	err := repository.DB.Raw(`
		SELECT 
			MONTH(order_create_at) as month,
			YEAR(order_create_at) as year,
			COALESCE(SUM(order_total), 0) as total_sales
		FROM customer_orders
		WHERE YEAR(order_create_at) = ? AND order_status = 1 AND order_customer_id = ?
		GROUP BY YEAR(order_create_at), MONTH(order_create_at)
		ORDER BY month
	`, year, customer_id).Scan(&results).Error

	if err != nil {
		return nil, err
	}

	// Map month names
	monthNames := []string{"", "January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December"}

	// Create a map for quick lookup
	salesMap := make(map[int]entity.MonthlySalesData)
	for _, r := range results {
		salesMap[r.Month] = r
	}

	// Fill in all 12 months
	var monthlySales []entity.MonthlySalesData
	for i := 1; i <= 12; i++ {
		if data, exists := salesMap[i]; exists {
			data.MonthName = monthNames[i]
			monthlySales = append(monthlySales, data)
		} else {
			monthlySales = append(monthlySales, entity.MonthlySalesData{
				Month:      i,
				MonthName:  monthNames[i],
				Year:       year,
				TotalSales: 0,
			})
		}
	}

	return monthlySales, nil
}

func (repository *dashboardCustomerRepository) GetTopProducts(limit int, customer_id string, year int) ([]entity.TopProductData, error) {
	var results []entity.TopProductData

	err := repository.DB.Raw(`
		SELECT
                        p.product_id,
                        p.product_name,
                        v.varian_name as variant_name,
                        COALESCE(SUM(od.order_detail_qty), 0) as total_quantity,
                        COALESCE(SUM(od.order_detail_subtotal), 0) as total_sales
                FROM customer_order_details od
                LEFT JOIN customer_orders t ON od.order_detail_parent_id = t.order_id
              	LEFT JOIN product_varians v ON od.order_detail_product_varian_id = v.product_varian_id
                LEFT JOIN products p ON v.product_id = p.product_id
                WHERE t.order_status = 1 AND t.order_customer_id = ? AND YEAR(t.order_create_at) = ?
                GROUP BY p.product_id, p.product_name, v.varian_name
                ORDER BY total_quantity DESC
		LIMIT ?
	`, customer_id, year, limit).Scan(&results).Error

	return results, err
}
