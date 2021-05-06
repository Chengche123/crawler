package server

import "testing"

func TestCheckCategoryDetailPage(t *testing.T) {
	s := ComicServer{}
	if s.checkCategoryDetailPage("https://nnv3api.dmzj1.com/classify/2304-8-13/0/50.json") == false {
		t.Error("")
	}

	if s.checkCategoryDetailPage("https://nnv3api.dmzj1.com/classify/2304-8-13/0/9999.json") == true {
		t.Error("")
	}
}
