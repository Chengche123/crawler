package jmodel

type ComicSpecial struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		ID         int    `json:"id"`
		Title      string `json:"title"`
		ShortTitle string `json:"short_title"`
		CreateTime int    `json:"create_time"`
		SmallCover string `json:"small_cover"`
		PageType   int    `json:"page_type"`
		Sort       int    `json:"sort"`
		PageURL    string `json:"page_url"`
	} `json:"data"`
}
