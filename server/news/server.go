package server

import model "crawler/model/news"

type NewsCategoryRepository interface {
	UpsertComicComment(entries []model.NewsCategory) (int, error)
}

type NewsServer struct {
	NewsCategoryRepository NewsCategoryRepository
}
