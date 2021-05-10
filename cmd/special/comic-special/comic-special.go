package main

import (
	dao "crawler/dao/special"
	server "crawler/server/special"

	"crawler/share/config"
	"crawler/share/log"

	"go.uber.org/zap"
)

func main() {
	repo, err := dao.NewComicSpecialRepository(config.MysqlDSN)
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
	defer repo.Close()

	s := server.SpecialServer{
		ComicSpecialRepository: repo,
	}

	err = s.StartCrawlComicSpecial(20)
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
}
