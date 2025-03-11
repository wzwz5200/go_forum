package service

import (
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

	return c.JSON(fiber.Map{
		"message": "User registered successfully",
		"user":    newUser,
	})
}
