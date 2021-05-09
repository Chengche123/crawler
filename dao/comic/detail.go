package dao

import (
	model "crawler/model/comic"
	mgorm "crawler/share/database/gorm"
	"fmt"

	"gorm.io/gorm"

	"gorm.io/gorm/clause"
)

type ComicDetailRepository struct {
	db *gorm.DB
	*mgorm.Closer
}

func NewComicDetailRepository(dsn string) (*ComicDetailRepository, error) {
	db, err := mgorm.NewMysqlGormWithTable(dsn, &model.ComicDetail{})
	if err != nil {
		return nil, fmt.Errorf("failed to create gorm: %v", err)
	}

	return &ComicDetailRepository{
		db: db,
		Closer: &mgorm.Closer{
			DB: db,
		},
	}, nil
}

func (r *ComicDetailRepository) UpsertComicDetail(entries []model.ComicDetail) (int, error) {
	tx := r.db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(entries)
	if tx.Error != nil {
		return int(tx.RowsAffected), fmt.Errorf("failed to insert entries: %v", tx.Error)
	}

	return int(tx.RowsAffected), nil
}

func (r *ComicDetailRepository) FindByHotDESC(offset, limit int) ([]model.ComicDetail, error) {
	var res []model.ComicDetail

	r.db.Offset(offset).Limit(limit).Order("hotnum DESC").Find(&res)

	return res, nil
}

func (r *ComicDetailRepository) Count() (int, error) {
	var count int64
	if err := r.db.Model(&model.ComicDetail{}).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count: %v", err)
	}
	return int(count), nil
}
