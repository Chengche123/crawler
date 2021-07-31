package server

import (
	comicModel "crawler/model/comic"
	commentModel "crawler/model/comment"
)

// 收敛repository

type ComicDetailRepository interface {
	FindByHotDESC(offset, limit int) ([]comicModel.ComicDetail, error)
	Count() (int, error)
}

type ComicCommentRepository interface {
	UpsertComicComment(entries []commentModel.ComicComment) (int, error)
	ComicCommentCount(comicid int) (int, error)
	FindAllComicId() ([]int, error)
}

type CommentServer struct {
	ComicDetailRepository  ComicDetailRepository
	ComicCommentRepository ComicCommentRepository
}
