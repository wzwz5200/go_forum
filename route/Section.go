package route

import (
	service "web/Service"

	"github.com/gofiber/fiber/v2"
)

func CreateSection(R fiber.Router) {

	R.Post("/create_section", service.CreateSection) //创建分区

}

func GetAllSection(R fiber.Router) {

	R.Get("/get_all_section",service.GetAllSection)
}