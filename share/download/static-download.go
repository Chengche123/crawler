package download

import (
	"crawler/share/log"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

type staticRes struct {
	url  string // http://images.dmzj1.com/webpic/19/jiubaotongxuebufangguowojsi.jpg
	path string // static/webpic/19/
	file string // jiubaotongxuebufangguowojsi.jpg
}

func downloadStatic(apis []string) error {
	res := make([]*staticRes, 0, len(apis))

	for _, api := range apis {
		s, err := parseStaticUrl(api)
		if err != nil {
			log.Logger.Info("failed to parse api", zap.String("api", api), zap.String("err", err.Error()))
			continue
		}
		res = append(res, s)
	}

	for _, v := range res {
		// 要用defer 用函数包起来
		func() {
			err := createPath(v.path)
			if err != nil {
				log.Logger.Info("failed to create path", zap.String("path", v.path), zap.String("err", err.Error()))
				return
			}

			filePath := v.path + v.file
			file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_EXCL, os.ModePerm)
			if err != nil {

				if os.IsExist(err) {
					log.Logger.Info("file already exist", zap.String("filePath", filePath))
					return
				}

				log.Logger.Info("failed to OpenFile",
					zap.String("filePath", filePath),
					zap.String("err", err.Error()))
				return
			}
			defer file.Close()

			r, err := http.Get(v.url)
			if err != nil {
				log.Logger.Info("failed to GET",
					zap.String("url", v.url),
					zap.String("err", err.Error()))
				return
			}
			defer r.Body.Close()

			_, err = io.Copy(file, r.Body)
			if err != nil {
				log.Logger.Info("failed to copy res.body to file", zap.String("err", err.Error()))
				return
			}
		}()
	}

	return nil
}

func createPath(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func parseStaticUrl(URL string) (*staticRes, error) {
	p, err := url.Parse(URL)
	if err != nil {
		return nil, fmt.Errorf("invalid url: %v", err)
	}
	file := filepath.Base(p.Path)
	return &staticRes{
		file: file,
		path: "static" + p.Path[:len(p.Path)-len(file)],
		url:  URL,
	}, nil
}
