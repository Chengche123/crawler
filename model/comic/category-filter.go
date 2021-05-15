package model

type ComicCategoryFilter struct {
	Id      int64  `gorm:"primaryKey;autoIncrement;not null"`
	Title   string `gorm:"not null;type:varchar(255)"`
	TagID   int
	TagName string
}
