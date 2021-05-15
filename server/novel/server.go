package server

import model "crawler/model/novel"

type NovelCategoryRepository interface {
	UpsertNovelCategory(entries []model.NovelCategory) (int, error)
	FindAll() ([]model.NovelCategory, error)
}

type NovelCategoryDetailRepository interface {
	UpsertNovelCategoryDetail(entries []model.NovelCategoryDetail) (int, error)
}

type NovelServer struct {
	NovelCategoryRepository       NovelCategoryRepository
	NovelCategoryDetailRepository NovelCategoryDetailRepository
}
