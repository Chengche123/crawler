package model

type Recommend []struct {
	CategoryID int    `json:"category_id"`
	Title      string `json:"title"`
	Sort       int    `json:"sort"`
	Data       []struct {
		Cover    string `json:"cover"`
		Title    string `json:"title"`
		SubTitle string `json:"sub_title"`
		Type     int    `json:"type"`
		URL      string `json:"url"`
		ObjID    int    `json:"obj_id"`
		Status   string `json:"status"`
	} `json:"data"`
}
