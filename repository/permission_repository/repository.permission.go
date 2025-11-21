package permission_repository

import (
	"context"
	"errors"

	"github.com/edwinjordan/e-canteen-backend/app/repository"
	"github.com/edwinjordan/e-canteen-backend/entity"
	"github.com/edwinjordan/e-canteen-backend/pkg/helpers"
	"gorm.io/gorm"
)

type PermissionRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.PermissionRepository {
	return &PermissionRepositoryImpl{
		DB: db,
	}
}

func (repo *PermissionRepositoryImpl) Create(ctx context.Context, permission entity.Permission) entity.Permission {
	permissionData := &Permission{}
	permissionData = permissionData.FromEntity(&permission)

	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&permissionData).Error
	helpers.PanicIfError(err)

	return *permissionData.ToEntity()
}

func (repo *PermissionRepositoryImpl) Update(ctx context.Context, permission entity.Permission) entity.Permission {
	permissionData := &Permission{}
	permissionData = permissionData.FromEntity(&permission)

	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).
		Model(&Permission{}).
		Where("permission_id = ?", permission.PermissionId).
		Updates(permissionData).Error
	helpers.PanicIfError(err)

	return *permissionData.ToEntity()
}

func (repo *PermissionRepositoryImpl) Delete(ctx context.Context, permissionId int) error {
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).
		Where("permission_id = ?", permissionId).
		Delete(&Permission{}).Error

	return err
}

func (repo *PermissionRepositoryImpl) FindById(ctx context.Context, permissionId int) (*entity.Permission, error) {
	permissionData := &Permission{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).
		Where("permission_id = ?", permissionId).
		First(&permissionData).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("permission not found")
		}
		return nil, err
	}

	return permissionData.ToEntity(), nil
}

func (repo *PermissionRepositoryImpl) FindByName(ctx context.Context, name string) entity.Permission {
	permissionData := &Permission{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).
		Where("permission_name = ?", name).
		First(&permissionData).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Permission{}
		}
		helpers.PanicIfError(err)
	}

	return *permissionData.ToEntity()
}

func (repo *PermissionRepositoryImpl) FindAll(ctx context.Context) []entity.Permission {
	permissions := []Permission{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Find(&permissions).Error
	helpers.PanicIfError(err)

	var tempData []entity.Permission
	for _, v := range permissions {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData
}

func (repo *PermissionRepositoryImpl) FindByRole(ctx context.Context, roleId int) []entity.Permission {
	permissions := []Permission{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).
		Joins("INNER JOIN permission_role ON permission_role.permission_id = permissions.permission_id").
		Where("permission_role.role_id = ?", roleId).
		Find(&permissions).Error
	helpers.PanicIfError(err)

	var tempData []entity.Permission
	for _, v := range permissions {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData
}

// FindByRoleAsTree returns permissions for a role in hierarchical tree structure
func (repo *PermissionRepositoryImpl) FindByRoleAsTree(ctx context.Context, roleId int) []*entity.PermissionTreeNode {
	permissions := []Permission{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).
		Joins("INNER JOIN permission_role ON permission_role.permission_id = permissions.permission_id").
		Where("permission_role.role_id = ?", roleId).
		Order("permission_parent_id ASC, permission_id ASC").
		Find(&permissions).Error
	helpers.PanicIfError(err)

	// Convert to tree nodes
	nodeMap := make(map[int]*entity.PermissionTreeNode)
	var rootNodes []*entity.PermissionTreeNode

	// First pass: create all nodes
	for _, p := range permissions {
		node := &entity.PermissionTreeNode{
			PermissionId:          p.PermissionId,
			PermissionName:        p.PermissionName,
			PermissionResource:    p.PermissionResource,
			PermissionAction:      p.PermissionAction,
			PermissionDescription: p.PermissionDescription,
			PermissionStatus:      p.PermissionStatus,
			PermissionParentId:    p.PermissionParentId,
			Children:              []*entity.PermissionTreeNode{},
		}
		nodeMap[p.PermissionId] = node
	}

	// Second pass: build tree structure
	for _, node := range nodeMap {
		if node.PermissionParentId == nil {
			// Root node
			rootNodes = append(rootNodes, node)
		} else {
			// Child node - add to parent
			if parent, exists := nodeMap[*node.PermissionParentId]; exists {
				parent.Children = append(parent.Children, node)
			} else {
				// Parent not in the list, treat as root
				rootNodes = append(rootNodes, node)
			}
		}
	}

	return rootNodes
}

