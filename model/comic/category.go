package model

type ComicCategory struct {
	TagID int    `json:"tag_id" gorm:"primaryKey;column:id"`
	Title string `json:"title"`
	Cover string `json:"cover"`
}
