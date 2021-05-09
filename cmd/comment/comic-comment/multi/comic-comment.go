package main

import (
	comicDao "crawler/dao/comic"
	commentDao "crawler/dao/comment"
	server "crawler/server/comment"
	"crawler/share/config"
	"crawler/share/log"
	"flag"

	"go.uber.org/zap"
)

var offset = flag.Int("offset", 0, "")
var limit = flag.Int("limit", 0, "")

func main() {
	flag.Parse()

	drepo, err := comicDao.NewComicDetailRepository(config.MysqlDSN)
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
	defer drepo.Close()

	crepo, err := commentDao.NewComicCommentRepository(config.MysqlDSN)
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
	defer crepo.Close()

	s := server.CommentServer{
		ComicDetailRepository:  drepo,
		ComicCommentRepository: crepo,
	}

	err = s.StartCrawlComicComment(*offset, *limit)
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
}
