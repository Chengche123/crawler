package main

import (
	dao "crawler/dao/comic"
	server "crawler/server/comic"
	"crawler/share/config"
	"crawler/share/log"

	"go.uber.org/zap"
)

func main() {
	repo, err := dao.NewComicCategoryRepository(config.MysqlDSN)
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
	defer repo.Close()

	s := server.ComicServer{
		CategoryRepository: repo,
	}

	err = s.StartCrawlCategory()
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
}
