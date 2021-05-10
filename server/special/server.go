package server

import (
	model "crawler/model/special"
)

type ComicSpecialRepository interface {
	UpsertComicSpecial(entries []model.ComicSpecial) (int, error)
}

type SpecialServer struct {
	ComicSpecialRepository ComicSpecialRepository
}
