package model

type Related struct {
	AuthorComics []struct {
		AuthorName string `json:"author_name"`
		AuthorID   int    `json:"author_id"`
		Data       []struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Cover  string `json:"cover"`
			Status string `json:"status"`
		} `json:"data"`
	} `json:"author_comics"`
	ThemeComics []struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Cover  string `json:"cover"`
		Status string `json:"status"`
	} `json:"theme_comics"`
	Novels []struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Cover  string `json:"cover"`
		Status string `json:"status"`
	} `json:"novels"`
}
