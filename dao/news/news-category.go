package dao

import (
	model "crawler/model/news"
	mgorm "crawler/share/database/gorm"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type NewsCategoryRepository struct {
	db *gorm.DB
	*mgorm.Closer
}

func NewNewsCategoryRepository(dsn string) (*NewsCategoryRepository, error) {
	db, err := mgorm.NewMysqlGormWithTable(dsn, &model.NewsCategory{})
	if err != nil {
		return nil, fmt.Errorf("failed to create gorm: %v", err)
	}

	return &NewsCategoryRepository{
		db: db,
		Closer: &mgorm.Closer{
			DB: db,
		},
	}, nil
}

func (r *NewsCategoryRepository) UpsertComicComment(entries []model.NewsCategory) (int, error) {
	tx := r.db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(entries)
	if tx.Error != nil && tx.Error != gorm.ErrEmptySlice {
		return int(tx.RowsAffected), fmt.Errorf("failed to insert entries: %v", tx.Error)
	}

	return int(tx.RowsAffected), nil
}
