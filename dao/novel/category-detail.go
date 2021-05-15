package dao

import (
	model "crawler/model/novel"
	mgorm "crawler/share/database/gorm"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type NovelCategoryDetailRepository struct {
	db *gorm.DB
	*mgorm.Closer
}

func NewNovelCategoryDetailRepository(dsn string) (*NovelCategoryDetailRepository, error) {
	db, err := mgorm.NewMysqlGormWithTable(dsn, &model.NovelCategoryDetail{})
	if err != nil {
		return nil, fmt.Errorf("failed to create gorm: %v", err)
	}

	return &NovelCategoryDetailRepository{
		db: db,
		Closer: &mgorm.Closer{
			DB: db,
		},
	}, nil
}

func (r *NovelCategoryDetailRepository) UpsertNovelCategoryDetail(entries []model.NovelCategoryDetail) (int, error) {
	tx := r.db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(entries)
	if tx.Error != nil && tx.Error != gorm.ErrEmptySlice {
		return int(tx.RowsAffected), fmt.Errorf("failed to insert entries: %v", tx.Error)
	}

	return int(tx.RowsAffected), nil
}
