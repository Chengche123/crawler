package api

import (
	rapi "crawler/server/api"
	"strconv"
)

func ComicChapterDetail(chapterId, comicId int) string {
	return rapi.Host + "/chapter/" + strconv.Itoa(comicId) + "/" + strconv.Itoa(chapterId) + ".json"
}
