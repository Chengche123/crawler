package server

import (
	dao "crawler/dao/novel"
	"crawler/share/config"
	"testing"
)

func TestNovelServer_getAllCategoryTagId(t *testing.T) {
	repo, err := dao.NewNovelCategoryRepository(config.MysqlDSN)
	if err != nil {
		t.Error(err)
		return
	}

	defer repo.Close()

	s := &NovelServer{
		NovelCategoryRepository: repo,
	}
	t.Log(s.getAllCategoryTagId())
}
