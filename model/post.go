package model

import "time"

type Post struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"size:200"`       // 标题
	Content   string `gorm:"type:text"`      // 内容
	
	AuthorID  uint   `gorm:"index"`          // 作者ID（外键）
	SectionID uint   `gorm:"index"`          // 所属板块ID（外键）
	
	CreatedAt time.Time                     // 发帖时间
	UpdatedAt time.Time                     // 最后更新时间
	ViewCount int     `gorm:"default:0"`    // 浏览数统计

	// 关联关系
	Author    User    `gorm:"foreignKey:AuthorID"` 
	Section   Section `gorm:"foreignKey:SectionID"`
	Comments  []Comment `gorm:"foreignKey:PostID"`
}