package dao

import (
	"crawler/share/config"
	"testing"
)

func TestNewComicSpecialRepository(t *testing.T) {
	_, err := NewComicSpecialRepository(config.MysqlDSN)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestComicSpecialRepository_FindAll(t *testing.T) {
	repo, err := NewComicSpecialRepository(config.MysqlDSN)
	if err != nil {
		t.Error(err)
		return
	}
	defer repo.Close()

	res, err := repo.FindAll()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(len(res))
}
