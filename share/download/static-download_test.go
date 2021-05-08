package download

import "testing"

func TestParseStaticUrl(t *testing.T) {
	r, err := parseStaticUrl("http://images.dmzj1.com/webpic/19/jiubaotongxuebufangguowojsi.jpg")
	if err != nil {
		t.Log(err)
		return
	}

	wantFile := "jiubaotongxuebufangguowojsi.jpg"
	if r.file != wantFile {
		t.Errorf("file want %q,got %q", wantFile, r.file)
	}

	wantPath := "static/webpic/19/"
	if r.path != wantPath {
		t.Errorf("path want %q,got %q", wantPath, r.path)
	}
}

func TestCreatePath(t *testing.T) {
	r, err := parseStaticUrl("http://images.dmzj1.com/webpic/19/jiubaotongxuebufangguowojsi.jpg")
	if err != nil {
		t.Log(err)
		return
	}

	err = createPath(r.path)
	if err != nil {
		t.Log(err)
		return
	}
}

func TestDownloadStatic(t *testing.T) {
	err := downloadStatic([]string{
		"http://images.dmzj1.com/webpic/4/kdjdnhai20210305a.jpg",
		"http://images.dmzj1.com/webpic/18/bdonjl20210424.jpg",
		"http://images.dmzj1.com/webpic/19/jiubaotongxuebufangguowojsi.jpg",
		"http://images.dmzj1.com/webpic/18/zaijianrenshengnihaolongsheng.jpg",
		"http://images.dmzj1.com/webpic/18/200612kdzy.jpg",
	})
	if err != nil {
		t.Log(err)
		return
	}
}
