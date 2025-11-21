package usecase_order

import (
	"database/sql"
	"fmt"
	"html"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/edwinjordan/e-canteen-backend/app/repository"
	"github.com/edwinjordan/e-canteen-backend/config"
	"github.com/edwinjordan/e-canteen-backend/entity"
	"github.com/edwinjordan/e-canteen-backend/handler"
	"github.com/edwinjordan/e-canteen-backend/pkg/exceptions"
	"github.com/edwinjordan/e-canteen-backend/pkg/helpers"
)

type UseCaseImpl struct {
	CustomerOrderRepository       repository.CustomerOrderRepository
	CustomerOrderDetailRepository repository.CustomerOrderDetailRepository
	BoothStockRepository          repository.BoothStockRepository
	VarianRepository              repository.VarianRepository
	TempCartRepository            repository.TempCartRepository
	TransactionRepository         repository.TransactionRepository
	TransactionDetailRepository   repository.TransactionDetailRepository
	UserRepository                repository.UserRepository
	Validate                      *validator.Validate
}

func NewUseCase(
	orderRepo repository.CustomerOrderRepository,
	orderDetailRepo repository.CustomerOrderDetailRepository,
	varianRepo repository.VarianRepository,
	tempCartRepo repository.TempCartRepository,
	transRepo repository.TransactionRepository,
	transDetailRepo repository.TransactionDetailRepository,
	stockBoothRepo repository.BoothStockRepository,
	userRepo repository.UserRepository,
	validate *validator.Validate,
) UseCase {
	return &UseCaseImpl{
		Validate:                      validate,
		CustomerOrderRepository:       orderRepo,
		CustomerOrderDetailRepository: orderDetailRepo,
		VarianRepository:              varianRepo,
		TempCartRepository:            tempCartRepo,
		TransactionRepository:         transRepo,
		TransactionDetailRepository:   transDetailRepo,
		BoothStockRepository:          stockBoothRepo,
		UserRepository:                userRepo,
	}
}

