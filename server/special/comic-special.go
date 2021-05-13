package server

import (
	model "crawler/model/special"
	jmodel "crawler/model/special/json"
	api "crawler/server/api/special"
	"crawler/share/log"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"go.uber.org/zap"
)

func (s *SpecialServer) StartCrawlComicSpecial(pages int) error {
	inchan := make(chan int, pages)
	for i := 0; i < pages; i++ {
		inchan <- i
	}
	close(inchan)
	var wg sync.WaitGroup

	const concur = 32
	for i := 0; i < concur; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for page := range inchan {
				URL := api.ComicSpecial(page)
				err := s.startCrawlComicSpecialByURL(URL)
				if err != nil {
					log.Logger.Info("failed to startCrawlComicSpecialByURL", zap.Error(err), zap.String("url", URL))
				}
			}
		}()
	}

	wg.Wait()

	return nil
}

func (s SpecialServer) startCrawlComicSpecialByURL(URL string) error {
	r, err := http.Get(URL)
	if err != nil {
		return fmt.Errorf("failed to GET: %v", err)
	}
	defer r.Body.Close()

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("stream occur error: %v", err)
	}

	var jmo jmodel.ComicSpecial

	err = json.Unmarshal(bs, &jmo)
	if err != nil {
		return fmt.Errorf("Unmarshal error: %v", err)
	}

	specials := make([]model.ComicSpecial, 0, len(jmo.Data))

	for _, special := range jmo.Data {
		specials = append(specials, model.ComicSpecial{
			ID:         special.ID,
			Title:      special.Title,
			ShortTitle: special.ShortTitle,
			CreateTime: special.CreateTime,
			SmallCover: special.SmallCover,
			PageType:   special.PageType,
			Sort:       special.Sort,
			PageURL:    special.PageURL,
		})
	}

	_, err = s.ComicSpecialRepository.UpsertComicSpecial(specials)
	if err != nil {
		return fmt.Errorf("dao error: %v", err)
	}

	return nil
}
