package dao

import (
	model "crawler/model/comic"
	mgorm "crawler/share/database/gorm"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryDetailRepository struct {
	db *gorm.DB
	*mgorm.Closer
}

func NewCategoryDetailRepository(dsn string) (*CategoryDetailRepository, error) {
	db, err := mgorm.NewMysqlGormWithTable(dsn, &model.CategoryDetail{})
	if err != nil {
		return nil, fmt.Errorf("failed to create gorm: %v", err)
	}

	return &CategoryDetailRepository{
		db: db,
		Closer: &mgorm.Closer{
			DB: db,
		},
	}, nil
}

func (r *CategoryDetailRepository) UpsertComicGategoryDetail(entries []model.CategoryDetail) (int, error) {
	tx := r.db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(entries)
	if tx.Error != nil && tx.Error != gorm.ErrEmptySlice {
		return int(tx.RowsAffected), fmt.Errorf("failed to insert entries: %v", tx.Error)
	}

	return int(tx.RowsAffected), nil
}
