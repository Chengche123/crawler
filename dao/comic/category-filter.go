package dao

import (
	model "crawler/model/comic"
	mgorm "crawler/share/database/gorm"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryFilterRepository struct {
	db *gorm.DB
	*mgorm.Closer
}

func NewCategoryFilterRepository(dsn string) (*CategoryFilterRepository, error) {
	db, err := mgorm.NewMysqlGormWithTable(dsn, &model.ComicCategoryFilter{})
	if err != nil {
		return nil, fmt.Errorf("failed to create gorm: %v", err)
	}

	return &CategoryFilterRepository{
		db: db,
		Closer: &mgorm.Closer{
			DB: db,
		},
	}, nil
}

func (r *CategoryFilterRepository) UpsertComicGategoryFilter(entries []model.ComicCategoryFilter) (int, error) {
	tx := r.db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(entries)
	if tx.Error != nil {
		return int(tx.RowsAffected), fmt.Errorf("failed to insert entries: %v", tx.Error)
	}

	return int(tx.RowsAffected), nil
}
