package route

import (
	service "web/Service"

	"github.com/gofiber/fiber/v2"
)

func UserLogin(R fiber.Router) {

	R.Get("/login", service.Login)

}

func UserRegister(R fiber.Router) {

	R.Get("/register", service.Register)

}
