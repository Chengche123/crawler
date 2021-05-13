package server

import model "crawler/model/novel"

type NovelCategoryRepository interface {
	UpsertNovelCategory(entries []model.NovelCategory) (int, error)
}

type NovelServer struct {
	NovelCategoryRepository NovelCategoryRepository
}
