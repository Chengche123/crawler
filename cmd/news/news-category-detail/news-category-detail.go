package main

import (
	dao "crawler/dao/news"
	server "crawler/server/news"
	"sync"

	"crawler/share/config"
	"crawler/share/log"

	"go.uber.org/zap"
)

func main() {
	repo, err := dao.NewNewsCategoryDetailRepository(config.MysqlDSN)
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
	defer repo.Close()

	s := server.NewsServer{
		NewsCategoryDetailRepository: repo,
	}

	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			err := s.StartCrawlNewsCategoryDetail(id, 10)
			if err != nil {
				log.Logger.Info("", zap.Error(err))
			}
		}(i)
	}

	wg.Wait()
}
