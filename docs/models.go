package docs

// Common response models

// WebResponse is the standard API response structure
// @Description Standard API response wrapper
type WebResponse struct {
    Error   bool        `json:"error" example:"false"`
    Message string      `json:"message" example:"Success"`
    Data    interface{} `json:"data"`
}

// LoginRequest represents the login payload
// @Description Login credentials
type LoginRequest struct {
	Email    string `json:"email" example:"admin@ecanteen.com" validate:"required,email"`
	Password string `json:"password" example:"admin123" validate:"required"`
}

// LoginResponse represents the login response
// @Description Login response with user data and token
type LoginResponse struct {
    UserId              string      `json:"user_id" example:"550e8400-e29b-41d4-a716-446655440001"`
    UserName            string      `json:"user_name" example:"Ahmad Fauzi"`
    UserEmail           string      `json:"user_email" example:"admin@ecanteen.com"`
    UserHasMobileAccess int         `json:"user_has_mobile_access" example:"1"`
    UserRoleId          string      `json:"user_role_id" example:"550e8400-e29b-41d4-a716-446655440010"`
    Token               string      `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
    Pegawai             Pegawai     `json:"pegawai"`
}

// LogoutRequest represents the logout payload
// @Description Logout request with token
type LogoutRequest struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." validate:"required"`
}

// VersionResponse represents version information
// @Description Application version information
type VersionResponse struct {
	VersionId       int    `json:"version_id" example:"1"`
	VersionNumber   string `json:"version_number" example:"1.2.0"`
	VersionCode     int    `json:"version_code" example:"3"`
	VersionChagelog string `json:"version_chagelog" example:"Improvement: Better reporting"`
	VersionDatetime string `json:"version_datetime" example:"2025-11-14T10:00:00Z"`
}

// Product represents a product
// @Description Product information
type Product struct {
    ProductId         string      `json:"product_id" example:"PRD001"`
    ProductCode       string      `json:"product_code" example:"PRD-001"`
    ProductName       string      `json:"product_name" example:"Nasi Goreng Spesial"`
    ProductCategoryId string      `json:"product_category_id" example:"550e8400-e29b-41d4-a716-446655440020"`
    ProductDesc       string      `json:"product_desc" example:"Nasi goreng dengan telur, ayam, dan sayuran"`
    CategoryName      string      `json:"category_name" example:"Makanan Berat"`
    ProductCreateAt   string      `json:"product_create_at" example:"2025-11-14T10:00:00Z"`
    ProductUpdateAt   string      `json:"product_update_at" example:"2025-11-14T10:00:00Z"`
    ProductDeleteAt   string      `json:"product_delete_at"`
    ProductPhoto      string      `json:"product_photo" example:"product1.jpg"`
    Varian            []ProductVariant `json:"varian"`
}

// Category represents a product category
// @Description Product category information
type Category struct {
	CategoryId       string `json:"category_id" example:"550e8400-e29b-41d4-a716-446655440020"`
	CategoryName     string `json:"category_name" example:"Makanan Berat"`
	CategoryDeleteAt string `json:"category_delete_at"`
}

// ProductVariant represents a product variant
// @Description Product variant with pricing and stock information
type ProductVariant struct {
	ProductVarianId           string `json:"product_varian_id" example:"550e8400-e29b-41d4-a716-446655440030"`
	ProductId                 string `json:"product_id" example:"PRD001"`
	ProductName               string `json:"product_name" example:"Nasi Goreng Spesial"`
	VarianName                string `json:"varian_name" example:"Reguler"`
	ProductVarianPrice        int    `json:"product_varian_price" example:"15000"`
	ProductVarianQtyBooth     int    `json:"product_varian_qty_booth" example:"50"`
	ProductVarianQtyWarehouse string `json:"product_varian_qty_warehouse" example:"100"`
	VarianId                  string `json:"varian_id" example:"VAR001"`
	ProductVarianQtyLeft      int    `json:"product_varian_qty_left" example:"50"`
}

