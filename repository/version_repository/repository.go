package version_repository

import (
	"context"

	"github.com/edwinjordan/e-canteen-backend/app/repository"
	"github.com/edwinjordan/e-canteen-backend/entity"
	"github.com/edwinjordan/e-canteen-backend/pkg/helpers"
	"gorm.io/gorm"
)

type VersionRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.VersionRepository {
	return &VersionRepositoryImpl{
		DB: db,
	}
}

func (repo *VersionRepositoryImpl) GetVersionAdmin(ctx context.Context) entity.VersionAdmin {
	versionData := &VersionAdmin{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Order("version_code DESC").
		First(&versionData).Error
	helpers.PanicIfError(err)
	return *versionData.ToEntity()
}

func (repo *VersionRepositoryImpl) GetVersionShop(ctx context.Context) entity.VersionShop {
	versionData := &VersionShop{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Order("version_code DESC").
		First(&versionData).Error
	helpers.PanicIfError(err)
	return *versionData.ToEntity()
}
