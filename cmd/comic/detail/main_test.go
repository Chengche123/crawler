package main

import (
	pmodel "crawler/model/comic/proto"
	api "crawler/server/api/comic"
	"crawler/share/crypto"
	"crawler/share/log"
	"io"

	"net/http"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestComicDetail(t *testing.T) {
	r, err := http.Get(api.ComicDetail(33322))
	if err != nil {
		t.Error(err)
		return
	}
	defer r.Body.Close()

	bs, _ := io.ReadAll(r.Body)

	dt, err := crypto.Decrypt(string(bs))
	if err != nil {
		t.Error(err)
		return
	}

	var obj pmodel.ComicDetailResponse
	err = proto.Unmarshal(dt, &obj)
	if err != nil {
		t.Error(err)
		return
	}

	log.LogJson(&obj)
}
