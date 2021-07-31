package dao

import (
	"crawler/share/config"
	"testing"
)

func TestNewComicChapterRepository(t *testing.T) {
	r, err := NewComicChapterRepository(config.MysqlDSN)
	if err != nil {
		t.Error(err)
		return
	}
	defer r.Close()
}
