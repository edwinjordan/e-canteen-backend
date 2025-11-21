package permission_repository

import (
	"context"

	"github.com/edwinjordan/e-canteen-backend/app/repository"
	"github.com/edwinjordan/e-canteen-backend/entity"
	"github.com/edwinjordan/e-canteen-backend/pkg/helpers"
	"gorm.io/gorm"
)

type PermissionRoleRepositoryImpl struct {
	DB *gorm.DB
}

func NewPermissionRole(db *gorm.DB) repository.PermissionRoleRepository {
	return &PermissionRoleRepositoryImpl{
		DB: db,
	}
}

func (repo *PermissionRoleRepositoryImpl) Assign(ctx context.Context, roleId int, permissionIds []int) error {
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	for _, permissionId := range permissionIds {
		permissionRole := &PermissionRole{
			PermissionId: permissionId,
			RoleId:       roleId,
		}

		// Check if already exists
		var count int64
		tx.WithContext(ctx).
			Model(&PermissionRole{}).
			Where("permission_id = ? AND role_id = ?", permissionId, roleId).
			Count(&count)

		if count == 0 {
			err := tx.WithContext(ctx).Create(&permissionRole).Error
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (repo *PermissionRoleRepositoryImpl) Revoke(ctx context.Context, roleId int, permissionId int) error {
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).
		Where("role_id = ? AND permission_id = ?", roleId, permissionId).
		Delete(&PermissionRole{}).Error

	return err
}

func (repo *PermissionRoleRepositoryImpl) FindByRole(ctx context.Context, roleId int) []entity.PermissionRole {
	permissionRoles := []PermissionRole{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).
		Where("role_id = ?", roleId).
		Find(&permissionRoles).Error
	helpers.PanicIfError(err)

	var tempData []entity.PermissionRole
	for _, v := range permissionRoles {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData
}

func (repo *PermissionRoleRepositoryImpl) CheckPermission(ctx context.Context, roleId int, resource string, action string) bool {
	var count int64
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	tx.WithContext(ctx).
		Table("permission_role").
		Joins("INNER JOIN permissions ON permissions.permission_id = permission_role.permission_id").
		Where("permission_role.role_id = ? AND permissions.permission_resource = ? AND permissions.permission_action = ?", roleId, resource, action).
		Count(&count)

	return count > 0
}
