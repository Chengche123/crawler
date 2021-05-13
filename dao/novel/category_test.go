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
