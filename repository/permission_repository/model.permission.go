package permission_repository

import (
	"github.com/edwinjordan/e-canteen-backend/entity"
)

type Permission struct {
	PermissionId          int    `gorm:"column:permission_id;primaryKey;autoIncrement"`
	PermissionName        string `gorm:"column:permission_name;unique;not null"`
	PermissionResource    string `gorm:"column:permission_resource;not null"`
	PermissionAction      string `gorm:"column:permission_action;not null"`
	PermissionDescription string `gorm:"column:permission_description"`
	PermissionStatus      string `gorm:"column:permission_status;type:enum('main_menu','submenu','action');default:action;not null"`
	PermissionParentId    *int   `gorm:"column:permission_parent_id"`
	PermissionUrutan      int    `gorm:"column:permission_urutan"`
	PermissionIcon        string `gorm:"column:permission_icon"`
	PermissionActive      int    `gorm:"column:permission_active"`
}

func (Permission) TableName() string {
	return "permissions"
}

func (Permission) FromEntity(e *entity.Permission) *Permission {
	return &Permission{
		PermissionId:          e.PermissionId,
		PermissionName:        e.PermissionName,
		PermissionResource:    e.PermissionResource,
		PermissionAction:      e.PermissionAction,
		PermissionDescription: e.PermissionDescription,
		PermissionStatus:      e.PermissionStatus,
		PermissionParentId:    e.PermissionParentId,
		PermissionUrutan:      e.PermissionUrutan,
		PermissionIcon:        e.PermissionIcon,
		PermissionActive:      e.PermissionActive,
	}
}

func (model *Permission) ToEntity() *entity.Permission {
	return &entity.Permission{
		PermissionId:          model.PermissionId,
		PermissionName:        model.PermissionName,
		PermissionResource:    model.PermissionResource,
		PermissionAction:      model.PermissionAction,
		PermissionDescription: model.PermissionDescription,
		PermissionStatus:      model.PermissionStatus,
		PermissionParentId:    model.PermissionParentId,
		PermissionUrutan:      model.PermissionUrutan,
		PermissionIcon:        model.PermissionIcon,
		PermissionActive:      model.PermissionActive,
	}
}

func (model *Permission) ToResponse() *entity.PermissionResponse {
	return &entity.PermissionResponse{
		PermissionId:          model.PermissionId,
		PermissionName:        model.PermissionName,
		PermissionResource:    model.PermissionResource,
		PermissionAction:      model.PermissionAction,
		PermissionDescription: model.PermissionDescription,
		PermissionStatus:      model.PermissionStatus,
		PermissionParentId:    model.PermissionParentId,
		PermissionUrutan:      model.PermissionUrutan,
		PermissionIcon:        model.PermissionIcon,
		PermissionActive:      model.PermissionActive,
	}
}
