package model

// TODO api返回的类容发生改变,导致jmodel和model不相同
type ComicCategory struct {
	TagID int    `json:"tag_id" gorm:"primaryKey;column:id"`
	Title string `json:"title"`
	Cover string `json:"cover"`
}
