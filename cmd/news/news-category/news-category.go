package main

import (
	dao "crawler/dao/news"
	server "crawler/server/news"

	"crawler/share/config"
	"crawler/share/log"

	"go.uber.org/zap"
)

func main() {
	repo, err := dao.NewNewsCategoryRepository(config.MysqlDSN)
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
	defer repo.Close()

	s := server.NewsServer{
		NewsCategoryRepository: repo,
	}

	err = s.StartCrawlNewsCategory()
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
}
