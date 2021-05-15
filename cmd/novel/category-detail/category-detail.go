package main

import (
	dao "crawler/dao/novel"
	server "crawler/server/novel"

	"crawler/share/config"
	"crawler/share/log"

	"go.uber.org/zap"
)

func main() {
	crepo, err := dao.NewNovelCategoryRepository(config.MysqlDSN)
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
	defer crepo.Close()

	drepo, err := dao.NewNovelCategoryDetailRepository(config.MysqlDSN)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	defer drepo.Close()

	s := server.NovelServer{
		NovelCategoryRepository:       crepo,
		NovelCategoryDetailRepository: drepo,
	}

	err = s.StartCrawlNovelCategoryDetail()
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
}
