package model

import (
	"time"
)

// 用户模型

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"size:50;uniqueIndex" validate:"required,min=3,max=50"` // 用户名：必填，3-50 字符
	Email    string `gorm:"size:100;uniqueIndex" validate:"required,email"`       // 邮箱：必填，格式必须是 email
	Password string `gorm:"size:100" validate:"required,min=6,max=100"`           // 密码：必填，最少 6 个字符

	AvatarURL string `gorm:"size:512;default:'https://i.stardots.io/aver/89828644_p0.jpg?width=500&quality=50&blur=0&rotate=0'"`

	CreatedAt time.Time // 注册时间
	UpdatedAt time.Time // 最后更新时间

	
	Posts    []Post    `gorm:"foreignKey:AuthorID"` // 用户发的帖子
	Comments []Comment `gorm:"foreignKey:UserID"`   // 用户发的评论
}
