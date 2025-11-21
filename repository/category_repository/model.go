package category_repository

import (
	"github.com/edwinjordan/e-canteen-backend/entity"
	"github.com/edwinjordan/e-canteen-backend/pkg/helpers"
	"gorm.io/gorm"
)

type Category struct {
	CategoryId   string `json:"category_id"`
	CategoryName string `json:"category_name"`
}

func (Category) TableName() string {
	return "categories"
}

func (Category) FromEntity(e *entity.Category) *Category {
	return &Category{
		CategoryId:   e.CategoryId,
		CategoryName: e.CategoryName,
	}
}

func (model *Category) ToEntity() *entity.CategoryResponse {
	modelData := &entity.CategoryResponse{
		CategoryId:   model.CategoryId,
		CategoryName: model.CategoryName,
	}
	return modelData
}

func (model *Category) BeforeCreate(tx *gorm.DB) (err error) {
	model.CategoryId = helpers.GenUUID()
	return
}
