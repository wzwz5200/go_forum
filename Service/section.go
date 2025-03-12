package service

import (
	initdb "web/cmd/Initdb"
	"web/model"
	request "web/model/Request"

	"github.com/gofiber/fiber/v2"
)

func CreateSection(c *fiber.Ctx) error {

	db := initdb.ReDB
	var request request.ReqSection

	c.BodyParser(&request)

	newSection := model.Section{
		Name:        request.Name,
		Description: request.Description,
	}

	if db.Create(&newSection).Error != nil {
		return c.Status(400).JSON(fiber.Map{"date": "创建数据错误，请检查是否重复创建"})
	}

	return c.Status(200).JSON(fiber.Map{"date": "创建成功"})

}
