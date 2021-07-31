package dao

import (
	model "crawler/model/comic"
	mgorm "crawler/share/database/gorm"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ComicChapterRepository struct {
	db *gorm.DB
	*mgorm.Closer
}

func NewComicChapterRepository(dsn string) (*ComicChapterRepository, error) {
	db, err := mgorm.NewMysqlGormWithTable(dsn, &model.ComicChapter{})
	if err != nil {
		return nil, fmt.Errorf("failed to create gorm: %v", err)
	}

	return &ComicChapterRepository{
		db: db,
		Closer: &mgorm.Closer{
			DB: db,
		},
	}, nil
}

func (r *ComicChapterRepository) UpsertComicComment(entries []model.ComicChapter) (int, error) {
	tx := r.db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(entries)
	if tx.Error != nil {
		return int(tx.RowsAffected), fmt.Errorf("failed to insert entries: %v", tx.Error)
	}

	return int(tx.RowsAffected), nil
}
