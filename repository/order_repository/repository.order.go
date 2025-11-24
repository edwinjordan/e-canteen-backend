package order_repository

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/edwinjordan/e-canteen-backend/app/repository"
	"github.com/edwinjordan/e-canteen-backend/entity"
	"github.com/edwinjordan/e-canteen-backend/pkg/helpers"
	"gorm.io/gorm"
)

type CustomerOrderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrder(db *gorm.DB) repository.CustomerOrderRepository {
	return &CustomerOrderRepositoryImpl{
		DB: db,
	}
}

func (repo *CustomerOrderRepositoryImpl) Create(ctx context.Context, order entity.CustomerOrder) entity.CustomerOrder {
	orderData := &CustomerOrder{}
	orderData = orderData.FromEntity(&order)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&orderData).Error
	helpers.PanicIfError(err)

	return *orderData.ToEntity()
}

func (repo *CustomerOrderRepositoryImpl) Update(ctx context.Context, order entity.CustomerOrder, selectField interface{}, where entity.CustomerOrder) entity.CustomerOrder {
	orderData := &CustomerOrder{}
	orderData = orderData.FromEntity(&order)

	whereData := &CustomerOrder{}
	whereData = whereData.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where(whereData).Select(selectField).Save(&orderData).Error
	helpers.PanicIfError(err)
	return *orderData.ToEntity()
}

func (repo *CustomerOrderRepositoryImpl) Delete(ctx context.Context, where entity.CustomerOrder) {
	order := &CustomerOrder{}
	whereData := order.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where(whereData).Delete(&order).Error
	helpers.PanicIfError(err)
}

func (repo *CustomerOrderRepositoryImpl) FindById(ctx context.Context, orderId string) (entity.CustomerOrder, error) {
	orderData := &CustomerOrder{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	// Query order dengan join ke view untuk mendapatkan product dan varian name
	err := tx.WithContext(ctx).
		Where("order_id = ?", orderId).
		Preload("Address").
		Preload("Customer").
		First(&orderData).Error
	if err != nil {
		return *orderData.ToEntity(), errors.New("data order tidak ditemukan")
	}

	// Load order detail dari view yang sudah include product_name dan varian_name
	var orderDetails []ViewOrderDetail
	tx.WithContext(ctx).
		Where("order_detail_parent_id = ?", orderId).
		Find(&orderDetails)

	// Convert ViewOrderDetail ke entity
	var detailEntities []entity.ViewOrderDetail
	for _, detail := range orderDetails {
		detailEntities = append(detailEntities, *detail.ToEntity())
	}

	// Map ke entity order
	result := orderData.ToEntity()
	result.OrderDetail = &detailEntities

	return *result, nil
}

func (repo *CustomerOrderRepositoryImpl) FindAll(ctx context.Context, customerId string) []entity.CustomerOrder {
	order := []CustomerOrder{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("order_customer_id = ?", customerId).Find(&order).Error
	helpers.PanicIfError(err)

	var tempData []entity.CustomerOrder
	for _, v := range order {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData
}

func (repo *CustomerOrderRepositoryImpl) FindSpesificData(ctx context.Context, where entity.CustomerOrder, config map[string]interface{}) []entity.CustomerOrder {
	order := []CustomerOrder{}
	orderWhere := &CustomerOrder{}
	orderWhere = orderWhere.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	orderString := "order_create_at DESC"

	if orderWhere.OrderStatus == 1 || orderWhere.OrderStatus == 3 {
		orderString = "order_finished_datetime DESC"
	} else if orderWhere.OrderStatus == 2 {
		orderString = "order_processed_datetime DESC"
	}

	query := tx.WithContext(ctx).
		Limit(config["limit"].(int)).
		Offset(config["offset"].(int)).
		Order(orderString).
		Where(orderWhere).
		Preload("Address").
		Preload("Customer")

	if val, ok := config["start_date"].(string); ok && val != "" {
		if val2, ok2 := config["end_date"].(string); ok2 && val2 != "" {
			query = query.Where("order_create_at BETWEEN ? AND ?", val, val2)
		} else {
			query = query.Where("order_create_at >= ?", val)
		}
	} else if val, ok := config["end_date"].(string); ok && val != "" {
		query = query.Where("order_create_at <= ?", val)
	}

	err := query.Find(&order).Error
	helpers.PanicIfError(err)

	var tempData []entity.CustomerOrder
	for _, v := range order {
		orderEntity := v.ToEntity()

		// Load order detail dari view
		var orderDetails []ViewOrderDetail
		tx.WithContext(ctx).
			Where("order_detail_parent_id = ?", v.OrderId).
			Find(&orderDetails)

		// Convert ViewOrderDetail ke entity
		var detailEntities []entity.ViewOrderDetail
		for _, detail := range orderDetails {
			detailEntities = append(detailEntities, *detail.ToEntity())
		}
		orderEntity.OrderDetail = &detailEntities

		tempData = append(tempData, *orderEntity)
	}
	return tempData

}

func (repo *CustomerOrderRepositoryImpl) GenInvoice(ctx context.Context) string {
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	invoice := map[string]interface{}{}

	month := fmt.Sprint(int(time.Now().Month()))
	if len(month) == 1 {
		month = "0" + month
	}
	day := fmt.Sprint(time.Now().Day())
	if len(day) == 1 {
		day = "0" + day
	}
	year := fmt.Sprint(int(time.Now().Year()) % 1e2)

	date := day + month + year

	tx.WithContext(ctx).Table("customer_orders").Select("IFNULL(order_inv_number,'') order_inv_number").Where("order_inv_number LIKE ?", "%"+date+"%").Order("order_inv_number DESC").Find(invoice)
	inv := ""
	if invoice["order_inv_number"] == nil {
		inv = "ORN-" + date + "-000"
	} else {
		inv = invoice["order_inv_number"].(string)
	}
	sort := inv[len(inv)-3:]
	newInv := inv[:len(inv)-3]
	str, _ := strconv.Atoi(sort)
	str += 1
	if str < 10 {
		sort = "00" + fmt.Sprint(str)
	} else if str < 100 {
		sort = "0" + fmt.Sprint(str)
	} else {
		sort = fmt.Sprint(str)
	}

	return newInv + sort
}

func (repo *CustomerOrderRepositoryImpl) GetOrderReport(ctx context.Context, startDate, endDate, status string) []entity.CustomerOrder {
	order := []CustomerOrder{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	query := tx.WithContext(ctx).
		Preload("Customer").
		Preload("Address").
		Preload("OrderDetail")

	if startDate != "" && endDate != "" {
		query = query.Where("DATE(order_create_at) BETWEEN ? AND ?", startDate, endDate)
	} else if startDate != "" {
		query = query.Where("DATE(order_create_at) >= ?", startDate)
	} else if endDate != "" {
		query = query.Where("DATE(order_create_at) <= ?", endDate)
	}

	// Add status filter if provided
	if status != "" {
		query = query.Where("order_status = ?", status)
	}

	err := query.Order("order_create_at DESC").Find(&order).Error
	helpers.PanicIfError(err)

	var tempData []entity.CustomerOrder
	for _, v := range order {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData
}
