package server

import (
	pmodel "crawler/model/comic/proto"
	"crawler/share/log"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestCrawlComicDetail(t *testing.T) {
	s := ComicServer{}
	// 50758
	bs, err := s.crawlComicDetail(33322)
	if err != nil {
		t.Log(err)
		return
	}

	var res pmodel.ComicDetailResponse
	err = proto.Unmarshal(bs, &res)
	if err != nil {
		t.Log(err)
		return
	}

	log.LogJson(&res)

}