// Customer represents a customer
// @Description Customer information
type Customer struct {
    CustomerId             string      `json:"customer_id" example:"550e8400-e29b-41d4-a716-446655440040"`
    CustomerCode           string      `json:"customer_code" example:"CST001"`
    CustomerName           string      `json:"customer_name" example:"Andi Wijaya"`
    CustomerGender         string      `json:"customer_gender" example:"L"`
    CustomerPhonenumber    string      `json:"customer_phonenumber" example:"081234560001"`
    CustomerEmail          string      `json:"customer_email" example:"andi@student.com"`
    CustomerDob            string      `json:"customer_dob" example:"2003-05-15"`
    CustomerProfilePic     string      `json:"customer_profile_pic" example:"profile1.jpg"`
    CustomerClass          string      `json:"customer_class" example:"3A"`
    CustomerMajorId        string      `json:"customer_major_id" example:"550e8400-e29b-41d4-a716-446655440050"`
    CustomerProfilePicPath string      `json:"customer_profile_pic_path" example:"https://canteensekolah.biz.id/profiles/profile1.jpg"`
    CustomerStatus         int         `json:"customer_status" example:"1"`
    CustomerLastStatus     int         `json:"customer_last_status" example:"1"`
    CustomerCreateAt       string      `json:"customer_create_at" example:"2025-11-14T10:00:00Z"`
    CustomerUpdateAt       string      `json:"customer_update_at" example:"2025-11-14T10:00:00Z"`
    Major                  Major       `json:"jurusan"`
    Address                []CustomerAddress `json:"alamat"`
}

// CustomerAddress represents a customer address
// @Description Customer delivery address
type CustomerAddress struct {
	AddressId         string `json:"address_id" example:"550e8400-e29b-41d4-a716-446655440060"`
	AddressCustomerId string `json:"address_customer_id" example:"550e8400-e29b-41d4-a716-446655440040"`
	AddressText       string `json:"address_text" example:"Jl. Raya No. 123, RT 01/RW 02"`
	AddressName       string `json:"address_name" example:"Rumah"`
	AddressProvinceId string `json:"address_province_id" example:"11"`
	AddressProvince   string `json:"address_province" example:"Jawa Barat"`
	AddressCityId     string `json:"address_city_id" example:"1101"`
	AddressCity       string `json:"address_city" example:"Bandung"`
	AddressDistrictId string `json:"address_district_id" example:"110101"`
	AddressDistrict   string `json:"address_district" example:"Coblong"`
	AddressVillageId  string `json:"address_village_id" example:"11010101"`
	AddressVillage    string `json:"address_village" example:"Lebak Siliwangi"`
	AddressPostalCode string `json:"address_postal_code" example:"40132"`
	AddressMain       int    `json:"address_main" example:"1"`
	AddressCreateAt   string `json:"address_create_at" example:"2025-11-14T10:00:00Z"`
	AddressUpdateAt   string `json:"address_update_at" example:"2025-11-14T10:00:00Z"`
}

// Transaction represents a transaction
// @Description Transaction/cashier sale information
type Transaction struct {
    TransId            string      `json:"trans_id" example:"550e8400-e29b-41d4-a716-446655440070"`
    TransUserId        string      `json:"trans_user_id" example:"550e8400-e29b-41d4-a716-446655440001"`
    TransCustomerId    string      `json:"trans_customer_id" example:"550e8400-e29b-41d4-a716-446655440040"`
    TransOrderId       string      `json:"trans_order_id" example:"550e8400-e29b-41d4-a716-446655440080"`
    TransInvoice       string      `json:"trans_invoice" example:"INV-20251114-001"`
    TransQtyTotal      int         `json:"trans_qty_total" example:"5"`
    TransProductTotal  int         `json:"trans_product_total" example:"3"`
    TransSubtotal      float64     `json:"trans_subtotal" example:"45000"`
    TransDiscount      float64     `json:"trans_discount" example:"5000"`
    TransTotal         float64     `json:"trans_total" example:"40000"`
    TransReceivedTotal float64     `json:"trans_received_total" example:"50000"`
    TransRefundTotal   float64     `json:"trans_refund_total" example:"10000"`
    TransStatus        int         `json:"trans_status" example:"1"`
    TransCreateAt      string      `json:"trans_create_at" example:"2025-11-14T10:00:00Z"`
    TransDetail        []TransactionDetail `json:"trans_detail"`
    Customer           CustomerResponse    `json:"customer"`
    User               User                `json:"user"`
}

