package jmodel

type ComicCategoryFilter struct {
	Title string `json:"title"`
	Items []struct {
		TagID   int    `json:"tag_id"`
		TagName string `json:"tag_name"`
	} `json:"items"`
}
