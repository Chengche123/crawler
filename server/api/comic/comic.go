package api

import (
	rapi "crawler/server/api"
	"strconv"
	"strings"
)

var ComicCategory = func() string {
	return rapi.Host + "/0/category.json"
}()

var ComicCategoryFilter = func() string {
	return rapi.Host + "/classify/filter.json"
}()

// ComicCategoryDetail 0:人气排序 1:更新排序
func ComicCategoryDetail(ids []int, sort, page int) string {
	var idstr string

	if len(ids) == 0 {
		idstr = "0"
	} else {
		tmps := make([]string, 0, len(ids))
		for _, v := range ids {
			tmps = append(tmps, strconv.Itoa(v))
		}

		idstr = strings.Join(tmps, "-")
	}

	return rapi.Host + "/classify/" + idstr + "/" + strconv.Itoa(sort) + "/" + strconv.Itoa(page) + ".json"
}
