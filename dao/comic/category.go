package dao

import (
	model "crawler/model/comic"
	mgorm "crawler/share/database/gorm"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ComicCategoryRepository struct {
	db *gorm.DB
	*mgorm.Closer
}

func NewComicCategoryRepository(dsn string) (*ComicCategoryRepository, error) {
	db, err := mgorm.NewMysqlGormWithTable(dsn, &model.ComicCategory{})
	if err != nil {
		return nil, fmt.Errorf("failed to create gorm: %v", err)
	}

	return &ComicCategoryRepository{
		db: db,
		Closer: &mgorm.Closer{
			DB: db,
		},
	}, nil
}

func (r *ComicCategoryRepository) InsertComicGategory(entries []model.ComicCategory) (int, error) {
	var tx *gorm.DB
	if tx = r.db.Create(entries); tx.Error != nil {
		return int(tx.RowsAffected), fmt.Errorf("failed to insert entries: %v", tx.Error)
	}

	return int(tx.RowsAffected), nil
}

func (r *ComicCategoryRepository) UpsertComicGategory(entries []model.ComicCategory) (int, error) {
	tx := r.db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(entries)
	if tx.Error != nil {
		return int(tx.RowsAffected), fmt.Errorf("failed to insert entries: %v", tx.Error)
	}

	return int(tx.RowsAffected), nil
}
