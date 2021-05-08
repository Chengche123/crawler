package model

import (
	pmodel "crawler/model/comic/proto"
)

type ComicDetail struct {
	ID                    int    `json:"Id" gorm:"primaryKey;autoIncrement;not null"`
	Title                 string `json:"Title"`
	Direction             int    `json:"Direction"`
	Islong                int    `json:"Islong"`
	Cover                 string `json:"Cover"`
	Description           string `json:"Description"`
	Lastupdatetime        int    `json:"LastUpdatetime"`
	Lastupdatechaptername string `json:"LastUpdateChapterName"`
	Firstletter           string `json:"FirstLetter"`
	Comicpy               string `json:"ComicPy"`
	Hotnum                int    `json:"HotNum"`
	Hitnum                int    `json:"HitNum"`
	Lastupdatechapterid   int    `json:"LastUpdateChapterId"`
	Subscribenum          int    `json:"SubscribeNum"`
	// Types                 []struct {
	// 	Tagid   int    `json:"TagId"`
	// 	Tagname string `json:"TagName"`
	// } `json:"Types"`
	// Status []struct {
	// 	Tagid   int    `json:"TagId"`
	// 	Tagname string `json:"TagName"`
	// } `json:"Status"`
	// Authors []struct {
	// 	Tagid   int    `json:"TagId"`
	// 	Tagname string `json:"TagName"`
	// } `json:"Authors"`
	// Chapters     []struct {
	// 	Title string `json:"Title"`
	// 	Data  []struct {
	// 		Chapterid    int    `json:"ChapterId"`
	// 		Chaptertitle string `json:"ChapterTitle"`
	// 		Updatetime   int    `json:"Updatetime"`
	// 		Filesize     int    `json:"Filesize"`
	// 		Chapterorder int    `json:"ChapterOrder"`
	// 	} `json:"Data"`
	// } `json:"Chapters"`
}

func NewComicDetail(pb *pmodel.ComicDetailResponse) *ComicDetail {
	return &ComicDetail{
		ID:                    int(pb.Data.Id),
		Title:                 pb.Data.Title,
		Direction:             int(pb.Data.Direction),
		Islong:                int(pb.Data.Islong),
		Cover:                 pb.Data.Cover,
		Description:           pb.Data.Description,
		Lastupdatetime:        int(pb.Data.LastUpdatetime),
		Lastupdatechaptername: pb.Data.LastUpdateChapterName,
		Firstletter:           pb.Data.FirstLetter,
		Comicpy:               pb.Data.ComicPy,
		Hotnum:                int(pb.Data.HotNum),
		Hitnum:                int(pb.Data.HitNum),
		Lastupdatechapterid:   int(pb.Data.LastUpdateChapterId),
		Subscribenum:          int(pb.Data.SubscribeNum),
	}
}
