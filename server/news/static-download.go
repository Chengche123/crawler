package server

import (
	"crawler/share/download"
	"crawler/share/log"
	"fmt"
	"sync"

	"go.uber.org/zap"
)

func (s *NewsServer) StartDownloadNewsCategoryDetail() error {
	models, err := s.NewsCategoryDetailRepository.FindAll()
	if err != nil {
		return fmt.Errorf("failed to find all models: %v", err)
	}

	type Pair struct {
		Rowpicurl, Cover string
	}

	inchan := make(chan *Pair, 30)
	go func() {
		for _, v := range models {
			inchan <- &Pair{
				Rowpicurl: v.Rowpicurl,
				Cover:     v.Cover,
			}
		}
		close(inchan)
	}()
	var wg sync.WaitGroup

	const concur = 32

	for i := 0; i < concur; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for pair := range inchan {
				err := download.DownloadStatic([]string{pair.Rowpicurl, pair.Cover})
				if err != nil {
					log.Logger.Info("failed to download static", zap.Error(err),
						zap.String("cover", pair.Cover), zap.String("Rowpicurl", pair.Rowpicurl))
					continue
				}
			}
		}()
	}

	wg.Wait()

	return nil
}
