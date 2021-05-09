package dao

import (
	"crawler/share/config"
	"testing"
)

func TestNewComicCommentRepository(t *testing.T) {
	repo, err := NewComicCommentRepository(config.MysqlDSN)
	if err != nil {
		t.Error(err)
		return
	}

	defer repo.Close()
}
