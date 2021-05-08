package api

import "testing"

func TestComicCategoryDetail(t *testing.T) {
	if ComicCategoryDetail([]int{2304, 8, 13}, 0, 0) != "https://nnv3api.dmzj1.com/classify/2304-8-13/0/0.json" {
		t.Error("")
	}
}

func TestComicDetail(t *testing.T) {
	t.Log(ComicDetail(33322))
}

func TestFoo(t *testing.T) {
	// 获取漫画分类
	t.Log(ComicCategory)
	// 获取漫画分类的条件
	t.Log(ComicCategoryFilter)
	// 根据id获取漫画详情
	t.Log(ComicDetail(50758))
	// 根据漫画分类获取漫画
	t.Log(ComicCategoryDetail([]int{2304}, 0, 0))
}
