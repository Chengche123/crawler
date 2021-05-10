package dao

import (
	"crawler/share/config"
	"testing"
)

func TestNewComicSpecialRepository(t *testing.T) {
	_, err := NewComicSpecialRepository(config.MysqlDSN)
	if err != nil {
		t.Error(err)
		return
	}
}
