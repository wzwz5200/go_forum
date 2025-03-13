package route

import (
	service "web/Service"

	"github.com/gofiber/fiber/v2"
)

func PostComment(R fiber.Router) {

	R.Post("/create_comment", service.Comment) //创建分区

}
