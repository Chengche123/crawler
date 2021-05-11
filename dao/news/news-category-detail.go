package dao

import (
	model "crawler/model/news"
	mgorm "crawler/share/database/gorm"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type NewsCategoryDetailRepository struct {
	db *gorm.DB
	*mgorm.Closer
}

func NewNewsCategoryDetailRepository(dsn string) (*NewsCategoryDetailRepository, error) {
	db, err := mgorm.NewMysqlGormWithTable(dsn, &model.NewsCategoryDetail{})
	if err != nil {
		return nil, fmt.Errorf("failed to create gorm: %v", err)
	}

	return &NewsCategoryDetailRepository{
		db: db,
		Closer: &mgorm.Closer{
			DB: db,
		},
	}, nil
}

func (r *NewsCategoryDetailRepository) UpsertNewsCategoryDetail(entries []model.NewsCategoryDetail) (int, error) {
	tx := r.db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(entries)
	if tx.Error != nil && tx.Error != gorm.ErrEmptySlice {
		return int(tx.RowsAffected), fmt.Errorf("failed to insert entries: %v", tx.Error)
	}

	return int(tx.RowsAffected), nil
}

func (r *NewsCategoryDetailRepository) FindAll() ([]model.NewsCategoryDetail, error) {
	var res []model.NewsCategoryDetail
	if err := r.db.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
