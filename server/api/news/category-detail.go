package api

import (
	rapi "crawler/server/api"
	"strconv"
)

// page从1开始
func NewsCategoryDetail(id, page int) string {
	// ${ApiUtil.BASE_URL_V4}/news/list/$id/${id == 0 ? 2 : 3}/$page

	return rapi.V4 + "/news/list/" + strconv.Itoa(id) + "/3/" + strconv.Itoa(page)
}
