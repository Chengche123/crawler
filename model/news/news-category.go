package model

type NewsCategory struct {
	TagID   int    `json:"tag_id" gorm:"primaryKey;column:id"`
	TagName string `json:"tag_name"`
}
