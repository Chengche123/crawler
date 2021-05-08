package server

import (
	pmodel "crawler/model/comic/proto"
	"encoding/json"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestCrawlComicDetail(t *testing.T) {
	s := ComicServer{}
	bs, err := s.crawlComicDetail(50758)
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

	jstr, err := json.Marshal(&res)
	if err != nil {
		t.Log(err)
		return
	}

	t.Log(string(jstr))
}
