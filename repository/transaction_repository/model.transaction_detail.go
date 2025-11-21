package transaction_repository

import (
	"github.com/edwinjordan/e-canteen-backend/entity"
	"github.com/edwinjordan/e-canteen-backend/pkg/helpers"
	"gorm.io/gorm"
)

type TransactionDetail struct {
	TransDetailId              string  `gorm:"column:trans_detail_id"`
	TransDetailParentId        string  `gorm:"column:trans_detail_parent_id"`
	TransDetailProductVarianId string  `gorm:"column:trans_detail_product_varian_id"`
	TransDetailQty             int     `gorm:"column:trans_detail_qty"`
	TransDetailPrice           float64 `gorm:"column:trans_detail_price"`
	TransDetailSubtotal        float64 `gorm:"column:trans_detail_subtotal"`
}

func (TransactionDetail) TableName() string {
	return "transaction_details"
}

func (model *TransactionDetail) BeforeCreate(tx *gorm.DB) (err error) {
	model.TransDetailId = helpers.GenUUID()
	return
}

func (TransactionDetail) FromEntity(e *entity.TransactionDetail) *TransactionDetail {
	return &TransactionDetail{
		TransDetailId:              e.TransDetailId,
		TransDetailParentId:        e.TransDetailParentId,
		TransDetailProductVarianId: e.TransDetailProductVarianId,
		TransDetailQty:             e.TransDetailQty,
		TransDetailPrice:           e.TransDetailPrice,
		TransDetailSubtotal:        e.TransDetailSubtotal,
	}
}

func (model *TransactionDetail) ToEntity() *entity.TransactionDetail {
	modelData := &entity.TransactionDetail{
		TransDetailId:              model.TransDetailId,
		TransDetailParentId:        model.TransDetailParentId,
		TransDetailProductVarianId: model.TransDetailProductVarianId,
		TransDetailQty:             model.TransDetailQty,
		TransDetailPrice:           model.TransDetailPrice,
		TransDetailSubtotal:        model.TransDetailSubtotal,
	}

	return modelData
}
