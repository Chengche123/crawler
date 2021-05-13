package main

import (
	dao "crawler/dao/novel"
	server "crawler/server/novel"

	"crawler/share/config"
	"crawler/share/log"

	"go.uber.org/zap"
)

func main() {
	repo, err := dao.NewNovelCategoryRepository(config.MysqlDSN)
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
	defer repo.Close()

	s := server.NovelServer{
		NovelCategoryRepository: repo,
	}

	err = s.StartCrawlNovelCategory()
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
}
