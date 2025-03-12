// config.go
package Config

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func GetFiberConfig() fiber.Config {
	return fiber.Config{
		ServerHeader:      "Fiber",
		AppName:           "Test App v1.0.1",
		EnablePrintRoutes: true,
	}
}

func GetJwtConfig() jwtware.Config {

	return jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("128wang123")},
		ContextKey: "jwt", // 上下文中存储令牌的键名
	}

}
