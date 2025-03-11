package model


type Section struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:50;uniqueIndex"` // 板块名称唯一
	Description string `gorm:"type:text"`           // 板块描述
	
	Posts []Post `gorm:"foreignKey:SectionID"`     // 板块下的帖子
}