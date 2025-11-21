package repository

import (
	"context"

	"github.com/edwinjordan/e-canteen-backend/entity"
)

type ProductRepository interface {
	FindById(ctx context.Context, productId string) (entity.Product, error)
	FindAll(ctx context.Context, where entity.Product, config map[string]interface{}) ([]entity.Product, int)
	FindSpesificData(ctx context.Context, where entity.Product) []entity.Product
	Insert(ctx context.Context, data entity.Product, varians []entity.Varian) error
	Update(ctx context.Context, data entity.Product, varians []entity.Varian) error
	Delete(ctx context.Context, productId string) error
}
