package dao

import (
	"crawler/share/config"
	"testing"
)

func TestCategoryDetailFindByIDASC(t *testing.T) {
	repo, err := NewCategoryDetailRepository(config.MysqlDSN)
	if err != nil {
		t.Log(err)
	}

	res, err := repo.FindByIDASC(5, 5)
	if err != nil {
		t.Log(err)
		return
	}

	for _, v := range res {
		t.Log(v.ID)
	}
}

func TestCategoryDetailFindByHotDESC(t *testing.T) {
	repo, err := NewCategoryDetailRepository(config.MysqlDSN)
	if err != nil {
		t.Log(err)
	}

	res, err := repo.FindByHotDESC(0, 5)
	if err != nil {
		t.Log(err)
		return
	}

	for _, v := range res {
		t.Log(v.ID)
	}
}
