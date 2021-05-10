package server

import (
	model "crawler/model/special"
)

type ComicSpecialRepository interface {
	UpsertComicSpecial(entries []model.ComicSpecial) (int, error)
	FindAll() ([]model.ComicSpecial, error)
}

type SpecialServer struct {
	ComicSpecialRepository ComicSpecialRepository
}
