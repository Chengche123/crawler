package main

import (
	dao "crawler/dao/comic"
	server "crawler/server/comic"
	"crawler/share/log"

	"go.uber.org/zap"
)

func main() {
	repo, err := dao.NewCategoryFilterRepository("root:root@tcp(127.0.0.1)/comic")
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
	defer repo.Close()

	s := server.ComicServer{
		CategoryFilterRepository: repo,
	}

	err = s.StartCrawlCategoryFilter()
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
}
