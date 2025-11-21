package dashboard_repository

import (
	"github.com/edwinjordan/e-canteen-backend/entity"
	"gorm.io/gorm"
)

type dashboardRepository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *dashboardRepository {
	return &dashboardRepository{DB: db}
}

func (repository *dashboardRepository) GetTotalSales() (float64, error) {
	var totalSales float64
	err := repository.DB.Table("transactions").
		Select("COALESCE(SUM(trans_total), 0) as total_sales").
		Where("trans_status = ?", 1).
		Scan(&totalSales).Error
	return totalSales, err
}

func (repository *dashboardRepository) GetTotalTransactions() (int, error) {
	var count int64
	err := repository.DB.Table("transactions").
		Where("trans_status = ?", 1).
		Count(&count).Error
	return int(count), err
}

func (repository *dashboardRepository) GetTotalCustomers() (int, error) {
	var count int64
	err := repository.DB.Table("customers").
		Where("customer_status = ?", 1).
		Count(&count).Error
	return int(count), err
}

func (repository *dashboardRepository) GetMonthlySales(year int) ([]entity.MonthlySalesData, error) {
	var results []entity.MonthlySalesData

	err := repository.DB.Raw(`
		SELECT 
			MONTH(trans_create_at) as month,
			YEAR(trans_create_at) as year,
			COALESCE(SUM(trans_total), 0) as total_sales
		FROM transactions
		WHERE YEAR(trans_create_at) = ? AND trans_status = 1
		GROUP BY YEAR(trans_create_at), MONTH(trans_create_at)
		ORDER BY month
	`, year).Scan(&results).Error

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

func (repository *dashboardRepository) GetTopProducts(limit int) ([]entity.TopProductData, error) {
	var results []entity.TopProductData

	err := repository.DB.Raw(`
		SELECT
                        p.product_id,
                        p.product_name,
                        v.varian_name as variant_name,
                        COALESCE(SUM(td.trans_detail_qty), 0) as total_quantity,
                        COALESCE(SUM(td.trans_detail_subtotal), 0) as total_sales
                FROM transaction_details td
                LEFT JOIN transactions t ON td.trans_detail_parent_id = t.trans_id
                LEFT JOIN product_varians v ON td.trans_detail_product_varian_id = v.product_varian_id
                LEFT JOIN products p ON v.product_id = p.product_id
                WHERE t.trans_status = 1
                GROUP BY p.product_id, p.product_name, v.varian_name
                ORDER BY total_quantity DESC
		LIMIT ?
	`, limit).Scan(&results).Error

	return results, err
}
