package dao

import (
	model "crawler/model/special"
	mgorm "crawler/share/database/gorm"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ComicSpecialRepository struct {
	db *gorm.DB
	*mgorm.Closer
}

func NewComicSpecialRepository(dsn string) (*ComicSpecialRepository, error) {
	db, err := mgorm.NewMysqlGormWithTable(dsn, &model.ComicSpecial{})
	if err != nil {
		return nil, fmt.Errorf("failed to create gorm: %v", err)
	}

	return &ComicSpecialRepository{
		db: db,
		Closer: &mgorm.Closer{
			DB: db,
		},
	}, nil
}

func (r *ComicSpecialRepository) UpsertComicSpecial(entries []model.ComicSpecial) (int, error) {
	tx := r.db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(entries)
	if tx.Error != nil && tx.Error != gorm.ErrEmptySlice {
		return int(tx.RowsAffected), fmt.Errorf("failed to insert entries: %v", tx.Error)
	}

	return int(tx.RowsAffected), nil
}
