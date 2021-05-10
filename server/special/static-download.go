package server

import (
	"crawler/share/download"
	"crawler/share/log"
	"fmt"
	"sync"

	"go.uber.org/zap"
)

func (s *SpecialServer) StartDownloadComicSpecial() error {
	// 因专题封面较少,直接返回所有model
	models, err := s.ComicSpecialRepository.FindAll()
	if err != nil {
		return fmt.Errorf("failed to find all models: %v", err)
	}

	inchan := make(chan string, len(models))
	for i := 0; i < len(models); i++ {
		inchan <- models[i].SmallCover
	}
	close(inchan)
	var wg sync.WaitGroup

	const concur = 32

	for i := 0; i < concur; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for cover := range inchan {
				err := download.DownloadStatic([]string{cover})
				if err != nil {
					log.Logger.Info("failed to download cover", zap.Error(err), zap.String("cover", cover))
					continue
				}
			}
		}()
	}

	wg.Wait()

	return nil
}
