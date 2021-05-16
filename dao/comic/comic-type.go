package dao

import (
	model "crawler/model/comic"
	mgorm "crawler/share/database/gorm"
	"fmt"

	"gorm.io/gorm"

	"gorm.io/gorm/clause"
)

type ComicTypeRepository struct {
	db *gorm.DB
	*mgorm.Closer
}

func NewComicTypeRepository(dsn string) (*ComicTypeRepository, error) {
	db, err := mgorm.NewMysqlGormWithTable(dsn, &model.ComicType{})
	if err != nil {
		return nil, fmt.Errorf("failed to create gorm: %v", err)
	}

	return &ComicTypeRepository{
		db: db,
		Closer: &mgorm.Closer{
			DB: db,
		},
	}, nil
}

func (r *ComicTypeRepository) UpsertComicType(entries []model.ComicType) (int, error) {
	tx := r.db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(entries)
	if tx.Error != nil {
		return int(tx.RowsAffected), fmt.Errorf("failed to insert entries: %v", tx.Error)
	}

	return int(tx.RowsAffected), nil
}
