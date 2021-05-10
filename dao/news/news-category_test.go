package dao

import (
	"crawler/share/config"
	"testing"
)

func TestNewNewsCategoryRepository(t *testing.T) {
	repo, err := NewNewsCategoryRepository(config.MysqlDSN)
	if err != nil {
		t.Error(err)
		return
	}

	defer repo.Close()
}
