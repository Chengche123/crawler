package dao

import (
	"crawler/share/config"
	"testing"
)

func TestNewNewsCategoryDetailRepository(t *testing.T) {
	repo, err := NewNewsCategoryDetailRepository(config.MysqlDSN)
	if err != nil {
		t.Error(err)
		return
	}

	defer repo.Close()
}
