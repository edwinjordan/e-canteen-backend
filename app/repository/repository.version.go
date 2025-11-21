package repository

import (
	"context"

	"github.com/edwinjordan/e-canteen-backend/entity"
)

type VersionRepository interface {
	GetVersionAdmin(ctx context.Context) entity.VersionAdmin
	GetVersionShop(ctx context.Context) entity.VersionShop
}