func (controller *UseCaseImpl) Create(w http.ResponseWriter, r *http.Request) {
	dataRequest := map[string]interface{}{}
	helpers.ReadFromRequestBody(r, &dataRequest)

	// Ambil order_details dari request
	orderDetails := dataRequest["order_details"].([]interface{})

	// Hitung total item dan total harga
	totalItem := len(orderDetails)
	var subtotal float64 = 0

	/* check apakah masih ada sisa untuk product tsb */
	for _, v := range orderDetails {
		dt := v.(map[string]interface{})
		detail, _ := controller.VarianRepository.FindById(r.Context(), dt["order_detail_product_varian_id"].(string))
		/* jika ada salah satu barang memiliki sisa kurang dari yang dipesan maka langsung batalkan pesanan */
		if detail.ProductVarianQtyBooth < int(dt["order_detail_qty"].(float64)) {
			panic(exceptions.NewBadRequestError("Tidak dapat membuat order karena salah satu barang dikeranjang tidak memiliki jumlah tersisa yang cukup"))
		}
		subtotal += dt["order_detail_subtotal"].(float64)
	}

	// Ambil nilai discount jika ada, default 0
	discount := 0.0
	if dataRequest["order_discount"] != nil {
		discount = dataRequest["order_discount"].(float64)
	}

	// Hitung total
	total := subtotal - discount

	// Ambil order_address_id jika ada
	addressId := ""
	// if dataRequest["order_address_id"] != nil && dataRequest["order_address_id"].(string) != "" {
	// 	addressId = dataRequest["order_address_id"].(string)
	// }

	// Ambil order_notes jika ada
	notes := ""
	if dataRequest["order_notes"] != nil && dataRequest["order_notes"].(string) != "" {
		notes = html.EscapeString(dataRequest["order_notes"].(string))
	}

	/* masukkan ke order */
	dataResponse := controller.CustomerOrderRepository.Create(r.Context(), entity.CustomerOrder{
		OrderInvNumber:    controller.CustomerOrderRepository.GenInvoice(r.Context()),
		OrderCustomerId:   dataRequest["order_customer_id"].(string),
		OrderAddressId:    addressId,
		OrderDeliveryType: dataRequest["order_delivery_type"].(string),
		OrderTotalItem:    totalItem,
		OrderSubtotal:     subtotal,
		OrderDiscount:     discount,
		OrderTotal:        total,
		OrderNotes:        notes,
	})

	/* masukkan ke table detail */
	for _, v := range orderDetails {
		dt := v.(map[string]interface{})
		controller.CustomerOrderDetailRepository.Create(r.Context(), entity.CustomerOrderDetail{
			OrderDetailParentId:        dataResponse.OrderId,
			OrderDetailProductVarianId: dt["order_detail_product_varian_id"].(string),
			OrderDetailQty:             int(dt["order_detail_qty"].(float64)),
			OrderDetailPrice:           dt["order_detail_price"].(float64),
			OrderDetailSubtotal:        dt["order_detail_subtotal"].(float64),
		})

		/* masukkan ke temporary cart jika ada temp_cart_order_id */
		if dataRequest["temp_cart_order_id"] != nil && dataRequest["temp_cart_order_id"].(string) != "" {
			controller.TempCartRepository.Create(r.Context(), entity.TempCart{
				TempCartProductVarianId: dt["order_detail_product_varian_id"].(string),
				TempCartUserId:          dataRequest["order_customer_id"].(string),
				TempCartQty:             int(dt["order_detail_qty"].(float64)),
				TempCartOrderId:         dataResponse.OrderId,
			})
		}
	}

	/* get ada admin */
	// users := controller.UserRepository.FindSpesificData(r.Context(), entity.User{
	// 	UserHasMobileAccess: 1,
	// })
	// for _, v := range users {
	// 	dt := helpers.GetFCMToken(v.UserId)
	// 	if len(dt) > 0 {
	// 		helpers.SendFCM(dt, map[string]interface{}{
	// 			"title": "Pesanan Baru",
	// 			"body":  "Ada pesanan baru senilai Rp. " + fmt.Sprint(int(dataResponse.OrderTotal)) + ", klik untuk melihat detail",
	// 			"data": map[string]interface{}{
	// 				"id":   dataResponse.OrderId,
	// 				"type": "order",
	// 			},
	// 		})
	// 	}
	// }
	/* kirim notif ke admin */

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessCreateData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["orderId"]
	dataRequest := entity.CustomerOrder{}
	helpers.ReadFromRequestBody(r, &dataRequest)
	err := controller.Validate.Struct(dataRequest)
	helpers.PanicIfError(err)
	_, err = controller.CustomerOrderRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	dataRequest.OrderId = id
	dataResponse := controller.CustomerOrderRepository.Update(r.Context(), dataRequest, "*", entity.CustomerOrder{
		OrderId: id,
	})
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessUpdateData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["orderId"]
	_, err := controller.CustomerOrderRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	controller.CustomerOrderRepository.Delete(r.Context(), entity.CustomerOrder{
		OrderId: id,
	})
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessDeleteData,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["orderId"]
	dataResponse, err := controller.CustomerOrderRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	where := entity.CustomerOrder{}

	if vars.Get("order_customer_id") != "" {
		where.OrderCustomerId = vars.Get("order_customer_id")
	}
	if vars.Get("status") != "" {
		status, _ := strconv.Atoi(vars.Get("status"))
		where.OrderStatus = status
	}

	Qlimit := vars.Get("limit")
	Qoffset := vars.Get("offset")

	if Qlimit == "" {
		Qlimit = "10"
	}

	if Qoffset == "" {
		Qoffset = "0"
	}

	limit, _ := strconv.Atoi(Qlimit)
	offset, _ := strconv.Atoi(Qoffset)

	nextOffset := limit + offset

	conf := map[string]interface{}{
		"limit":  limit,
		"offset": offset,
	}

	w.Header().Add("offset", fmt.Sprint(nextOffset))
	w.Header().Add("Access-Control-Expose-Headers", "offset")

	dataResponse := controller.CustomerOrderRepository.FindSpesificData(r.Context(), where, conf)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) GetOrderDetail(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	where := entity.ViewOrderDetail{}

	if vars.Get("order_detail_id") != "" {
		where.OrderDetailId = vars.Get("order_detail_id")
	}

	if vars.Get("order_detail_parent_id") != "" {
		where.OrderDetailParentId = vars.Get("order_detail_parent_id")
	}

	dataResponse := controller.CustomerOrderDetailRepository.FindSpesificData(r.Context(), where)
	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) OrderCanceled(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["orderId"]
	dataRequest := map[string]interface{}{}

	helpers.ReadFromRequestBody(r, &dataRequest)
	order, err := controller.CustomerOrderRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	/* ambil data dari jwt jika bukan kasir maka hanya bisa membatalkan jika statusnya belum diproses */
	data := r.Context().Value("userslogin").(jwt.MapClaims)
	/* check apakah orderan sudah di proses */
	if order.OrderStatus == 2 && data["HasAccessCashier"].(float64) == 0 {
		panic(exceptions.NewBadRequestError("Tidak dapat membatalkan pesanan karena sudah diproses, silahkan hubungi kasir untuk melakukan pembatalan"))
	}
	notes := ""
	cancelBy := order.OrderCustomerId
	if data["HasAccessCashier"].(float64) == 0 {
		notes = "Dibatalkan oleh pelanggan."
	} else {
		if dataRequest["message"] != nil {
			notes = dataRequest["message"].(string)
		} else {
			notes = "Dibatalkan oleh kasir."
		}
		cancelBy = data["UserId"].(string)
	}

	controller.CustomerOrderRepository.Update(r.Context(), entity.CustomerOrder{
		OrderStatus: 3,
		OrderFinishedDatetime: sql.NullTime{
			Valid: true,
			Time:  helpers.CreateDateTime(),
		},
		OrderFinishedBy:  cancelBy,
		OrderCancelNotes: notes,
	}, []string{"order_status", "order_finished_datetime", "order_cancel_notes", "order_finished_by"}, entity.CustomerOrder{
		OrderId: id,
	})

	/* hapus data temporary cart */
	controller.TempCartRepository.DeleteSpesificData(r.Context(), entity.TempCart{
		TempCartOrderId: id,
	})

	/* kirim notif ke kasir dan pelanggan */
	users := controller.UserRepository.FindSpesificData(r.Context(), entity.User{
		UserHasMobileAccess: 1,
	})
	for _, v := range users {
		dt := helpers.GetFCMToken(v.UserId)
		if len(dt) > 0 {
			// helpers.SendFCM(dt, map[string]interface{}{
			// 	"title": "Pesanan Dibatalkan",
			// 	"body":  "Pesanan dengan nomor order " + order.OrderInvNumber + " telah dibatalkan, klik untuk melihat detail",
			// 	"data": map[string]interface{}{
			// 		"id":   order.OrderId,
			// 		"type": "order",
			// 	},
			// })
		}
	}

	dtCust := helpers.GetFCMToken(order.OrderCustomerId)
	if len(dtCust) > 0 {
		// helpers.SendFCM(dtCust, map[string]interface{}{
		// 	"title": "Pesanan Dibatalkan",
		// 	"body":  "Pesanan dengan nomor order " + order.OrderInvNumber + " telah dibatalkan, klik untuk melihat detail",
		// 	"data": map[string]interface{}{
		// 		"id":   order.OrderId,
		// 		"type": "order",
		// 	},
		// })
	}

	webResponse := handler.WebResponse{
		Error:   false,
		Message: "Berhasil membatalkan pesanan",
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) OrderProcessed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["orderId"]
	reqData := map[string]interface{}{}

	helpers.ReadFromRequestBody(r, &reqData)

	// Validasi order_status ada
	if reqData["order_status"] == nil {
		panic(exceptions.NewBadRequestError("order_status diperlukan"))
	}

	dataRequest := entity.CustomerOrder{
		OrderStatus: int(reqData["order_status"].(float64)),
	}
	_, err := controller.CustomerOrderRepository.FindById(r.Context(), id)

	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	dataUser := r.Context().Value("userslogin").(jwt.MapClaims)
	selectField := []string{"order_status", "order_processed_datetime", "order_processed_by"}
	message := ""
	// fcmMessage := "Pesanan dengan nomor order " + order.OrderInvNumber + " telah diproses, klik untuk melihat detail"
	// fcmTitle := "Pesanan Diproses"

	if dataRequest.OrderStatus == 2 {
		// Status 2: Diproses
		dataRequest.OrderProcessedDatetime = sql.NullTime{
			Valid: true,
			Time:  helpers.CreateDateTime(),
		}
		dataRequest.OrderProcessedBy = dataUser["UserId"].(string)
		message = "Berhasil memproses pesanan"

	} else if dataRequest.OrderStatus == 1 {
		// Status 1: Selesai - Perlu data transaksi
		// fcmMessage = "Pesanan dengan nomor order " + order.OrderInvNumber + " telah selesai, klik untuk melihat detail"
		// fcmTitle = "Pesanan Selesai"

		selectField = []string{"order_status", "order_finished_datetime", "order_finished_by"}
		dataRequest.OrderFinishedDatetime = sql.NullTime{
			Valid: true,
			Time:  helpers.CreateDateTime(),
		}
		dataRequest.OrderFinishedBy = dataUser["UserId"].(string)
		message = "Berhasil menyelesaikan pesanan"

		// Validasi data transaksi ada
		// if reqData["parent"] == nil {
		// 	panic(exceptions.NewBadRequestError("Data parent diperlukan untuk menyelesaikan pesanan"))
		// }

		/* masukkan ke table transaksi */
		parent := reqData["parent"].(map[string]interface{})

		// Ambil cart_detail dari customer_order_detail berdasarkan order_id
		orderDetails := controller.CustomerOrderDetailRepository.FindSpesificData(r.Context(), entity.ViewOrderDetail{
			OrderDetailParentId: id,
		})

		// Konversi ke format cartDetails
		cartDetails := make([]map[string]interface{}, 0)
		qtyTotal := 0
		for _, detail := range orderDetails {
			cartDetails = append(cartDetails, map[string]interface{}{
				"product_varian_id": detail.OrderDetailProductVarianId,
				"product_qty":       float64(detail.OrderDetailQty),
				"product_price":     detail.OrderDetailPrice,
			})
			qtyTotal += detail.OrderDetailQty
		}

		invoice := controller.TransactionRepository.GenInvoice(r.Context())

		/* insert into parent data */
		dataTransaction := entity.Transaction{
			TransUserId:        parent["user_id"].(string),
			TransInvoice:       invoice,
			TransOrderId:       id,
			TransQtyTotal:      qtyTotal,
			TransProductTotal:  len(cartDetails),
			TransSubtotal:      parent["total_price"].(float64),
			TransDiscount:      parent["total_discount"].(float64),
			TransTotal:         parent["total_price"].(float64) - parent["total_discount"].(float64),
			TransReceivedTotal: parent["total_receive"].(float64),
			TransRefundTotal:   parent["total_receive"].(float64) - (parent["total_price"].(float64) - parent["total_discount"].(float64)),
			TransCustomerId:    parent["customer_id"].(string),
			TransStatus:        1,
		}
		trans := controller.TransactionRepository.Create(r.Context(), dataTransaction)

		/* input detail */
		for _, dt := range cartDetails {
			subtotal := dt["product_qty"].(float64) * dt["product_price"].(float64)
			transDetail := entity.TransactionDetail{
				TransDetailParentId:        trans.TransId,
				TransDetailProductVarianId: dt["product_varian_id"].(string),
				TransDetailQty:             int(dt["product_qty"].(float64)),
				TransDetailPrice:           dt["product_price"].(float64),
				TransDetailSubtotal:        subtotal,
			}
			controller.TransactionDetailRepository.Create(r.Context(), transDetail)

			/* input kartu stok */
			varian, _ := controller.VarianRepository.FindById(r.Context(), dt["product_varian_id"].(string))
			lastStock := varian.ProductVarianQtyBooth - int(dt["product_qty"].(float64))
			controller.BoothStockRepository.Create(r.Context(), entity.StockBooth{
				ProductStokProductVarianId: dt["product_varian_id"].(string),
				ProductStokFirstQty:        varian.ProductVarianQtyBooth,
				ProductStokQty:             int(dt["product_qty"].(float64)),
				ProductStokLastQty:         lastStock,
				ProductStokJenis:           "keluar",
				ProductStokPegawaiId:       parent["user_id"].(string),
			})

			/* kurangi stok booth */
			controller.VarianRepository.UpdateStock(r.Context(), dt["product_varian_id"].(string), lastStock)
		}

		/* hapus table temp cart */
		controller.TempCartRepository.DeleteSpesificData(r.Context(), entity.TempCart{
			TempCartOrderId: id,
		})

	} else {
		panic(exceptions.NewBadRequestError("Status order tidak valid. Gunakan 2 untuk diproses atau 1 untuk selesai"))
	}
	controller.CustomerOrderRepository.Update(r.Context(), dataRequest, selectField, entity.CustomerOrder{
		OrderId: id,
	})

	/* kirim notif ke pelanggan */

	// dtCust := helpers.GetFCMToken(order.OrderCustomerId)
	// if len(dtCust) > 0 {
	// 	helpers.SendFCM(dtCust, map[string]interface{}{
	// 		"title": fcmTitle,
	// 		"body":  fcmMessage,
	// 		"data": map[string]interface{}{
	// 			"id":   order.OrderId,
	// 			"type": "order",
	// 		},
	// 	})
	// }

	webResponse := handler.WebResponse{
		Error:   false,
		Message: message,
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) OrderFinished(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["orderId"]
	reqData := map[string]interface{}{}

	helpers.ReadFromRequestBody(r, &reqData)

	// Cek apakah order ada
	order, err := controller.CustomerOrderRepository.FindById(r.Context(), id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	// Validasi order harus sudah diproses (status 2)
	if order.OrderStatus != 2 {
		panic(exceptions.NewBadRequestError("Order harus dalam status diproses terlebih dahulu"))
	}

	// Validasi parent data
	// if reqData["parent"] == nil {
	// 	panic(exceptions.NewBadRequestError("Data parent diperlukan untuk menyelesaikan pesanan"))
	// }

	dataUser := r.Context().Value("userslogin").(jwt.MapClaims)
	//parent := reqData["parent"].(map[string]interface{})

	// Ambil cart_detail dari customer_order_detail berdasarkan order_id
	orderDetails := controller.CustomerOrderDetailRepository.FindSpesificData(r.Context(), entity.ViewOrderDetail{
		OrderDetailParentId: id,
	})

	// Konversi orderDetails ke format cartDetails dan hitung total
	cartDetails := make([]map[string]interface{}, 0)
	qtyTotal := 0
	subtotal := 0.0

	for _, detail := range orderDetails {
		cartDetails = append(cartDetails, map[string]interface{}{
			"product_varian_id": detail.OrderDetailProductVarianId,
			"product_qty":       float64(detail.OrderDetailQty),
			"product_price":     detail.OrderDetailPrice,
		})
		qtyTotal += detail.OrderDetailQty
		subtotal += detail.OrderDetailSubtotal
	}

	// Hitung total dari order
	discount := order.OrderDiscount
	total := subtotal - discount

	invoice := controller.TransactionRepository.GenInvoice(r.Context())

	/* insert into parent data */
	dataTransaction := entity.Transaction{
		TransUserId:        dataUser["UserId"].(string),
		TransInvoice:       invoice,
		TransOrderId:       id,
		TransQtyTotal:      qtyTotal,
		TransProductTotal:  len(cartDetails),
		TransSubtotal:      subtotal,
		TransDiscount:      discount,
		TransTotal:         total,
		TransReceivedTotal: 0,
		TransRefundTotal:   0,
		TransCustomerId:    order.OrderCustomerId,
		TransStatus:        1,
	}
	trans := controller.TransactionRepository.Create(r.Context(), dataTransaction)

	/* input detail */
	for _, dt := range cartDetails {
		subtotal := dt["product_qty"].(float64) * dt["product_price"].(float64)
		transDetail := entity.TransactionDetail{
			TransDetailParentId:        trans.TransId,
			TransDetailProductVarianId: dt["product_varian_id"].(string),
			TransDetailQty:             int(dt["product_qty"].(float64)),
			TransDetailPrice:           dt["product_price"].(float64),
			TransDetailSubtotal:        subtotal,
		}
		controller.TransactionDetailRepository.Create(r.Context(), transDetail)

		/* input kartu stok */
		varian, _ := controller.VarianRepository.FindById(r.Context(), dt["product_varian_id"].(string))
		lastStock := varian.ProductVarianQtyBooth - int(dt["product_qty"].(float64))
		controller.BoothStockRepository.Create(r.Context(), entity.StockBooth{
			ProductStokProductVarianId: dt["product_varian_id"].(string),
			ProductStokFirstQty:        varian.ProductVarianQtyBooth,
			ProductStokQty:             int(dt["product_qty"].(float64)),
			ProductStokLastQty:         lastStock,
			ProductStokJenis:           "keluar",
			ProductStokPegawaiId:       dataUser["UserId"].(string),
		})

		/* kurangi stok booth */
		controller.VarianRepository.UpdateStock(r.Context(), dt["product_varian_id"].(string), lastStock)
	}

	/* update order status menjadi finished (1) */
	controller.CustomerOrderRepository.Update(r.Context(), entity.CustomerOrder{
		OrderStatus: 1,
		OrderFinishedDatetime: sql.NullTime{
			Valid: true,
			Time:  helpers.CreateDateTime(),
		},
		OrderFinishedBy: dataUser["UserId"].(string),
	}, []string{"order_status", "order_finished_datetime", "order_finished_by"}, entity.CustomerOrder{
		OrderId: id,
	})

	/* hapus table temp cart */
	controller.TempCartRepository.DeleteSpesificData(r.Context(), entity.TempCart{
		TempCartOrderId: id,
	})

	/* kirim notif ke pelanggan */
	dtCust := helpers.GetFCMToken(order.OrderCustomerId)
	if len(dtCust) > 0 {
		// helpers.SendFCM(dtCust, map[string]interface{}{
		// 	"title": "Pesanan Selesai",
		// 	"body":  "Pesanan dengan nomor order " + order.OrderInvNumber + " telah selesai, klik untuk melihat detail",
		// 	"data": map[string]interface{}{
		// 		"id":   order.OrderId,
		// 		"type": "order",
		// 	},
		// })
	}

	webResponse := handler.WebResponse{
		Error:   false,
		Message: "Berhasil menyelesaikan pesanan",
	}
	helpers.WriteToResponseBody(w, webResponse)
}

func (controller *UseCaseImpl) GetOrderReport(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if startDate == "" && endDate == "" {
		panic(exceptions.NewBadRequestError("Parameter start_date atau end_date diperlukan"))
	}

	dataResponse := controller.CustomerOrderRepository.GetOrderReport(r.Context(), startDate, endDate)

	webResponse := handler.WebResponse{
		Error:   false,
		Message: config.LoadMessage().SuccessGetData,
		Data:    dataResponse,
	}
	helpers.WriteToResponseBody(w, webResponse)
}
