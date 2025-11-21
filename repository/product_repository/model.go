package product_repository

import (
	"time"

	"github.com/edwinjordan/e-canteen-backend/entity"
	"github.com/edwinjordan/e-canteen-backend/repository/varian_repository"
)

type Product struct {
	ProductId         string                      `gorm:"column:product_id"`
	ProductCode       string                      `gorm:"column:product_code"`
	ProductName       string                      `gorm:"column:product_name"`
	ProductCategoryId string                      `gorm:"column:product_category_id"`
	ProductDesc       string                      `gorm:"column:product_desc"`
	CategoryName      string                      `gorm:"column:category_name"`
	ProductCreateAt   time.Time                   `gorm:"column:product_create_at"`
	ProductUpdateAt   time.Time                   `gorm:"column:product_update_at"`
	ProductDeleteAt   time.Time                   `gorm:"column:product_delete_at"`
	ProductPhoto      string                      `gorm:"column:product_photo"`
	Varian            *[]varian_repository.Varian `gorm:"foreignKey:ProductId;references:ProductId"`
}

func (Product) TableName() string {
	return "v_ms_product"
}

func (Product) FromEntity(e *entity.Product) *Product {
	return &Product{
		ProductId:         e.ProductId,
		ProductCode:       e.ProductCode,
		ProductName:       e.ProductName,
		ProductCategoryId: e.ProductCategoryId,
		ProductDesc:       e.ProductDesc,
		CategoryName:      e.CategoryName,
		ProductCreateAt:   e.ProductCreateAt,
		ProductUpdateAt:   e.ProductUpdateAt,
		ProductDeleteAt:   e.ProductDeleteAt,
		ProductPhoto:      e.ProductPhoto,
	}
}

func (model *Product) ToEntity() *entity.Product {
	modelData := &entity.Product{
		ProductId:         model.ProductId,
		ProductCode:       model.ProductCode,
		ProductName:       model.ProductName,
		ProductCategoryId: model.ProductCategoryId,
		ProductDesc:       model.ProductDesc,
		CategoryName:      model.CategoryName,
		ProductCreateAt:   model.ProductCreateAt,
		ProductUpdateAt:   model.ProductUpdateAt,
		ProductDeleteAt:   model.ProductDeleteAt,
		ProductPhoto:      model.ProductPhoto,
	}

	if model.Varian != nil {
		var tempMenu []entity.Varian
		for _, v := range *model.Varian {
			tempMenu = append(tempMenu, *v.ToEntity())
		}
		modelData.Varian = &tempMenu
	}

	return modelData
}

// ProductInsert model for inserting into ms_product table
type ProductInsert struct {
	ProductId         string    `gorm:"column:product_id;primaryKey"`
	ProductCode       string    `gorm:"column:product_code"`
	ProductName       string    `gorm:"column:product_name"`
	ProductCategoryId string    `gorm:"column:product_category_id"`
	ProductDesc       string    `gorm:"column:product_desc"`
	ProductCreateAt   time.Time `gorm:"column:product_create_at"`
	ProductUpdateAt   time.Time `gorm:"column:product_update_at"`
	ProductPhoto      string    `gorm:"column:product_photo"`
}

func (ProductInsert) TableName() string {
	return "products"
}

// func (model *ProductInsert) BeforeCreate(tx *gorm.DB) (err error) {
// 	model.ProductId = helpers.GenUUID()
// 	return
// }
