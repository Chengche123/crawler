package server

import model "crawler/model/news"

type NewsCategoryRepository interface {
	UpsertNewsCategory(entries []model.NewsCategory) (int, error)
}

type NewsCategoryDetailRepository interface {
	UpsertNewsCategoryDetail(entries []model.NewsCategoryDetail) (int, error)
}

type NewsServer struct {
	NewsCategoryRepository       NewsCategoryRepository
	NewsCategoryDetailRepository NewsCategoryDetailRepository
}
