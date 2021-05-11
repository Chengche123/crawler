package pmodel

import (
	api "crawler/server/api/news"
	"crawler/share/crypto"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestFoo(t *testing.T) {
	URL := api.NewsCategoryDetail(1, 1)
	r, err := http.Get(URL)
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

	msg, err := crypto.Decrypt(string(bs))
	if err != nil {
		t.Error(err)
		return
	}

	var detail NewsListResponse
	err = proto.Unmarshal(msg, &detail)
	if err != nil {
		t.Error(err)
		return
	}

	jbs, err := json.MarshalIndent(&detail, "", "  ")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(jbs))
}
