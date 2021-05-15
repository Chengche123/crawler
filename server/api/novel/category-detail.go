package api

import (
	rapi "crawler/server/api"
	"strconv"
	"strings"
)

/*
/// 轻小说类目详情
static String novelCategoryDetail(
	{int cateId = 0, int status = 0, int sort = 0, int page = 0}) {
return "$apiHost/novel/$cateId/$status/$sort/$page.json?${defaultParameter()}";
}
*/

func c(i int) string {
	return strconv.Itoa(i)
}

// status为进度id,可以在category-filter.json里找
func NovelCategoryDetail(id, status, sort, page int) string {
	params := []string{c(id), c(status), c(sort), c(page)}

	param := strings.Join(params, "/")

	return rapi.Host + "/novel/" + param + ".json"
}
