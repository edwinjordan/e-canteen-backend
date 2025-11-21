package repository

import (
	"context"

	"github.com/edwinjordan/e-canteen-backend/entity"
)

type VarianRepository interface {
	FindById(ctx context.Context, varianId string) (entity.Varian, error)
	FindSpesificData(ctx context.Context, where entity.Varian) []entity.Varian
	UpdateStock(ctx context.Context, id string, stok int)
	Insert(ctx context.Context, data entity.Varian) error
}
