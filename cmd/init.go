package cmd

import (
	"web/Config"
	initdb "web/cmd/Initdb"
	"web/route"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func InitFiber() {
	// 初始化数据库
	initdb.Initdb()
	app := fiber.New(Config.GetFiberConfig())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// 创建 API 父路由组
	api := app.Group("/api") // <-- 新增 API 父级路由组

	// 测试demo api（保持原路径/Human）
	Human := app.Group("/Human")
	Human.Use(jwtware.New(Config.GetJwtConfig()))
	route.SetupHumanRoutes(Human)

	// 用户api组（路径改为 /api/user）
	User := api.Group("/user") // <-- 改为 API 的子路由组
	{
		// 开放路由（无需 JWT）
		route.UserLogin(User)    // 路径: /api/user/login
		route.UserRegister(User) // 路径: /api/user/register

		// 受保护路由（应用 JWT 中间件）
		User.Use(jwtware.New(Config.GetJwtConfig()))
		route.UserTest(User) // 路径: /api/user/test
	}

	// 帖子API组（保持原路径/post）
	POST := api.Group("/post")
	{
		route.GetPost(POST)

		POST.Use(jwtware.New(Config.GetJwtConfig()))
		route.CreatePost(POST)
	}
	route.GetPost(POST)

	app.Listen(":3000")
}
