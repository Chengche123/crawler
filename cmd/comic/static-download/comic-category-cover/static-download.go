package main

import (
	dao "crawler/dao/comic"
	"crawler/share/config"
	"crawler/share/download"
	"crawler/share/log"
	"flag"

	"go.uber.org/zap"
)

func main() {
	flag.Parse()

	repo, err := dao.NewComicCategoryRepository(config.MysqlDSN)
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
	defer repo.Close()

	s := &download.Downloader{
		Repository: repo,
	}

	err = s.StartDownload()
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
}
