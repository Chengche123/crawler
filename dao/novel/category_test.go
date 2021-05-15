package dao

import (
	"crawler/share/config"
	"testing"
)

func TestNewNovelCategoryRepository(t *testing.T) {
	repo, err := NewNovelCategoryRepository(config.MysqlDSN)
	if err != nil {
		t.Error(err)
		return
	}

	defer repo.Close()
}

func TestNovelCategoryRepository_FindAll(t *testing.T) {
	repo, err := NewNovelCategoryRepository(config.MysqlDSN)
	if err != nil {
		t.Error(err)
		return
	}

	defer repo.Close()

	rcs, err := repo.FindAll()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(rcs)
}
