package route

import (
	service "web/Service"

	"github.com/gofiber/fiber/v2"
)

func CreateSection(R fiber.Router) {

	R.Post("/create_section", service.CreateSection) //创建分区

}

//获取使用分区
func GetAllSection(R fiber.Router) {

	R.Get("/get_all_section",service.GetAllSection)
}

//传入分区id获取分区帖子
func GetSectionAllPost(R fiber.Router) {

	R.Get("/get_all_section_post",service.GetSectionPost)
}