package model

type ComicCommentCount struct {
	Result int    `json:"result"`
	Msg    string `json:"msg"`
	Data   int    `json:"data"`
}

type ComicComment struct {
	ID               int    `json:"id" gorm:"primaryKey"` // 评论id
	IsPassed         int    `json:"is_passed"`            // 审核是否通过,无论审核是否通过自己发的评论都能看到
	TopStatus        int    `json:"top_status"`           // -
	IsGoods          int    `json:"is_goods"`             // -
	UploadImages     string `json:"upload_images"`        // 评论图像
	ObjID            int    `json:"obj_id"`               // 漫画/新闻/专题/轻小说 id
	Content          string `json:"content"`              // 评论内容
	SenderUID        int    `json:"sender_uid"`           // 评论用户id
	LikeAmount       int    `json:"like_amount"`          // 点赞数量
	CreateTime       int    `json:"create_time"`          // 添加时间
	ToUID            int    `json:"to_uid"`               // 被评论用户,可能为空
	ToCommentID      int    `json:"to_comment_id"`        // 被评论的评论,可能为空
	OriginCommentID  int    `json:"origin_comment_id"`    // -
	ReplyAmount      int    `json:"reply_amount"`         // 回复数量
	HotCommentAmount int    `json:"hot_comment_amount"`   // 和回复数量相同
	Cover            string `json:"cover"`                // -
	Nickname         string `json:"nickname"`             // 用户名
	AvatarURL        string `json:"avatar_url"`           // 用户头像
	Sex              int    `json:"sex"`                  // 用户性别
	Mastercommentnum int    `json:"masterCommentNum"`     // 主评论数量
	// Mastercomment    []struct {
	// 	ID               int    `json:"id"`
	// 	IsPassed         int    `json:"is_passed"`
	// 	TopStatus        int    `json:"top_status"`
	// 	IsGoods          int    `json:"is_goods"`
	// 	UploadImages     string `json:"upload_images"`
	// 	ObjID            int    `json:"obj_id"`
	// 	Content          string `json:"content"`
	// 	SenderUID        int    `json:"sender_uid"`
	// 	LikeAmount       int    `json:"like_amount"`
	// 	CreateTime       int    `json:"create_time"`
	// 	ToUID            int    `json:"to_uid"`
	// 	ToCommentID      int    `json:"to_comment_id"`
	// 	OriginCommentID  int    `json:"origin_comment_id"`
	// 	ReplyAmount      int    `json:"reply_amount"`
	// 	Cover            string `json:"cover"`
	// 	Nickname         string `json:"nickname"`
	// 	HotCommentAmount int    `json:"hot_comment_amount"`
	// 	Sex              int    `json:"sex"`
	// } `json:"masterComment,omitempty"`
}
