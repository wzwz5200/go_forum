package service

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {
	return c.SendString("Hello Login")
}

func Register(c *fiber.Ctx) error {
	
	return c.SendString("Hello Register")
}
