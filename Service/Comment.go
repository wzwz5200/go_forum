package service

import (
	"log"
	"strconv"
	initdb "web/cmd/Initdb"
	"web/model"
	request "web/model/Request"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Comment(c *fiber.Ctx) error {

	db := initdb.ReDB
	var req request.ReqComment

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

	NewComment := model.Comment{
		UserID:  uint(userID),
		Content: req.Content,
		PostID:  req.PostID,
	}

	if err = db.Create(&NewComment).Error; err != nil {

		return c.Status(400).JSON(fiber.Map{"error": "发生评论错误", "req": req})
	}

	return c.Status(200).JSON(fiber.Map{"status": "发送成功"})
}
