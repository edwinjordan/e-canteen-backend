package varian_repository

import (
	"context"
	"errors"

	"github.com/edwinjordan/e-canteen-backend/app/repository"
	"github.com/edwinjordan/e-canteen-backend/entity"
	"github.com/edwinjordan/e-canteen-backend/pkg/helpers"
	"gorm.io/gorm"
)

type VarianRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.VarianRepository {

	return &VarianRepositoryImpl{
		DB: db,
	}
}

func (repo *VarianRepositoryImpl) FindById(ctx context.Context, varianId string) (entity.Varian, error) {
	data := &Varian{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("product_varian_id = ?", varianId).First(&data).Error
	if err != nil {
		return *data.ToEntity(), errors.New("data tidak ditemukan")
	}

	return *data.ToEntity(), nil
}

func (repo *VarianRepositoryImpl) UpdateStock(ctx context.Context, id string, stok int) {
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	tx.WithContext(ctx).Exec("UPDATE product_varians SET product_varian_qty_booth = ? WHERE product_varian_id = ? ", stok, id)
}

func (repo *VarianRepositoryImpl) FindSpesificData(ctx context.Context, where entity.Varian) []entity.Varian {
	varian := []Varian{}
	varianWhere := &Varian{}
	varianWhere = varianWhere.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where(varianWhere).Find(&varian).Error
	helpers.PanicIfError(err)

	var tempData []entity.Varian
	for _, v := range varian {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData

}

func (repo *VarianRepositoryImpl) Insert(ctx context.Context, data entity.Varian) error {
	varianData := &VarianInsert{
		ProductVarianId:           data.ProductVarianId,
		ProductId:                 data.ProductId,
		VarianId:                  data.VarianId,
		VarianName:                data.VarianName,
		ProductVarianPrice:        data.ProductVarianPrice,
		ProductVarianQtyBooth:     data.ProductVarianQtyBooth,
		ProductVarianQtyWarehouse: 0, // default 0 untuk warehouse
	}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Create(&varianData).Error
	if err != nil {
		return err
	}
	return nil
}
