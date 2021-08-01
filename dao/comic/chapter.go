package dao

import (
	model "crawler/model/comic"
	mgorm "crawler/share/database/gorm"
	zlog "crawler/share/log"
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
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

func (r *ComicChapterRepository) FindStaticUrls() (<-chan []string, error) {
	var mos []model.ComicChapter
	if err := r.db.Find(&mos).Error; err != nil {
		return nil, err
	}

	ch := make(chan []string, 64)
	go func() {
		for _, v := range mos {
			var pageUrls []string
			err := json.Unmarshal([]byte(v.PageUrl), &pageUrls)
			if err != nil {
				zlog.Logger.Info("cannot Unmarshal page_url", zap.Error(err))
				continue
			}

			ch <- pageUrls
		}

		close(ch)
	}()

	return ch, nil
}
