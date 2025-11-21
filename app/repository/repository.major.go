package repository

import (
	"context"

	"github.com/edwinjordan/e-canteen-backend/entity"
)

type MajorRepository interface {
	Create(ctx context.Context, major entity.Major) entity.MajorResponse
	Update(ctx context.Context, major entity.Major, majorId string) entity.MajorResponse
	Delete(ctx context.Context, majorId string)
	FindById(ctx context.Context, majorId string) (entity.MajorResponse, error)
	FindAll(ctx context.Context, conf map[string]interface{}) ([]entity.MajorResponse, int)
	FindSpesificData(ctx context.Context, where entity.Major) []entity.MajorResponse
}
