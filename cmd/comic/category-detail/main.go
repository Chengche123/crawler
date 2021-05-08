package main

import (
	dao "crawler/dao/comic"
	server "crawler/server/comic"
	"crawler/share/config"
	"crawler/share/log"

	"go.uber.org/zap"
)

func main() {
	repo, err := dao.NewCategoryDetailRepository(config.MysqlDSN)
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
	defer repo.Close()

	s := server.ComicServer{
		CategoryDetailRepository: repo,
	}

	err = s.StartCrawlCategoryDetail([]int{2304}, 0, &server.Pages{
		Start:      0,
		End:        1800,
		AllowEmpty: true,
	})
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
}
