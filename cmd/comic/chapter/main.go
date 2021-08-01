package main

import (
	dao "crawler/dao/comic"
	server "crawler/server/comic"
	"crawler/share/config"
	"crawler/share/log"

	"go.uber.org/zap"
)

func main() {
	rrepo, err := dao.NewComicDetailRepository(config.MysqlDSN)
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
	defer rrepo.Close()

	wrepo, err := dao.NewComicChapterRepository(config.MysqlDSN)
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
	defer wrepo.Close()

	s := server.ComicServer{
		ComicChapterRepository: wrepo,
		ComicDetailRepository:  rrepo,
	}

	err = s.StartCrawlComicChapter()
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
}
