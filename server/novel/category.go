package server

import (
	model "crawler/model/novel"
	api "crawler/server/api/novel"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (s *NovelServer) StartCrawlNovelCategory() error {
	r, err := http.Get(api.NovelCategory)
	if err != nil {
		return fmt.Errorf("failed to GET: %v", err)
	}
	defer r.Body.Close()

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("stream error: %v", err)
	}

	var categorys []model.NovelCategory
	err = json.Unmarshal(bs, &categorys)
	if err != nil {
		return fmt.Errorf("Unmarshal error: %v", err)
	}

	_, err = s.NovelCategoryRepository.UpsertNovelCategory(categorys)
	if err != nil {
		return err
	}

	return nil
}
