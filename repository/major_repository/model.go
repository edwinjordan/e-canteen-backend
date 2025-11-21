package major_repository

import (
	"github.com/edwinjordan/e-canteen-backend/entity"
	"github.com/edwinjordan/e-canteen-backend/pkg/helpers"
	"gorm.io/gorm"
)

type Major struct {
	MajorId   string `gorm:"column:major_id"`
	MajorName string `gorm:"column:major_name"`
}

func (Major) TableName() string {
	return "majors"
}

func (Major) FromEntity(e *entity.Major) *Major {
	return &Major{
		MajorId:   e.MajorId,
		MajorName: e.MajorName,
	}
}

func (model *Major) ToEntity() *entity.MajorResponse {
	modelData := &entity.MajorResponse{
		MajorId:   model.MajorId,
		MajorName: model.MajorName,
	}

	return modelData
}

func (model *Major) BeforeCreate(tx *gorm.DB) (err error) {
	model.MajorId = helpers.GenUUID()
	return
}
