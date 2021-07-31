package model

type ComicChapter struct {
	Chapterid    int `gorm:"primaryKey;autoIncrement;not null;column:id"`
	Chaptertitle string
	Updatetime   int
	Filesize     int
	Chapterorder int

	Title   string
	ComicId int

	Direction    int
	PageUrl      string
	Picnum       int
	CommentCount int
}
