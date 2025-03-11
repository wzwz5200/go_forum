package service

import (
	"strconv"
	initdb "web/cmd/Initdb"
	"web/model"

	"github.com/gofiber/fiber/v2"
)

func GetAllPost(c *fiber.Ctx) error {
	db := initdb.ReDB
	cursor, _ := strconv.Atoi(c.Query("cursor", "0")) // 默认0表示首次请求
	limit := 20

	var posts []model.Post

	// 关键修改：添加预加载
	query := db.
		Preload("Author").  // 加载用户数据
		Preload("Section"). // 加载板块数据
		Order("id ASC").
		Limit(limit + 1) // 多取1条用于判断是否有更多数据

	if cursor > 0 {
		query = query.Where("id > ?", cursor)
	}

	result := query.Find(&posts)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "数据库查询失败"})
	}

	// 分页逻辑
	hasMore := len(posts) > limit
	if hasMore {
		posts = posts[:limit] // 截断多余的那条
	}

	// 获取下一次的游标
	nextCursor := uint(0)
	if len(posts) > 0 {
		nextCursor = posts[len(posts)-1].ID
	}

	return c.JSON(fiber.Map{
		"data": posts,
		"pagination": fiber.Map{
			"cursor":   nextCursor,
			"has_more": hasMore,
		},
	})
}
