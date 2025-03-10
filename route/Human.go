package route

import "github.com/gofiber/fiber/v2"

func SetupHumanRoutes(Human fiber.Router) {
	Human.Get("/say", SayHello) // 绑定路由处理函数
}

// SayHello 处理函数（需首字母大写，允许跨包访问）
func SayHello(c *fiber.Ctx) error {
	return c.SendString("Hello from another file!")
}
