package main

import (
	dao "crawler/dao/comic"
	server "crawler/server/comic"
	"crawler/share/config"
	"crawler/share/log"

	"go.uber.org/zap"
)

func main() {
	crepo, err := dao.NewCategoryDetailRepository(config.MysqlDSN)
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
	defer crepo.Close()

	drepo, err := dao.NewComicDetailRepository(config.MysqlDSN)
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
	defer drepo.Close()

	s := server.ComicServer{
		CategoryDetailRepository: crepo,
		ComicDetailRepository:    drepo,
	}

	err = s.StartCrawlComicDetail(0, 26491)
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
}
