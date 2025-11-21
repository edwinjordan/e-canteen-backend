package tempcart_repository

import (
	"github.com/edwinjordan/e-canteen-backend/entity"
	"github.com/edwinjordan/e-canteen-backend/pkg/helpers"
	"github.com/edwinjordan/e-canteen-backend/repository/varian_repository"
	"gorm.io/gorm"
)

type TempCart struct {
	TempCartId              string                    `gorm:"column:temp_cart_id"`
	TempCartOrderId         string                    `gorm:"column:temp_cart_order_id"`
	TempCartProductVarianId string                    `gorm:"column:temp_cart_product_varian_id"`
	TempCartUserId          string                    `gorm:"column:temp_cart_user_id"`
	TempCartQty             int                       `gorm:"column:temp_cart_qty"`
	ProductVarian           *varian_repository.Varian `gorm:"foreignKey:TempCartProductVarianId;references:ProductVarianId"`
}

func (TempCart) TableName() string {
	return "temp_cart"
}

func (model *TempCart) BeforeCreate(tx *gorm.DB) (err error) {
	model.TempCartId = helpers.GenUUID()
	return
}

func (TempCart) FromEntity(e *entity.TempCart) *TempCart {
	return &TempCart{
		TempCartId:              e.TempCartId,
		TempCartProductVarianId: e.TempCartProductVarianId,
		TempCartOrderId:         e.TempCartOrderId,
		TempCartUserId:          e.TempCartUserId,
		TempCartQty:             e.TempCartQty,
	}
}

func (model *TempCart) ToEntity() *entity.TempCart {
	modelData := &entity.TempCart{
		TempCartId:              model.TempCartId,
		TempCartOrderId:         model.TempCartOrderId,
		TempCartProductVarianId: model.TempCartProductVarianId,
		TempCartUserId:          model.TempCartUserId,
		TempCartQty:             model.TempCartQty,
	}

	if model.ProductVarian != nil {
		modelData.ProductVarian = model.ProductVarian.ToEntity()
	}

	return modelData
}
