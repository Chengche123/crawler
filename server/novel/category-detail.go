package server

import (
	model "crawler/model/novel"
	jmodel "crawler/model/novel/json"
	api "crawler/server/api/novel"
	"crawler/share/log"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"go.uber.org/zap"
)

func (s *NovelServer) StartCrawlNovelCategoryDetail() error {
	tagids, err := s.getAllCategoryTagId()
	if err != nil {
		return fmt.Errorf("cannot get tagid: %v", err)
	}

	var wg sync.WaitGroup
	for _, v := range tagids {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			err := s.startCrawlNovelCategoryDetail(id)
			if err != nil {
				log.Logger.Info("failed to startCrawlNovelCategoryDetail",
					zap.Error(err), zap.Int("category id", id))
			}

		}(v)
	}

	wg.Wait()

	return nil
}

func (s *NovelServer) startCrawlNovelCategoryDetail(id int) error {
	const pages = 20
	const concur = 8

	inchan := make(chan int, pages)
	for i := 0; i < pages; i++ {
		inchan <- i
	}
	close(inchan)

	var wg sync.WaitGroup
	for i := 0; i < concur; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for page := range inchan {

				URL := api.NovelCategoryDetail(id, 0, 0, page)
				err := s.startCrawlNovelCategoryDetailByURL(URL, id)
				if err != nil {
					log.Logger.Info("failed to startCrawlNovelCategoryDetailByURL",
						zap.Error(err), zap.String("URL", URL))
					continue
				}

			}
		}()
	}

	wg.Wait()

	return nil
}

func (s *NovelServer) startCrawlNovelCategoryDetailByURL(URL string, tagid int) error {
	r, err := http.Get(URL)
	if err != nil {
		return fmt.Errorf("failed to GET: %v", err)
	}
	defer r.Body.Close()

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("stream error: %v", err)
	}

	var jmos []jmodel.NovelCategoryDetail
	err = json.Unmarshal(bs, &jmos)
	if err != nil {
		return err
	}

	mos := make([]model.NovelCategoryDetail, len(jmos))
	for i := 0; i < len(jmos); i++ {
		mos[i].Authors = jmos[i].Authors
		mos[i].Cover = jmos[i].Cover
		mos[i].ID = jmos[i].ID
		mos[i].Name = jmos[i].Name
		mos[i].TagId = tagid
	}

	_, err = s.NovelCategoryDetailRepository.UpsertNovelCategoryDetail(mos)
	if err != nil {
		return err
	}

	return nil
}

func (s *NovelServer) getAllCategoryTagId() ([]int, error) {
	rcs, err := s.NovelCategoryRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find all category: %v", err)
	}

	res := make([]int, len(rcs))
	for i := 0; i < len(rcs); i++ {
		res[i] = rcs[i].TagID
	}

	return res, nil
}
