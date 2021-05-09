package dao

import (
	"crawler/share/config"
	"testing"
)

func TestComicDetailRepository(t *testing.T) {
	repo, err := NewComicDetailRepository(config.MysqlDSN)
	if err != nil {
		t.Error(err)
		return
	}
	defer repo.Close()

	n, err := repo.Count()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(n)

	// res, err := repo.FindByHotDESC(0, 10)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }

	// bs, err := json.MarshalIndent(&res, "", "  ")
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }

	// t.Log(string(bs))
}
