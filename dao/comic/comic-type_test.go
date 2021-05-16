package dao

import (
	"crawler/share/config"
	"testing"
)

func TestNewComicTypeRepository(t *testing.T) {
	repo, err := NewComicTypeRepository(config.MysqlDSN)
	if err != nil {
		t.Error(err)
		return
	}

	defer repo.Close()
}
