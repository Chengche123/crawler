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

var comicid = flag.Int("id", 0, "")

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

	err = s.StartCrawlComicCommentByID(*comicid)
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
}
