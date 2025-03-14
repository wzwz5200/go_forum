package route

import (
	service "web/Service"

	"github.com/gofiber/fiber/v2"
)

func GetPost(R fiber.Router) {

	R.Get("/allpost", service.GetAllPost) //获取所有帖子

}

func CreatePost(R fiber.Router) {

	R.Post("/createpost", service.CreatePost) //创建帖子

}

func GetPostDetails(R fiber.Router) {

	R.Get("/postd/:id", service.GetPostDetails)
}
