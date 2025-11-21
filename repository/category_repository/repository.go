package category_repository

import (
	"context"
	"errors"
	"html"

	"github.com/edwinjordan/e-canteen-backend/app/repository"
	"github.com/edwinjordan/e-canteen-backend/entity"
	"github.com/edwinjordan/e-canteen-backend/pkg/helpers"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.CategoryRepository {
	return &CategoryRepositoryImpl{
		DB: db,
	}
}

func (repo *CategoryRepositoryImpl) FindById(ctx context.Context, categoryId string) (entity.CategoryResponse, error) {
	categoryData := &Category{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Where("category_id = ?", categoryId).
		First(&categoryData).Error
	if err != nil {
		return *categoryData.ToEntity(), errors.New("data kategori tidak ditemukan")
	}
	return *categoryData.ToEntity(), nil
}

func (repo *CategoryRepositoryImpl) FindAll(ctx context.Context, conf map[string]interface{}) ([]entity.CategoryResponse, int) {
	category := []Category{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	var tempData []entity.CategoryResponse

	whereLike := ""
	if conf["search"].(string) != "" {
		search := html.EscapeString(conf["search"].(string))
		whereLike = "(category_name LIKE '%" + search + "%' )"
	}

	/* ambil data customer yang dipilih */
	whereNot := ""
	if whereLike == "" && conf["category"].(string) != "" {
		whereNot = conf["category"].(string)
		if conf["offset"].(int) == 0 {
			cust, err := repo.FindById(ctx, conf["category"].(string))
			if err == nil {
				tempData = append(tempData, cust)
			}
		}
	}

	// Count total rows
	var totalRows int64
	countQuery := tx.WithContext(ctx).Model(&Category{}).
		Where("category_id != ? ", whereNot).
		Where(whereLike)
	err := countQuery.Count(&totalRows).Error
	helpers.PanicIfError(err)

	err = tx.WithContext(ctx).
		Limit(conf["limit"].(int)).
		Offset(conf["offset"].(int)).
		Where(whereLike).
		Find(&category).Error
	helpers.PanicIfError(err)

	for _, v := range category {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData, int(totalRows)
}

func (repo *CategoryRepositoryImpl) Create(ctx context.Context, category entity.Category) entity.CategoryResponse {
	categoryData := &Category{}
	categoryData = categoryData.FromEntity(&category)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&categoryData).Error
	helpers.PanicIfError(err)

	return *categoryData.ToEntity()
}

func (repo *CategoryRepositoryImpl) Update(ctx context.Context, selectField interface{}, category entity.Category, categoryId string) entity.CategoryResponse {
	categoryData := &Category{}
	categoryData = categoryData.FromEntity(&category)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("category_id = ?", categoryId).Select(selectField).Updates(&categoryData).Error
	helpers.PanicIfError(err)
	return *categoryData.ToEntity()
}

func (repo *CategoryRepositoryImpl) Delete(ctx context.Context, categoryId string) {
	category := &Category{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("category_id = ?", categoryId).Delete(&category).Error
	helpers.PanicIfError(err)
}
