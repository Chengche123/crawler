package api

import (
	"crawler/server/api"
	"strconv"
)

func ComicSpecial(page int) string {
	// https://nnv3api.dmzj1.com/subject/0/0.json?channel=windows&version3.0.0&timestamp=1620621502
	return api.Host + "/subject/0/" + strconv.Itoa(page) + ".json"
}
