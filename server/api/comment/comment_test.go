package api

import "testing"

func TestFoo(t *testing.T) {
	t.Log(ComicCommentV2(50758, 0, 0))
	t.Log(ComicCommentCountV2(50758))
}
