package main

import (
	dao "crawler/dao/comic"
	server "crawler/server/comic"
	"crawler/share/config"
	"crawler/share/log"
	"flag"

	"go.uber.org/zap"
)

var offset = flag.Int("offset", 0, "")
var limit = flag.Int("limit", 0, "")

func main() {
	flag.Parse()

	repo, err := dao.NewCategoryDetailRepository(config.MysqlDSN)
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
	defer repo.Close()

	s := server.ComicServer{
		CategoryDetailRepository: repo,
	}

	err = s.StartDownloadComicCover(*offset, *limit)
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
}
