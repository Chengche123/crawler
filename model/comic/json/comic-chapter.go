package jmodel

type ComicChapterDetail struct {
	ChapterID    int      `json:"chapter_id"`
	ComicID      int      `json:"comic_id"`
	Title        string   `json:"title"`
	ChapterOrder int      `json:"chapter_order"`
	Direction    int      `json:"direction"`
	PageURL      []string `json:"page_url"`
	Picnum       int      `json:"picnum"`
	CommentCount int      `json:"comment_count"`
}
