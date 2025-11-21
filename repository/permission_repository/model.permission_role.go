package permission_repository

import (
	"github.com/edwinjordan/e-canteen-backend/entity"
)

type PermissionRole struct {
	PermissionId int `gorm:"column:permission_id;primaryKey"`
	RoleId       int `gorm:"column:role_id;primaryKey"`
}

func (PermissionRole) TableName() string {
	return "permission_role"
}

func (PermissionRole) FromEntity(e *entity.PermissionRole) *PermissionRole {
	return &PermissionRole{
		PermissionId: e.PermissionId,
		RoleId:       e.RoleId,
	}
}

func (model *PermissionRole) ToEntity() *entity.PermissionRole {
	return &entity.PermissionRole{
		PermissionId: model.PermissionId,
		RoleId:       model.RoleId,
	}
}
