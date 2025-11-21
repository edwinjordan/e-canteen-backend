package entity

type CustomerOrderDetail struct {
	OrderDetailId              string  `json:"order_detail_id"`
	OrderDetailParentId        string  `json:"order_detail_parent_id"`
	OrderDetailProductVarianId string  `json:"order_detail_product_varian_id"`
	OrderDetailQty             int     `json:"order_detail_qty"`
	OrderDetailPrice           float64 `json:"order_detail_price"`
	OrderDetailSubtotal        float64 `json:"order_detail_subtotal"`
}

type ViewOrderDetail struct {
	OrderDetailId              string  `json:"order_detail_id"`
	OrderDetailParentId        string  `json:"order_detail_parent_id"`
	OrderDetailProductVarianId string  `json:"order_detail_product_varian_id"`
	OrderDetailQty             int     `json:"order_detail_qty"`
	OrderDetailPrice           float64 `json:"order_detail_price"`
	OrderDetailSubtotal        float64 `json:"order_detail_subtotal"`
	CustomerName               string  `json:"customer_name"`
	ProductName                string  `json:"product_name"`
	VarianName                 string  `json:"varian_name"`
}

// CreateOrderRequest represents the request payload for creating a customer order
type CreateOrderRequest struct {
	OrderCustomerId   string                     `json:"order_customer_id" example:"1bee61b0-c11f-11f0-9feb-482ae3a0bcc9"`
	OrderDeliveryType string                     `json:"order_delivery_type" example:"pickup"`
	OrderAddressId    string                     `json:"order_address_id,omitempty" example:""`
	OrderNotes        string                     `json:"order_notes,omitempty" example:""`
	OrderDiscount     float64                    `json:"order_discount,omitempty" example:"0"`
	TempCartOrderId   string                     `json:"temp_cart_order_id,omitempty" example:"ORDER-GUEST-1763326901093-JEGMFSW"`
	OrderDetails      []CreateOrderDetailRequest `json:"order_details"`
}

// CreateOrderDetailRequest represents each order item in the request
type CreateOrderDetailRequest struct {
	OrderDetailProductVarianId string  `json:"order_detail_product_varian_id" example:"0058ae56cc7a089dc710b0fff0f1e312"`
	OrderDetailQty             int     `json:"order_detail_qty" example:"1"`
	OrderDetailPrice           float64 `json:"order_detail_price" example:"15000"`
	OrderDetailSubtotal        float64 `json:"order_detail_subtotal" example:"15000"`
}
