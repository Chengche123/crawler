package model

type ComicType struct {
	Id      int `gorm:"primaryKey;column:id"`
	ComicId int
	TagId   int
}
