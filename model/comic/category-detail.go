package model

type CategoryDetail struct {
	ID             int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Title          string `json:"title"`
	Authors        string `json:"authors"`
	Status         string `json:"status"`
	Cover          string `json:"cover"`
	Types          string `json:"types"`
	LastUpdatetime int    `json:"last_updatetime"`
	Num            int    `json:"num"`
}
