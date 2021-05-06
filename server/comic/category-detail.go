package server

import (
	model "crawler/model/comic"
	api "crawler/server/api/comic"
	"crawler/share/log"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"go.uber.org/zap"
)

type Pages struct {
	Start, End int
	AllowEmpty bool
}

// StartCrawlCategoryDetail 每页15项
func (s *ComicServer) StartCrawlCategoryDetail(ids []int, sort int, pages *Pages) error {
	if pages.Start < 0 {
		return fmt.Errorf("invalid start of pages: %d", pages.Start)
	}

	exist := s.checkCategoryDetailPage(api.ComicCategoryDetail(ids, sort, pages.End))
	if !exist && !pages.AllowEmpty {
		return fmt.Errorf("page is empty: %d", pages.End)
	}

	var wg sync.WaitGroup
	inchan := make(chan int, pages.End-pages.Start+1)
	for i := pages.Start; i <= pages.End; i++ {
		inchan <- i
	}
	close(inchan)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for page := range inchan {
				err := s.crawlCategoryDetail(api.ComicCategoryDetail(ids, sort, page))
				if err != nil {
					log.Logger.Error("", zap.Error(err))
					continue
				}
			}
		}()
	}

	wg.Wait()

	return nil
}

func (s *ComicServer) crawlCategoryDetail(api string) error {
	r, err := http.Get(api)
	if err != nil {
		return fmt.Errorf("filed to GET: %v", err)
	}
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)

	// read open bracket
	_, err = dec.Token()
	if err != nil {
		return fmt.Errorf("cannot read open bracket")
	}

	var entries []model.CategoryDetail

	for dec.More() {
		var m model.CategoryDetail

		err := dec.Decode(&m)
		if err != nil {
			log.Logger.Info("failed to decode a model: %v", zap.Error(err))
			continue
		}

		entries = append(entries, m)
	}

	if len(entries) == 0 {
		log.Logger.Info("content is empty", zap.String("url", api))
		return nil
	}

	_, err = s.CategoryDetailRepository.UpsertComicGategoryDetail(entries)
	if err != nil {
		return fmt.Errorf("failed to insert data to repository: %v", err)
	}

	_, err = dec.Token()
	if err != nil {
		log.Logger.Info("failed to read closing bracket", zap.Error(err))
	}

	return nil
}

func (s *ComicServer) checkCategoryDetailPage(api string) bool {
	r, err := http.Get(api)
	if err != nil {
		log.Logger.Error("network error", zap.Error(err))
		return false
	}
	defer r.Body.Close()

	bs := make([]byte, 2)
	_, err = r.Body.Read(bs)
	if err != nil {
		log.Logger.Error("network error", zap.Error(err))
		return false
	}

	if bs[1] == ']' {
		return false
	}

	return true
}
