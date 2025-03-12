package route

import (
	service "web/Service"

	"github.com/gofiber/fiber/v2"
)

func CreateSection(R fiber.Router) {

	R.Post("/create_section", service.CreateSection) //获取所有帖子

}
