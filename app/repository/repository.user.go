package repository

import (
	"context"

	"github.com/edwinjordan/e-canteen-backend/entity"
)

type UserRepository interface {
	FindById(ctx context.Context, user entity.User, userId string) (entity.User, error)
	FindAll(ctx context.Context) []entity.User
	FindSpesificData(ctx context.Context, where entity.User) []entity.User
	CheckMaintenanceMode(ctx context.Context, where map[string]interface{}) map[string]interface{}
	UpdateFcm(ctx context.Context, userId string, fcmToken string) error
}
