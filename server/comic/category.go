package server

import (
	model "crawler/model/comic"
	api "crawler/server/api/comic"
	"crawler/share/log"
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func (s *ComicServer) StartCrawlCategory() error {
	r, err := http.Get(api.ComicCategory)
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

	var cates []model.ComicCategory

	for dec.More() {
		var m model.ComicCategory

		err := dec.Decode(&m)
		if err != nil {
			log.Logger.Info("failed to decode a model: %v", zap.Error(err))
			continue
		}

		cates = append(cates, m)
	}

	_, err = s.CategoryRepository.UpsertComicGategory(cates)
	if err != nil {
		return fmt.Errorf("failed to insert data to repository: %v", err)
	}

	_, err = dec.Token()
	if err != nil {
		log.Logger.Info("failed to read closing bracket", zap.Error(err))
	}

	return nil
}
