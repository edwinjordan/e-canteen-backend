package repository

import (
	"context"

	"github.com/edwinjordan/e-canteen-backend/entity"
)

type PermissionRepository interface {
	Create(ctx context.Context, permission entity.Permission) entity.Permission
	Update(ctx context.Context, permission entity.Permission) entity.Permission
	Delete(ctx context.Context, permissionId int) error
	FindById(ctx context.Context, permissionId int) (*entity.Permission, error)
	FindByName(ctx context.Context, name string) entity.Permission
	FindAll(ctx context.Context) []entity.Permission
	FindByRole(ctx context.Context, roleId int) []entity.Permission
	FindByRoleAsTree(ctx context.Context, roleId int) []*entity.PermissionTreeNode
}

type PermissionRoleRepository interface {
	Assign(ctx context.Context, roleId int, permissionIds []int) error
	Revoke(ctx context.Context, roleId int, permissionId int) error
	FindByRole(ctx context.Context, roleId int) []entity.PermissionRole
	CheckPermission(ctx context.Context, roleId int, resource string, action string) bool
}
