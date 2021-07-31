package main

import (
	commentDao "crawler/dao/comment"

	dao "crawler/dao/comic"
	server "crawler/server/comic"
	"crawler/share/config"
	"crawler/share/log"

	"go.uber.org/zap"
)

func main() {
	rrepo, err := commentDao.NewComicCommentRepository(config.MysqlDSN)
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
		ComicCommentRepository: rrepo,
	}

	err = s.StartCrawlComicChapter()
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
}
