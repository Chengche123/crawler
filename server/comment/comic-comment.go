package server

import (
	model "crawler/model/comment"
	api "crawler/server/api/comment"
	"crawler/share/log"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"go.uber.org/zap"
)

const (
	// 漫画的原来评论数不足此数时重新爬取
	commentLowerBound = 510
	// 希望爬取的评论数
	expectedComment = 600
	commentsPerPage = 30
)

func (s *CommentServer) StartCrawlComicComment(offset, limit int) error {
	details, err := s.ComicDetailRepository.FindByHotDESC(offset, limit)
	if err != nil {
		return fmt.Errorf("failed to get comic details: %v", err)
	}

	comicids := make([]int, 0, len(details))
	for _, v := range details {
		comicids = append(comicids, v.ID)
	}

	const concurLimit = 64

	inchan := make(chan int, len(comicids))
	for _, v := range comicids {
		inchan <- v
	}
	close(inchan)
	var wg sync.WaitGroup

	for i := 0; i < concurLimit; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for id := range inchan {
				err := s.startCrawlComicComment(id)
				if err != nil {
					log.Logger.Info("failed to get comments", zap.Error(err), zap.Int("comicid", id))
					continue
				}
			}
		}()
	}

	wg.Wait()

	return nil
}

func (s *CommentServer) StartCrawlComicCommentByID(comicid int) error {
	return s.startCrawlComicComment(comicid)
}

func (s *CommentServer) startCrawlComicComment(comicid int) error {
	sum, err := getComicCommentSum(comicid)
	if err != nil {
		return fmt.Errorf("failed to getComicCommentSum(%d): %v", comicid, err)
	}

	n, err := s.ComicCommentRepository.ComicCommentCount(comicid)
	if err != nil {
		return fmt.Errorf("cannot count comic comments: %v", err)
	}

	if sum == n {
		return fmt.Errorf("comments already crawl completed,comicid: %d,nums: %d", comicid, sum)
	}

	if n >= commentLowerBound {
		return fmt.Errorf("comic comment already >= %d", commentLowerBound)
	}

	const concurLimit = 20

	pages := expectedComment / commentsPerPage
	inchan := make(chan int, pages)
	for i := 1; i <= pages; i++ {
		inchan <- i
	}
	close(inchan)
	var wg sync.WaitGroup

	for i := 0; i < concurLimit; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for page := range inchan {
				func() {
					URL := api.ComicCommentV2(comicid, page, 0)
					r, err := http.Get(URL)
					if err != nil {
						log.Logger.Info("failed to GET", zap.Error(err), zap.Int("comicid", comicid), zap.Int("page", page))
						return
					}

					defer r.Body.Close()

					bs, err := ioutil.ReadAll(r.Body)
					if err != nil {
						log.Logger.Info("failed to ReadAll", zap.Error(err))
					}

					var comments []model.ComicComment

					err = json.Unmarshal(bs, &comments)
					if err != nil {
						log.Logger.Info("failed to Unmarshal", zap.Error(err), zap.String("URL", URL))
						return
					}

					_, err = s.ComicCommentRepository.UpsertComicComment(comments)
					if err != nil {
						log.Logger.Info("failed to UpsertComicComment", zap.Error(err))
						return
					}
				}()
			}
		}()
	}

	wg.Wait()

	return nil
}

func getComicCommentSum(comicid int) (int, error) {
	r, err := http.Get(api.ComicCommentCountV2(comicid))
	if err != nil {
		return 0, err
	}
	defer r.Body.Close()

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return 0, err
	}

	var res model.ComicCommentCount
	err = json.Unmarshal(bs, &res)
	if err != nil {
		return 0, err
	}

	if res.Data == 0 {
		return 0, fmt.Errorf("invalid count: 0")
	}

	return res.Data, nil
}
