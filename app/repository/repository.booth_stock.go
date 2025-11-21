package repository

import (
	"context"

	"github.com/edwinjordan/e-canteen-backend/entity"
)

type BoothStockRepository interface {
	Create(ctx context.Context, stockBooth entity.StockBooth) entity.StockBooth
}
