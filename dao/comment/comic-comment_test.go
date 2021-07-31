package dao

import (
	"crawler/share/config"
	"testing"
)

func TestNewComicCommentRepository(t *testing.T) {
	repo, err := NewComicCommentRepository(config.MysqlDSN)
	if err != nil {
		t.Error(err)
		return
	}

	defer repo.Close()
}

func TestComicCommentRepository_ComicCommentCount(t *testing.T) {
	repo, err := NewComicCommentRepository(config.MysqlDSN)
	if err != nil {
		t.Error(err)
		return
	}

	defer repo.Close()

	n, err := repo.ComicCommentCount(33322)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("count:", n)
}

func TestComicCommentRepository_FindAllComicId(t *testing.T) {
	repo, err := NewComicCommentRepository(config.MysqlDSN)
	if err != nil {
		t.Error(err)
		return
	}

	defer repo.Close()

	ids, err := repo.FindAllComicId()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(ids)
}
