package server

import (
	model "crawler/model/comic"
)

// 收敛repository

type CategoryFilterRepository interface {
	UpsertComicGategoryFilter(entries []model.ComicCategoryFilter) (int, error)
}

type CategoryRepository interface {
	UpsertComicGategory(entries []model.ComicCategory) (int, error)
}

type CategoryDetailRepository interface {
	UpsertComicGategoryDetail(entries []model.CategoryDetail) (int, error)
	FindByHotDESC(offset, limit int) ([]model.CategoryDetail, error)
}

type ComicDetailRepository interface {
	UpsertComicDetail(entries []model.ComicDetail) (int, error)
}

type ComicChapterRepository interface {
	UpsertComicComment(entries []model.ComicChapter) (int, error)
}

type ComicCommentRepository interface {
	FindAllComicId() ([]int, error)
}

type ComicServer struct {
	CategoryRepository       CategoryRepository
	CategoryFilterRepository CategoryFilterRepository
	CategoryDetailRepository CategoryDetailRepository
	ComicDetailRepository    ComicDetailRepository
	ComicChapterRepository   ComicChapterRepository
	ComicCommentRepository   ComicCommentRepository
}
