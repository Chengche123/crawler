package dao

import (
	"crawler/share/config"
	"testing"
)

func TestNewNovelCategoryDetailRepository(t *testing.T) {
	repo, err := NewNovelCategoryDetailRepository(config.MysqlDSN)
	if err != nil {
		t.Error(err)
		return
	}

	defer repo.Close()
}