// TransactionDetail represents a transaction detail item
// @Description Individual item in a transaction
type TransactionDetail struct {
    TransDetailId              string  `json:"trans_detail_id" example:"550e8400-e29b-41d4-a716-446655440090"`
    TransDetailParentId        string  `json:"trans_detail_parent_id" example:"550e8400-e29b-41d4-a716-446655440070"`
    TransDetailProductVarianId string  `json:"trans_detail_product_varian_id" example:"550e8400-e29b-41d4-a716-446655440030"`
    TransDetailQty             int     `json:"trans_detail_qty" example:"2"`
    TransDetailPrice           float64 `json:"trans_detail_price" example:"15000"`
    TransDetailSubtotal        float64 `json:"trans_detail_subtotal" example:"30000"`
}

// CustomerOrderDetail represents order detail item
// @Description Item in a customer order
type CustomerOrderDetail struct {
    OrderDetailId              string  `json:"order_detail_id" example:"550e8400-e29b-41d4-a716-446655440200"`
    OrderDetailParentId        string  `json:"order_detail_parent_id" example:"550e8400-e29b-41d4-a716-446655440080"`
    OrderDetailProductVarianId string  `json:"order_detail_product_varian_id" example:"550e8400-e29b-41d4-a716-446655440030"`
    OrderDetailQty             int     `json:"order_detail_qty" example:"2"`
    OrderDetailPrice           float64 `json:"order_detail_price" example:"15000"`
    OrderDetailSubtotal        float64 `json:"order_detail_subtotal" example:"30000"`
}

// CreateTransactionRequest represents the transaction creation payload
// @Description Create new transaction request
type CreateTransactionRequest struct {
	TransUserId        string  `json:"trans_user_id" example:"550e8400-e29b-41d4-a716-446655440001" validate:"required"`
	TransCustomerId    string  `json:"trans_customer_id" example:"550e8400-e29b-41d4-a716-446655440040"`
	TransOrderId       string  `json:"trans_order_id" example:"550e8400-e29b-41d4-a716-446655440080"`
	TransSubtotal      float64 `json:"trans_subtotal" example:"45000" validate:"required"`
	TransDiscount      float64 `json:"trans_discount" example:"5000"`
	TransTotal         float64 `json:"trans_total" example:"40000" validate:"required"`
	TransReceivedTotal float64 `json:"trans_received_total" example:"50000" validate:"required"`
}

// TempCart represents a temporary cart item
// @Description Temporary cart for building transactions
type TempCart struct {
	TempCartId              string `json:"temp_cart_id" example:"550e8400-e29b-41d4-a716-446655440100"`
	TempCartOrderId         string `json:"temp_cart_order_id" example:"550e8400-e29b-41d4-a716-446655440080"`
	TempCartProductVarianId string `json:"temp_cart_product_varian_id" example:"550e8400-e29b-41d4-a716-446655440030"`
	TempCartUserId          string `json:"temp_cart_user_id" example:"550e8400-e29b-41d4-a716-446655440001"`
	TempCartQty             int    `json:"temp_cart_qty" example:"2"`
}

// Major represents a major/department
// @Description Student major or department
type Major struct {
	MajorId   string `json:"major_id" example:"550e8400-e29b-41d4-a716-446655440050"`
	MajorName string `json:"major_name" example:"Teknik Informatika"`
}

