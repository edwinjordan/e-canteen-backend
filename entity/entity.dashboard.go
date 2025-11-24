package entity

type DashboardStats struct {
	TotalSales            float64            `json:"total_sales"`
	TotalTransactions     int                `json:"total_transactions"`
	TotalCustomers        int                `json:"total_customers"`
	TotalProductCustomers int                `json:"total_product_customers"`
	MonthlySales          []MonthlySalesData `json:"monthly_sales"`
	TopProducts           []TopProductData   `json:"top_products"`
}

type MonthlySalesData struct {
	Month      int     `json:"month"`
	MonthName  string  `json:"month_name"`
	Year       int     `json:"year"`
	TotalSales float64 `json:"total_sales"`
}

type TopProductData struct {
	ProductId     string  `json:"product_id"`
	ProductName   string  `json:"product_name"`
	VariantName   string  `json:"variant_name"`
	TotalQuantity int     `json:"total_quantity"`
	TotalSales    float64 `json:"total_sales"`
}
