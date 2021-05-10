package model

type ComicSpecial struct {
	ID         int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title      string `json:"title"`
	ShortTitle string `json:"short_title"`
	CreateTime int    `json:"create_time"`
	SmallCover string `json:"small_cover"`
	PageType   int    `json:"page_type"` // -
	Sort       int    `json:"sort"`      // -
	PageURL    string `json:"page_url"`  // http://m.dmzj.com/zhuanti/$page_url
}
