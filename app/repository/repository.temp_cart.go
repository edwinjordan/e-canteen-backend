package repository

import (
	"context"

	"github.com/edwinjordan/e-canteen-backend/entity"
)

type TempCartRepository interface {
	Create(ctx context.Context, tempCart entity.TempCart)
	Update(ctx context.Context, tempCart entity.TempCart, tempCartId string)
	Delete(ctx context.Context, tempCartId string)
	DeleteSpesificData(ctx context.Context, data entity.TempCart)
	FindSpesificData(ctx context.Context, where entity.TempCart) []entity.TempCart
	FindByUserId(ctx context.Context, userId string) []entity.TempCart
	DeleteByUserId(ctx context.Context, userId string)
}
