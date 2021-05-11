package model

type NewsCategoryDetail struct {
	Articleid     int    `json:"ArticleId" gorm:"primaryKey;column:id"`
	Title         string `json:"Title"`
	Fromname      string `json:"FromName,omitempty"`
	Createtime    int    `json:"CreateTime"`
	Intro         string `json:"Intro"`
	Authorid      int    `json:"AuthorId"`
	Status        int    `json:"Status"`
	Rowpicurl     string `json:"RowPicUrl"`
	Colpicurl     string `json:"ColPicUrl"`
	Pageurl       string `json:"PageUrl"`
	Authoruid     int    `json:"AuthorUid"`
	Cover         string `json:"Cover"`
	Nickname      string `json:"Nickname"`
	Moodamount    int    `json:"MoodAmount,omitempty"`
	Commentamount int    `json:"CommentAmount,omitempty"`

	TagId int
}
