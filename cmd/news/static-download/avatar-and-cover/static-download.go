package main

import (
	dao "crawler/dao/news"
	server "crawler/server/news"

	"crawler/share/config"
	"crawler/share/log"

	"go.uber.org/zap"
)

func main() {
	repo, err := dao.NewNewsCategoryDetailRepository(config.MysqlDSN)
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
	defer repo.Close()

	s := server.NewsServer{
		NewsCategoryDetailRepository: repo,
	}

	err = s.StartDownloadNewsCategoryDetail()
	if err != nil {
		log.Logger.Info("", zap.Error(err))
		return
	}

}
