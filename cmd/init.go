package cmd

import (
	"web/Config"
	initdb "web/cmd/Initdb"
	"web/route"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func InitFiber() {

	//初始化数据库
	initdb.Initdb()
	app := fiber.New(Config.GetFiberConfig())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//测试demo api
	Human := app.Group("/Human")
	Human.Use(jwtware.New(Config.GetJwtConfig()))

	route.SetupHumanRoutes(Human)

	// 用户api组
	User := app.Group("/user")
	route.UserLogin(User)
	route.UserRegister(User)
	User.Use(jwtware.New(Config.GetJwtConfig()))

	route.UserTest(User)

	app.Listen(":3000")
}
