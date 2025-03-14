package service

import (
	"errors"
	"log"
	"strconv"
	initdb "web/cmd/Initdb"
	"web/model"
	request "web/model/Request"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func GetAllPost(c *fiber.Ctx) error {
	db := initdb.ReDB
	cursor, _ := strconv.Atoi(c.Query("cursor", "0")) // 默认0表示首次请求
	limit := 20

	var posts []model.Post

	// 关键修改：添加预加载
	query := db.
		Select("id", "title", "author_id", "section_id", "created_at", "updated_at", "view_count").
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

//创建帖子

func CreatePost(c *fiber.Ctx) error {
	db := initdb.ReDB.Debug() // 开启调试模式
	var req request.CreatePostRequest

	c.BodyParser(&req)
	// 类型安全获取用户ID
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	// 提取字段
	userIDStr, ok := claims["user"].(string)
	if !ok {
		log.Printf("JWT Claims 中缺少 user 字段或类型不匹配")
		return c.Status(400).JSON(fiber.Map{"error": "无效的 JWT"})
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Printf("userID 转换失败: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "无效的 JWT"})
	}
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("事务回滚: %v", r)
		}
	}()

	// 验证用户存在
	var author model.User
	if err := tx.First(&author, userID).Error; err != nil {
		tx.Rollback()
		log.Printf("用户查询失败: %v", err)
		return c.Status(404).JSON(fiber.Map{"error": "用户不存在"})
	}

	// 验证板块存在
	var section model.Section
	if err := tx.First(&section, req.SectionID).Error; err != nil {
		tx.Rollback()
		log.Printf("板块查询失败: %v", req)
		return c.Status(404).JSON(fiber.Map{"error": "板块不存在"})
	}

	// 创建帖子
	newPost := model.Post{
		Title:     req.Title,
		Content:   req.Content,
		AuthorID:  1,
		SectionID: req.SectionID,
	}

	if err := tx.Create(&newPost).Error; err != nil {
		tx.Rollback()
		log.Printf("帖子创建失败: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"error":   "创建失败",
			"details": err.Error(),
		})
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("事务提交失败: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "提交失败"})
	}

	// 返回数据（可选预加载）
	var createdPost model.Post
	if err := db.Preload("Author").Preload("Section").First(&createdPost, newPost.ID).Error; err != nil {
		log.Printf("数据预加载失败: %v", err)
	}

	return c.Status(201).JSON(fiber.Map{
		"data":    createdPost,
		"message": "创建成功",
	})
}

func GetSectionPost(c *fiber.Ctx) error {

	db := initdb.ReDB

	var section model.Section
	var Req request.ReqGetSection

	c.BodyParser(&Req)
	if err := db.Preload("Posts").
		Where("name = ?", Req.Name).
		First(&section).
		Error; err != nil {

		return c.Status(400).JSON(fiber.Map{"err": "数据库错误"})
	}
	if section.Posts == nil || len(section.Posts) == 0 {

		return c.Status(201).JSON(fiber.Map{"err": "当前分区无帖子"})

	}
	return c.Status(201).JSON(fiber.Map{"date": section.Posts, "req": Req})
}

func GetPostDetails(c *fiber.Ctx) error {

	db := initdb.ReDB
	postID := c.Params("id")
	if postID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Post ID is required"})
	}

	// 转换ID为uint类型
	id, err := strconv.ParseUint(postID, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid post ID format"})
	}

	var post model.Post
	// 查询数据库并预加载关联数据
	if err := db.Preload("Author").
		Preload("Section").
		Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Author") // 现在可以正确加载评论作者
		}).
		First(&post, id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Post not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve post"})
	}

	// 原子递增浏览数并更新返回的数据
	if err := db.Model(&post).UpdateColumn("view_count", gorm.Expr("view_count + 1")).Error; err != nil {
		log.Printf("View count update error: %v", err)
	} else {
		post.ViewCount++ // 手动递增以保证返回数据正确
	}

	// 返回帖子详情（依赖模型结构的json标签过滤敏感字段）
	return c.JSON(post)

}
