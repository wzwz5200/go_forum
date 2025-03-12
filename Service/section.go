package service

import "github.com/gofiber/fiber/v2"

func CreateSection(c *fiber.Ctx) error {

	c.SendString("Hello, World!")
	return c.SendStatus(200)

}
