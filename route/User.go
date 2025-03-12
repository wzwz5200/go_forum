package route

import (
	service "web/Service"

	"github.com/gofiber/fiber/v2"
)

func UserLogin(R fiber.Router) {

	R.Post("/login", service.Login)

}

func UserRegister(R fiber.Router) {

	R.Post("/register", service.Register)

}

func UserTest(R fiber.Router) {

	R.Get("/test", service.Test)

}
