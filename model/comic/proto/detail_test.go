package pmodel

import (
	api "crawler/server/api/comic"
	"crawler/share/crypto"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestUnmashalComicDetailResponse(t *testing.T) {
	r, err := http.Get(api.ComicDetail(50758))
	if err != nil {
		t.Error(err)
		return
	}
	defer r.Body.Close()

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Error(err)
		return
	}

	cipher, err := crypto.Decrypt(string(bs))
	if err != nil {
		t.Error(err)
		return
	}

	res, err := UnmashalComicDetailResponse(cipher)
	if err != nil {
		t.Error(err)
		return
	}

	jbs, _ := json.MarshalIndent(&res, "", "  ")
	t.Log(string(jbs))
}
