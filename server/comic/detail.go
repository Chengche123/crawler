package server

import (
	model "crawler/model/comic"
	pmodel "crawler/model/comic/proto"
	api "crawler/server/api/comic"
	"crawler/share/crypto"
	"crawler/share/log"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"go.uber.org/zap"
)

func (s *ComicServer) StartCrawlComicDetail(offset, limit int) error {
	cdetails, err := s.CategoryDetailRepository.FindByHotDESC(offset, limit)
	if err != nil {
		return fmt.Errorf("failed to get category detail: %v", err)
	}

	var wg sync.WaitGroup
	inchan := make(chan int, len(cdetails))
	for i := 0; i < len(cdetails); i++ {
		inchan <- cdetails[i].ID
	}
	close(inchan)
	outchan := make(chan *model.ComicDetail, 15)

	for i := 0; i < 64; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for id := range inchan {
				buf, err := s.crawlComicDetail(id)
				if err != nil {
					// log.Logger.Info("failed to crawlComicDetail",
					// 	zap.String("id", strconv.Itoa(id)),
					// 	zap.String("error", err.Error()))
					continue
				}
				pb, err := pmodel.UnmashalComicDetailResponse(buf)
				if err != nil || pb.Data == nil {
					if pb.Errno == 3 {
						// 暂时无法观看,直接忽略log
						continue
					}

					log.Logger.Info("failed to UnmashalComicDetailResponse",
						zap.String("id", strconv.Itoa(id)),
						zap.String("obj", fmt.Sprintf("%+v", pb)))
					continue
				}
				mo := model.NewComicDetail(pb)

				outchan <- mo

				// _, err = s.ComicDetailRepository.UpsertComicDetail([]model.ComicDetail{*mo})
				// if err != nil {
				// 	log.Logger.Info("failed to UpsertComicDetail", zap.String("id", strconv.Itoa(id)))
				// }
			}

		}()
	}

	go func() {
		wg.Wait()
		close(outchan)
	}()

	mos := make([]model.ComicDetail, 15)
	i := 0
	for mo := range outchan {
		if i == 15 {
			_, err := s.ComicDetailRepository.UpsertComicDetail(mos)
			if err != nil {
				log.Logger.Info("failed to UpsertComicDetail", zap.String("id", "unknown"))
			}

			i = 0
		}

		mos[i] = *mo

		i++
	}

	_, err = s.ComicDetailRepository.UpsertComicDetail(mos[:i])
	if err != nil {
		log.Logger.Info("failed to UpsertComicDetail", zap.String("id", "unknown"))
	}

	return nil
}

// 返回的数据为protobuf
func (s *ComicServer) crawlComicDetail(id int) ([]byte, error) {
	r, err := http.Get(api.ComicDetail(id))
	if err != nil {
		return nil, fmt.Errorf("failed to GET: %v", err)
	}
	defer r.Body.Close()

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %v", err)
	}

	return crypto.DecryptComicDetail(string(bs))
}
