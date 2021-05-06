package dao

import (
	model "crawler/model/comic"
	mgorm "crawler/share/database/gorm"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const dsn = "root:root@tcp(127.0.0.1)/comic"

type ComicCategoryRepository struct {
	db *gorm.DB
}

func NewComicCategoryRepository() (*ComicCategoryRepository, error) {
	db, err := mgorm.NewMysqlGormByDSN(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to init gorm: %v", err)
	}

	if !db.Migrator().HasTable(&model.ComicCategory{}) {
		err := db.Migrator().CreateTable(&model.ComicCategory{})
		if err != nil {
			return nil, fmt.Errorf("cannot create new table: %v", err)
		}
	}

	return &ComicCategoryRepository{
		db: db,
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
