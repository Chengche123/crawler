package server

import (
	model "crawler/model/news"
	pmodel "crawler/model/news/proto"
	api "crawler/server/api/news"
	"crawler/share/crypto"
	"crawler/share/log"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

func (s *NewsServer) StartCrawlComicNewsCategoryDetail() error {
	// 100页左右
	return s.StartCrawlNewsCategoryDetail(2, 100)
}

func (s *NewsServer) StartCrawlNewsCategoryDetail(id, pages int) error {
	inchan := make(chan string, 15)
	go func() {
		// page从1开始
		for i := 1; i <= pages; i++ {
			inchan <- api.NewsCategoryDetail(id, i)
		}
		close(inchan)
	}()
	var wg sync.WaitGroup

	const concur = 32
	for i := 0; i < concur; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for URL := range inchan {
				err := s.startCrawlNewsCategoryDetailByURL(id, URL)
				if err != nil {
					log.Logger.Info("failed to startCrawlNewsCategoryDetailByURL",
						zap.Error(err), zap.String("URL", URL))
					continue
				}

			}
		}()
	}

	wg.Wait()

	return nil
}

func (s *NewsServer) startCrawlNewsCategoryDetailByURL(tagid int, URL string) error {
	r, err := http.Get(URL)
	if err != nil {
		return fmt.Errorf("faeild to GET: %v", err)
	}
	defer r.Body.Close()

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("stream read error: %v", err)
	}

	pbbuf, err := crypto.Decrypt(string(bs))
	if err != nil {
		return fmt.Errorf("failed to Decrypt: %v", err)
	}

	var detailpb pmodel.NewsListResponse

	err = proto.Unmarshal(pbbuf, &detailpb)
	if err != nil {
		return fmt.Errorf("failed to Unmarshal pb: %v", err)
	}

	ds := detailpb.Data

	details := make([]model.NewsCategoryDetail, len(ds))
	for i := 0; i < len(ds); i++ {
		details[i].Articleid = int(ds[i].ArticleId)
		details[i].Authorid = int(ds[i].AuthorId)
		details[i].Authoruid = int(ds[i].AuthorUid)
		details[i].Colpicurl = ds[i].ColPicUrl
		details[i].Commentamount = int(ds[i].CommentAmount)
		details[i].Cover = ds[i].Cover
		details[i].Createtime = int(ds[i].CreateTime)
		details[i].Fromname = ds[i].FromName
		details[i].Intro = ds[i].Intro
		details[i].Moodamount = int(ds[i].MoodAmount)
		details[i].Nickname = ds[i].Nickname
		details[i].Pageurl = ds[i].PageUrl
		details[i].Rowpicurl = ds[i].RowPicUrl
		details[i].Status = int(ds[i].Status)
		details[i].Title = ds[i].Title

		details[i].TagId = tagid
	}

	_, err = s.NewsCategoryDetailRepository.UpsertNewsCategoryDetail(details)
	if err != nil {
		return fmt.Errorf("failed to UpsertNewsCategoryDetail: %v", err)
	}

	return nil
}
