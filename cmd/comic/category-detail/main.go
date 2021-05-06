package main

import (
	dao "crawler/dao/comic"
	server "crawler/server/comic"
	"crawler/share/log"

	"go.uber.org/zap"
)

func main() {
	repo, err := dao.NewCategoryDetailRepository("root:root@tcp(127.0.0.1)/comic")
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
	defer repo.Close()

	s := server.ComicServer{
		CategoryDetailRepository: repo,
	}

	err = s.StartCrawlCategoryDetail([]int{2304, 8}, 0, &server.Pages{
		Start:      301,
		End:        500,
		AllowEmpty: true,
	})
	if err != nil {
		log.Logger.Error("", zap.Error(err))
		return
	}
}
