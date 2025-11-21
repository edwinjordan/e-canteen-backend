package router

import (
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/edwinjordan/e-canteen-backend/app/usecase/usecase_order"
	"github.com/edwinjordan/e-canteen-backend/repository/order_repository"
	"github.com/edwinjordan/e-canteen-backend/repository/stock_repository"
	"github.com/edwinjordan/e-canteen-backend/repository/tempcart_repository"
	"github.com/edwinjordan/e-canteen-backend/repository/transaction_repository"
	"github.com/edwinjordan/e-canteen-backend/repository/user_repository"
	"github.com/edwinjordan/e-canteen-backend/repository/varian_repository"
	"gorm.io/gorm"
)

func OrderRouter(db *gorm.DB, validate *validator.Validate, router *mux.Router) {
	orderRepository := order_repository.NewOrder(db)
	orderDetailRepository := order_repository.NewOrderDetail(db)
	varianRepository := varian_repository.New(db)
	tempCartRepo := tempcart_repository.New(db)
	transRepo := transaction_repository.NewTrans(db)
	transDetailRepo := transaction_repository.NewTransDetail(db)
	stockBoothRepo := stock_repository.NewBooth(db)
	userRepository := user_repository.New(db)
	orderController := usecase_order.NewUseCase(orderRepository, orderDetailRepository, varianRepository, tempCartRepo, transRepo, transDetailRepo, stockBoothRepo, userRepository, validate)
	router.HandleFunc("/api/order", orderController.Create).Methods("POST")
	router.HandleFunc("/api/order", orderController.FindAll).Methods("GET")
	router.HandleFunc("/api/order/{orderId}", orderController.FindById).Methods("GET")
	router.HandleFunc("/api/order_detail", orderController.GetOrderDetail).Methods("GET")
	router.HandleFunc("/api/order_report", orderController.GetOrderReport).Methods("GET")
	router.HandleFunc("/api/order_canceled/{orderId}", orderController.OrderCanceled).Methods("PUT")
	router.HandleFunc("/api/kasir/order_processed/{orderId}", orderController.OrderProcessed).Methods("PUT")
	router.HandleFunc("/api/kasir/order_finished/{orderId}", orderController.OrderFinished).Methods("PUT")
}
