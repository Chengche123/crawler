package server

import (
	model "crawler/model/comic"
	jmodel "crawler/model/comic/json"
	pmodel "crawler/model/comic/proto"
	comicApi "crawler/server/api/comic"
	"crawler/share/log"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"unsafe"

	"go.uber.org/zap"
)

func (s *ComicServer) StartCrawlComicChapter() error {
	cdetails, err := s.ComicCommentRepository.FindAllComicId()
	if err != nil {
		return fmt.Errorf("failed to get comic id from comment repository: %v", err)
	}

	inchan := make(chan int, len(cdetails))
	for i := 0; i < len(cdetails); i++ {
		inchan <- cdetails[i]
	}
	close(inchan)
	outchan := make(chan *pmodel.ComicDetailResponse, 16)

	var wg sync.WaitGroup
	concur := 64

	for i := 0; i < concur; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for id := range inchan {
				buf, err := s.crawlComicDetail(id)
				if err != nil {
					log.Logger.Info("",
						zap.Int("comic id", id),
						zap.Error(err),
					)
					continue
				}

				pb, err := pmodel.UnmashalComicDetailResponse(buf)
				if err != nil {
					log.Logger.Info("failed to UnmashalComicDetailResponse",
						zap.String("id", strconv.Itoa(id)),
						zap.String("obj", fmt.Sprintf("%+v", pb)))
					continue
				}

				if pb.Errno == 3 {
					// 暂时无法观看
					log.Logger.Info("暂时无法观看",
						zap.Int("comic id", id))
					continue
				}

				outchan <- pb
			}

		}()
	}

	go func() {
		wg.Wait()
		close(outchan)
	}()

	var syncWg sync.WaitGroup
	for po := range outchan {
		syncWg.Add(1)

		go func(po *pmodel.ComicDetailResponse) {
			defer syncWg.Done()
			s.SaveComicChapter(po)
		}(po)

	}

	syncWg.Wait()

	return nil
}

func (s *ComicServer) SaveComicChapter(o *pmodel.ComicDetailResponse) {
	if o.Data == nil || o.Data.Chapters == nil {
		return
	}

	// TODO: 优化
	mos := make([]model.ComicChapter, 0)

	for _, v := range o.Data.Chapters {
		for _, v1 := range v.Data {
			mos = append(mos, model.ComicChapter{
				Chapterid:    int(v1.ChapterId),
				Chaptertitle: v1.ChapterTitle,
				Updatetime:   int(v1.Updatetime),
				Filesize:     int(v1.Filesize),
				Chapterorder: int(v1.ChapterOrder),

				Title:   v.Title,
				ComicId: int(o.Data.Id),
			})
		}
	}

	// 补全page_url等字段
	inchan := make(chan *model.ComicChapter, len(mos))
	for i := 0; i < len(mos); i++ {
		inchan <- &mos[i]
	}
	close(inchan)
	outchan := make(chan *model.ComicChapter, 1)

	var wg sync.WaitGroup
	concur := 16
	if concur > len(mos) {
		concur = len(mos)
	}

	for i := 0; i < concur; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for v := range inchan {

				func() { // 方便defer
					url := comicApi.ComicChapterDetail(v.Chapterid, v.ComicId)
					r, err := http.Get(url)
					if err != nil {
						return
					}
					defer r.Body.Close()

					var jo jmodel.ComicChapterDetail
					bs, err := io.ReadAll(r.Body)
					if err != nil {
						return
					}
					err = json.Unmarshal(bs, &jo)
					if err != nil {
						return
					}

					v.Picnum = jo.Picnum
					v.Direction = jo.Direction
					v.CommentCount = jo.CommentCount

					bs1, err := json.Marshal(&jo.PageURL)
					if err != nil {
						return
					}
					v.PageUrl = *(*string)(unsafe.Pointer(&bs1))

					outchan <- v
				}()
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outchan)
	}()

	fmos := make([]model.ComicChapter, 0, len(mos))
	for v := range outchan {
		fmos = append(fmos, *v)
	}

	_, err := s.ComicChapterRepository.UpsertComicComment(fmos)
	if err != nil {
		log.Logger.Info("",
			zap.Error(err),
			zap.Any("pb", o),
		)
	}
}
