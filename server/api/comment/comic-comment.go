package api

import (
	"fmt"
	"time"
)

func ComicCommentCountV2(id int) string {
	return fmt.Sprintf("https://interface.dmzj1.com/api/NewComment2/total?type=%d&obj_id=%d&countType=1&authorId=&_=%d",
		TYPE_COMIC, id, time.Now().Unix())
}

func ComicCommentV2(id, page, isHot int) string {
	return fmt.Sprintf("https://interface.dmzj1.com/api/NewComment2/list?type=%d&obj_id=%d&hot=%d&page_index=%d&_=1620531683470",
		TYPE_COMIC, id, isHot, page)
}
