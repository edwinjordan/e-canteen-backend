package product_repository

import (
	"context"
	"errors"
	"html"
	"time"

	"github.com/edwinjordan/e-canteen-backend/app/repository"
	"github.com/edwinjordan/e-canteen-backend/entity"
	"github.com/edwinjordan/e-canteen-backend/pkg/helpers"
	"github.com/edwinjordan/e-canteen-backend/repository/varian_repository"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.ProductRepository {

	return &ProductRepositoryImpl{
		DB: db,
	}
}

func (repo *ProductRepositoryImpl) FindById(ctx context.Context, product_id string) (entity.Product, error) {
	productData := &Product{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Preload("Varian").Where("product_id = ?", product_id).First(&productData).Error
	if err != nil {
		return *productData.ToEntity(), errors.New("data tidak ditemukan")
	}

	return *productData.ToEntity(), nil
}

func (repo *ProductRepositoryImpl) FindAll(ctx context.Context, where entity.Product, config map[string]interface{}) ([]entity.Product, int) {
	/* search */
	whereLike := ""
	if config["search"].(string) != "" {
		whereLike = "(product_name LIKE '%" + html.EscapeString(config["search"].(string)) + "%')"
	}

	product := []Product{}
	whereProduct := &Product{}
	whereProduct = whereProduct.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	// Count total rows
	var totalRows int64
	countQuery := tx.WithContext(ctx).Model(&Product{}).
		Where("product_delete_at IS NULL").
		Where(whereProduct).
		Where(whereLike)
	err := countQuery.Count(&totalRows).Error
	helpers.PanicIfError(err)

	// Get data with pagination
	err = tx.WithContext(ctx).
		Preload("Varian").
		Limit(config["limit"].(int)).
		Offset(config["offset"].(int)).
		Where("product_delete_at IS NULL").
		Where(whereProduct).
		Where(whereLike).
		Find(&product).Error
	helpers.PanicIfError(err)

	var tempData []entity.Product
	for _, v := range product {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData, int(totalRows)
}

func (repo *ProductRepositoryImpl) FindSpesificData(ctx context.Context, where entity.Product) []entity.Product {
	product := []Product{}
	whereProduct := &Product{}
	whereProduct = whereProduct.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where(whereProduct).Find(&product).Error
	helpers.PanicIfError(err)

	var tempData []entity.Product
	for _, v := range product {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData

}

func (repo *ProductRepositoryImpl) Insert(ctx context.Context, data entity.Product, varians []entity.Varian) error {
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	// Insert product
	now := time.Now()
	productData := &ProductInsert{
		ProductId:         data.ProductId,
		ProductCode:       data.ProductCode,
		ProductName:       data.ProductName,
		ProductCategoryId: data.ProductCategoryId,
		ProductDesc:       data.ProductDesc,
		ProductCreateAt:   now,
		ProductUpdateAt:   now,
		ProductPhoto:      data.ProductPhoto,
	}

	err := tx.WithContext(ctx).Create(&productData).Error
	if err != nil {
		return err
	}

	// Insert varian
	for _, varian := range varians {
		varianData := &varian_repository.VarianInsert{
			ProductVarianId:           varian.ProductVarianId,
			ProductId:                 data.ProductId,
			VarianId:                  varian.VarianId,
			VarianName:                varian.VarianName,
			ProductVarianPrice:        varian.ProductVarianPrice,
			ProductVarianQtyBooth:     varian.ProductVarianQtyBooth,
			ProductVarianQtyWarehouse: 0, // default 0 untuk warehouse
		}
		err := tx.WithContext(ctx).Create(&varianData).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *ProductRepositoryImpl) Update(ctx context.Context, data entity.Product, varians []entity.Varian) error {
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	// Check if product exists
	var existing ProductInsert
	err := tx.WithContext(ctx).Table("products").Where("product_id = ?", data.ProductId).First(&existing).Error
	if err != nil {
		return errors.New("product tidak ditemukan")
	}

	// Update product
	now := time.Now()
	updateData := map[string]interface{}{
		"product_code":        data.ProductCode,
		"product_name":        data.ProductName,
		"product_category_id": data.ProductCategoryId,
		"product_desc":        data.ProductDesc,
		"product_update_at":   now,
		"product_photo":       data.ProductPhoto,
	}

	err = tx.WithContext(ctx).Table("products").Where("product_id = ?", data.ProductId).Updates(updateData).Error
	if err != nil {
		return err
	}

	// Sync variants
	// 1. Get existing variant IDs
	var existingVariantIds []string
	err = tx.WithContext(ctx).Table("product_varians").Where("product_id = ?", data.ProductId).Pluck("product_varian_id", &existingVariantIds).Error
	if err != nil {
		return err
	}

	// 2. Collect incoming variant IDs
	incomingVariantIds := make(map[string]bool)
	for _, v := range varians {
		if v.ProductVarianId != "" {
			incomingVariantIds[v.ProductVarianId] = true
		}
	}

	// 3. Delete variants that are not in incoming list
	for _, id := range existingVariantIds {
		if !incomingVariantIds[id] {
			err = tx.WithContext(ctx).Table("product_varians").Where("product_varian_id = ?", id).Delete(&varian_repository.VarianInsert{}).Error
			if err != nil {
				return err
			}
		}
	}

	// 4. Upsert incoming variants
	for _, varian := range varians {
		if varian.ProductVarianId == "" {
			// Insert new varian
			// Generate UUID if empty
			varian.ProductVarianId = helpers.GenUUID()

			varianData := &varian_repository.VarianInsert{
				ProductVarianId:           varian.ProductVarianId,
				ProductId:                 data.ProductId,
				VarianId:                  varian.VarianId,
				VarianName:                varian.VarianName,
				ProductVarianPrice:        varian.ProductVarianPrice,
				ProductVarianQtyBooth:     varian.ProductVarianQtyBooth,
				ProductVarianQtyWarehouse: varian.ProductVarianQtyWarehouse,
			}
			err = tx.WithContext(ctx).Table("product_varians").Create(&varianData).Error
			if err != nil {
				return err
			}
		} else {
			// Check if exists to decide Update or Insert (in case ID is provided but not in DB)
			var count int64
			tx.WithContext(ctx).Table("product_varians").Where("product_varian_id = ?", varian.ProductVarianId).Count(&count)

			if count == 0 {
				// Insert with provided ID
				varianData := &varian_repository.VarianInsert{
					ProductVarianId:           varian.ProductVarianId,
					ProductId:                 data.ProductId,
					VarianId:                  varian.VarianId,
					VarianName:                varian.VarianName,
					ProductVarianPrice:        varian.ProductVarianPrice,
					ProductVarianQtyBooth:     varian.ProductVarianQtyBooth,
					ProductVarianQtyWarehouse: varian.ProductVarianQtyWarehouse,
				}
				err = tx.WithContext(ctx).Table("product_varians").Create(&varianData).Error
				if err != nil {
					return err
				}
			} else {
				// Update existing
				updateVarianData := map[string]interface{}{
					"varian_id":                    varian.VarianId,
					"varian_name":                  varian.VarianName,
					"product_varian_price":         varian.ProductVarianPrice,
					"product_varian_qty_booth":     varian.ProductVarianQtyBooth,
					"product_varian_qty_warehouse": varian.ProductVarianQtyWarehouse,
				}
				err = tx.WithContext(ctx).Table("product_varians").Where("product_varian_id = ?", varian.ProductVarianId).Updates(updateVarianData).Error
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (repo *ProductRepositoryImpl) Delete(ctx context.Context, productId string) error {
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	// Check if product exists
	var existing ProductInsert
	err := tx.WithContext(ctx).Table("products").Where("product_id = ?", productId).First(&existing).Error
	if err != nil {
		return errors.New("product tidak ditemukan")
	}

	// Soft delete - set product_delete_at
	now := time.Now()
	err = tx.WithContext(ctx).Table("products").Where("product_id = ?", productId).Update("product_delete_at", now).Error
	if err != nil {
		return err
	}

	return nil
}
