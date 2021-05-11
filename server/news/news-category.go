package server

import (
	model "crawler/model/news"
	api "crawler/server/api/news"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (s *NewsServer) StartCrawlNewsCategory() error {
	r, err := http.Get(api.NewsCategory)
	if err != nil {
		return fmt.Errorf("failed to GET: %v", err)
	}
	defer r.Body.Close()

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("failed to read stream: %v", err)
	}

	var mos []model.NewsCategory
	err = json.Unmarshal(bs, &mos)
	if err != nil {
		return fmt.Errorf("failed to Unmarshal: %v", err)
	}

	_, err = s.NewsCategoryRepository.UpsertNewsCategory(mos)
	if err != nil {
		return fmt.Errorf("failed to UpsertComicComment: %v", err)
	}

	return nil
}
