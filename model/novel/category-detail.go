package model

type NovelCategoryDetail struct {
	ID      int `gorm:"primaryKey;column:id"`
	Cover   string
	Name    string
	Authors string

	TagId int
}
