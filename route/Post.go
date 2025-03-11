package route

import (
	service "web/Service"

	"github.com/gofiber/fiber/v2"
)

func GetPost(R fiber.Router) {

	R.Get("/allpost", service.GetAllPost) //获取所有帖子

}
