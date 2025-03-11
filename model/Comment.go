package model

import "time"

// 评论模型
type Comment struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"type:text"` // 评论内容

	PostID uint `gorm:"index"` // 所属帖子ID（外键）
	UserID uint `gorm:"index"` // 评论者ID（外键）

	CreatedAt time.Time // 评论时间

	// 关联关系
	Post Post `gorm:"foreignKey:PostID"`
	User User `gorm:"foreignKey:UserID"`
}
