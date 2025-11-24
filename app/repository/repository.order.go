package repository

import (
	"context"

	"github.com/edwinjordan/e-canteen-backend/entity"
)

type CustomerOrderRepository interface {
	Create(ctx context.Context, order entity.CustomerOrder) entity.CustomerOrder
	Update(ctx context.Context, order entity.CustomerOrder, selectField interface{}, where entity.CustomerOrder) entity.CustomerOrder
	Delete(ctx context.Context, where entity.CustomerOrder)
	FindById(ctx context.Context, orderId string) (entity.CustomerOrder, error)
	FindAll(ctx context.Context, customerId string) []entity.CustomerOrder
	FindSpesificData(ctx context.Context, where entity.CustomerOrder, conf map[string]interface{}) []entity.CustomerOrder
	GenInvoice(ctx context.Context) string
	GetOrderReport(ctx context.Context, startDate, endDate, status string) []entity.CustomerOrder
}
