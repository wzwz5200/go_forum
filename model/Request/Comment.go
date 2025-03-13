package request

type ReqComment struct {
	Content string `json:"content"`
	PostID  uint   `json:"id"`
}

type ReqCommentDb struct {
	Content string `gorm:"type:text"` // 评论内容

	PostID uint `gorm:"index"` // 所属帖子ID（外键）
	UserID uint `gorm:"index"` // 评论者ID（外键）
}
