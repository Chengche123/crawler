package dao

import (
	model "crawler/model/comic"
	"testing"
)

func TestInitTable(t *testing.T) {
	_, err := NewComicCategoryRepository()
	if err != nil {
		t.Error(err)
	}
}

func TestUpsertComicGategory(t *testing.T) {
	db, err := NewComicCategoryRepository()
	if err != nil {
		t.Error(err)
	}

	n, err := db.UpsertComicGategory([]model.ComicCategory{
		{
			TagID: 1,
			Title: "少年漫",
			Cover: "http://...",
		},
		{
			TagID: 1,
			Title: "少女漫",
			Cover: "http://...",
		},
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(n, err)
}
