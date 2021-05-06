package api

import "testing"

func TestComicCategoryDetail(t *testing.T) {
	if ComicCategoryDetail([]int{2304, 8, 13}, 0, 0) != "https://nnv3api.dmzj1.com/classify/2304-8-13/0/0.json" {
		t.Error("")
	}
}
