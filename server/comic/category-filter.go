package server

import (
	model "crawler/model/comic"
	jmodel "crawler/model/comic/json"
	api "crawler/server/api/comic"
	"crawler/share/log"
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func (s *ComicServer) StartCrawlCategoryFilter() error {
	r, err := http.Get(api.ComicCategoryFilter)
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

	var entries []model.ComicCategoryFilter

	for dec.More() {
		var m jmodel.ComicCategoryFilter
		err := dec.Decode(&m)
		if err != nil {
			log.Logger.Info("failed to decode a model: %v", zap.Error(err))
			continue
		}

		entry := make([]model.ComicCategoryFilter, 0, len(m.Items))
		for _, v := range m.Items {
			entry = append(entry, model.ComicCategoryFilter{
				TagID:   v.TagID,
				Title:   m.Title,
				TagName: v.TagName,
			})
		}
		entries = append(entries, entry...)
	}

	_, err = s.CategoryFilterRepository.UpsertComicGategoryFilter(entries)
	if err != nil {
		return fmt.Errorf("failed to insert data to repository: %v", err)
	}

	_, err = dec.Token()
	if err != nil {
		log.Logger.Info("failed to read closing bracket", zap.Error(err))
	}

	return nil
}
