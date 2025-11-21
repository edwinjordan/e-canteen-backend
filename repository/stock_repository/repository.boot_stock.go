package stock_repository

import (
	"context"

	"github.com/edwinjordan/e-canteen-backend/app/repository"
	"github.com/edwinjordan/e-canteen-backend/entity"
	"github.com/edwinjordan/e-canteen-backend/pkg/helpers"
	"gorm.io/gorm"
)

type BoothStockRepositoryImpl struct {
	DB *gorm.DB
}

func NewBooth(db *gorm.DB) repository.BoothStockRepository {
	return &BoothStockRepositoryImpl{
		DB: db,
	}
}

func (repo *BoothStockRepositoryImpl) Create(ctx context.Context, stockBooth entity.StockBooth) entity.StockBooth {
	stockData := &StockBooth{}
	stockData = stockData.FromEntity(&stockBooth)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&stockData).Error
	helpers.PanicIfError(err)

	return *stockData.ToEntity()
}
