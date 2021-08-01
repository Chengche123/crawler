package server

import (
	"crawler/share/download"
	"crawler/share/log"
	"fmt"
	"sync"

	"go.uber.org/zap"
)

const (
	step   = 15
	concur = 32
)

// StartDownloadComicCover 实际下载的资源数量可能会大于limit
func (s *ComicServer) StartDownloadComicCover(offset, limit int) error {
	if limit < step {
		return fmt.Errorf("limit must >= %d", step)
	}

	type pair struct {
		Offset, Limit int
	}
	pairs := make([]*pair, 0)

	i := offset
	end := offset + limit
	for i < end {
		pairs = append(pairs, &pair{
			Offset: i,
			Limit:  step,
		})
		i = i + step
	}

	inchan := make(chan *pair, len(pairs))
	for _, v := range pairs {
		inchan <- v
	}
	close(inchan)
	var wg sync.WaitGroup

	for i := 0; i < concur; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range inchan {
				err := s.startDownloadComicCover(v.Offset, v.Limit)

				if err != nil {
					log.Logger.Info("failed to startDownloadComicCover", zap.Error(err))
					continue
				}

			}
		}()
	}

	wg.Wait()

	return nil
}

func (s *ComicServer) startDownloadComicCover(offset, limit int) error {
	mos, err := s.CategoryDetailRepository.FindByHotDESC(offset, limit)
	if err != nil {
		return fmt.Errorf("failed to FindByHotDESC(%d,%d): %v", offset, limit, err)
	}

	apis := make([]string, len(mos))
	for i := 0; i < len(mos); i++ {
		apis[i] = mos[i].Cover
	}

	return download.DownloadStatic(apis)
}