// CustomerOrder represents a customer order
// @Description Customer order information
type CustomerOrder struct {
    OrderId                string      `json:"order_id" example:"550e8400-e29b-41d4-a716-446655440080"`
    OrderCustomerId        string      `json:"order_customer_id" example:"550e8400-e29b-41d4-a716-446655440040"`
    OrderInvNumber         string      `json:"order_inv_number" example:"ORD-20251114-001"`
    OrderAddressId         string      `json:"order_address_id" example:"550e8400-e29b-41d4-a716-446655440060"`
    OrderDeliveryType      string      `json:"order_delivery_type" example:"DELIVERY"`
    OrderTotalItem         int         `json:"order_total_item" example:"5"`
    OrderSubtotal          float64     `json:"order_subtotal" example:"45000"`
    OrderDiscount          float64     `json:"order_discount" example:"5000"`
    OrderTotal             float64     `json:"order_total" example:"40000"`
    OrderNotes             string      `json:"order_notes" example:"Extra pedas"`
    OrderCancelNotes       string      `json:"order_cancel_notes"`
    OrderStatus            int         `json:"order_status" example:"1"`
    OrderProcessedDatetime string      `json:"order_processed_datetime" example:"2025-11-14T10:00:00Z"`
    OrderProcessedBy       string      `json:"order_processed_by" example:"550e8400-e29b-41d4-a716-446655440001"`
    OrderFinishedDatetime  string      `json:"order_finished_datetime" example:"2025-11-14T10:30:00Z"`
    OrderFinishedBy        string      `json:"order_finished_by" example:"550e8400-e29b-41d4-a716-446655440001"`
    OrderCreateAt          string      `json:"order_create_at" example:"2025-11-14T10:00:00Z"`
    OrderDetail            []CustomerOrderDetail `json:"order_detail"`
    Customer               CustomerResponse      `json:"customer"`
    Address                CustomerAddress       `json:"address"`
}

// User represents system user
// @Description Cashier/admin user
type User struct {
    UserId              string  `json:"user_id" example:"550e8400-e29b-41d4-a716-446655440001"`
    UserName            string  `json:"user_name" example:"Ahmad Fauzi"`
    UserEmail           string  `json:"user_email" example:"admin@ecanteen.com"`
    UserPegawaiId       string  `json:"user_pegawai_id" example:"550e8400-e29b-41d4-a716-446655440002"`
    UserHasMobileAccess int     `json:"user_has_mobile_access" example:"1"`
    UserRoleId          string  `json:"user_role_id" example:"550e8400-e29b-41d4-a716-446655440010"`
    Pegawai             Pegawai `json:"pegawai"`
}

// Pegawai represents employee data
// @Description Employee profile linked to user
type Pegawai struct {
    PegawaiId          string `json:"pegawai_id" example:"550e8400-e29b-41d4-a716-446655440002"`
    PegawaiCode        string `json:"pegawai_code" example:"EMP-001"`
    PegawaiName        string `json:"pegawai_name" example:"Budi"`
    PegawaiGender      string `json:"pegawai_gender" example:"L"`
    PegawaiPhonenumber string `json:"pegawai_phonenumber" example:"08123456789"`
}

// CustomerResponse mirrors Customer with hidden password field
// @Description Customer response entity
type CustomerResponse struct {
    CustomerId             string           `json:"customer_id" example:"550e8400-e29b-41d4-a716-446655440040"`
    CustomerCode           string           `json:"customer_code" example:"CST001"`
    CustomerName           string           `json:"customer_name" example:"Andi Wijaya"`
    CustomerGender         string           `json:"customer_gender" example:"L"`
    CustomerPhonenumber    string           `json:"customer_phonenumber" example:"081234560001"`
    CustomerEmail          string           `json:"customer_email" example:"andi@student.com"`
    CustomerDob            string           `json:"customer_dob" example:"2003-05-15"`
    CustomerProfilePic     string           `json:"customer_profile_pic" example:"profile1.jpg"`
    CustomerClass          string           `json:"customer_class" example:"3A"`
    CustomerMajorId        string           `json:"customer_major_id" example:"550e8400-e29b-41d4-a716-446655440050"`
    CustomerProfilePicPath string           `json:"customer_profile_pic_path" example:"https://canteensekolah.biz.id/profiles/profile1.jpg"`
    CustomerStatus         int              `json:"customer_status" example:"1"`
    CustomerLastStatus     int              `json:"customer_last_status" example:"1"`
    CustomerCreateAt       string           `json:"customer_create_at" example:"2025-11-14T10:00:00Z"`
    CustomerUpdateAt       string           `json:"customer_update_at" example:"2025-11-14T10:00:00Z"`
    Major                  Major            `json:"jurusan"`
    Address                []CustomerAddress `json:"alamat"`
}
