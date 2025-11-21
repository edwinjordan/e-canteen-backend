package major_repository

import (
	"context"
	"errors"
	"html"

	"github.com/edwinjordan/e-canteen-backend/app/repository"
	"github.com/edwinjordan/e-canteen-backend/entity"
	"github.com/edwinjordan/e-canteen-backend/pkg/helpers"
	"gorm.io/gorm"
)

type MajorRepositoryImpl struct {
	DB *gorm.DB
}

func New(db *gorm.DB) repository.MajorRepository {
	return &MajorRepositoryImpl{
		DB: db,
	}
}

func (repo *MajorRepositoryImpl) Create(ctx context.Context, major entity.Major) entity.MajorResponse {
	majorData := &Major{}
	majorData = majorData.FromEntity(&major)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := tx.WithContext(ctx).Create(&majorData).Error
	helpers.PanicIfError(err)

	return *majorData.ToEntity()
}

func (repo *MajorRepositoryImpl) Update(ctx context.Context, major entity.Major, majorId string) entity.MajorResponse {
	majorData := &Major{}
	majorData = majorData.FromEntity(&major)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("major_id = ?", majorId).Save(&majorData).Error
	helpers.PanicIfError(err)
	return *majorData.ToEntity()
}

func (repo *MajorRepositoryImpl) Delete(ctx context.Context, majorId string) {
	major := &Major{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where("major_id = ?", majorId).Delete(&major).Error
	helpers.PanicIfError(err)
}

func (repo *MajorRepositoryImpl) FindById(ctx context.Context, majorId string) (entity.MajorResponse, error) {
	majorData := &Major{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).
		Where("major_id = ?", majorId).
		First(&majorData).Error
	if err != nil {
		return *majorData.ToEntity(), errors.New("data jurusan tidak ditemukan")
	}
	return *majorData.ToEntity(), nil
}

func (repo *MajorRepositoryImpl) FindAll(ctx context.Context, conf map[string]interface{}) ([]entity.MajorResponse, int) {
	major := []Major{}
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	var tempData []entity.MajorResponse

	whereLike := ""
	if conf["search"].(string) != "" {
		search := html.EscapeString(conf["search"].(string))
		whereLike = "(major_name LIKE '%" + search + "%' )"
	}

	/* ambil data customer yang dipilih */
	whereNot := ""
	if whereLike == "" && conf["major"].(string) != "" {
		whereNot = conf["major"].(string)
		if conf["offset"].(int) == 0 {
			cust, err := repo.FindById(ctx, conf["major"].(string))
			if err == nil {
				tempData = append(tempData, cust)
			}
		}
	}

	// Count total rows
	var totalRows int64
	countQuery := tx.WithContext(ctx).Model(&Major{}).
		Where("major_id != ? ", whereNot).
		Where(whereLike)
	err := countQuery.Count(&totalRows).Error
	helpers.PanicIfError(err)

	err = tx.WithContext(ctx).
		Limit(conf["limit"].(int)).
		Offset(conf["offset"].(int)).
		Where(whereLike).
		Find(&major).Error
	helpers.PanicIfError(err)

	for _, v := range major {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData, int(totalRows)
}

func (repo *MajorRepositoryImpl) FindSpesificData(ctx context.Context, where entity.Major) []entity.MajorResponse {
	major := []Major{}
	majorWhere := &Major{}
	majorWhere = majorWhere.FromEntity(&where)
	tx := repo.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	err := tx.WithContext(ctx).Where(majorWhere).Find(&major).Error
	helpers.PanicIfError(err)

	var tempData []entity.MajorResponse
	for _, v := range major {
		tempData = append(tempData, *v.ToEntity())
	}
	return tempData

}
