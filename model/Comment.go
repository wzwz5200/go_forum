package model

import "time"

// 评论模型
type Comment struct {
    ID        uint      `gorm:"primaryKey"`
    Content   string    `gorm:"type:text"`     // 评论内容
    PostID    uint      `gorm:"index"`         // 外键：posts.id
    UserID    uint      `gorm:"index"`         // 外键：users.id
    CreatedAt time.Time                        // 评论时间

    // 关联关系（保持与Post结构体相同的Author命名）
    Author    User      `gorm:"foreignKey:UserID"`  // 修改字段名称为 Author
    Post      Post      `gorm:"foreignKey:PostID"`
}