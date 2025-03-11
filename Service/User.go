package service

import (
	hashp "web/HashP"
	initdb "web/cmd/Initdb"
	"web/model"
	request "web/model/Request"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func Login(c *fiber.Ctx) error {

	return c.SendString("Hello Login")
}

func Register(c *fiber.Ctx) error {

	newUser := request.User{}
	db := initdb.ReDB

	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// 进行数据验证
	if err := validate.Struct(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "数据不合法",
		})
	}
	//检查name,email是否存在
	var existingUser model.User
	if err := db.Where("email = ? OR username = ?", newUser.Email, newUser.Username).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Username or email already exists",
		})
	}

	//加密密码

	hashPwd, err := hashp.HashPassword(newUser.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "加密密码错误",
		})

	}

	//创建用户
	newUsers := model.User{
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: hashPwd,
	}

	if err := db.Create(&newUsers).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User registered successfully",
		"pwd":     hashPwd,
	})
}
