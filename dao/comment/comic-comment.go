package dao

import (
	model "crawler/model/comment"
	mgorm "crawler/share/database/gorm"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ComicCommentRepository struct {
	db *gorm.DB
	*mgorm.Closer
}

func NewComicCommentRepository(dsn string) (*ComicCommentRepository, error) {
	db, err := mgorm.NewMysqlGormWithTable(dsn, &model.ComicComment{})
	if err != nil {
		return nil, fmt.Errorf("failed to create gorm: %v", err)
	}

	return &ComicCommentRepository{
		db: db,
		Closer: &mgorm.Closer{
			DB: db,
		},
	}, nil
}

func (r *ComicCommentRepository) UpsertComicComment(entries []model.ComicComment) (int, error) {
	tx := r.db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(entries)
	if tx.Error != nil && tx.Error != gorm.ErrEmptySlice {
		return int(tx.RowsAffected), fmt.Errorf("failed to insert entries: %v", tx.Error)
	}

	return int(tx.RowsAffected), nil
}

func (r *ComicCommentRepository) ComicCommentCount(comicid int) (int, error) {
	var count int64
	if err := r.db.Model(&model.ComicComment{}).Where("obj_id = ?", comicid).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}
