package repository

import (
	"context"

	"github.com/edwinjordan/e-canteen-backend/entity"
)

type CategoryRepository interface {
	FindById(ctx context.Context, categoryId string) (entity.CategoryResponse, error)
	FindAll(ctx context.Context, conf map[string]interface{}) ([]entity.CategoryResponse, int)
	Create(ctx context.Context, category entity.Category) entity.CategoryResponse
	Update(ctx context.Context, selectField interface{}, category entity.Category, categoryId string) entity.CategoryResponse
	Delete(ctx context.Context, categoryId string)
}
