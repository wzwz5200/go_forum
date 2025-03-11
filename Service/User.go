package service

import (
	"time"
	hashp "web/HashP"
	initdb "web/cmd/Initdb"
	"web/model"
	request "web/model/Request"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

//登陆

func Login(c *fiber.Ctx) error {

	db := initdb.ReDB
	NewUser := model.User{}
	Req := request.AuthRequest{}

	c.BodyParser(&Req)

	if err := db.Where("username = ?", Req.Username).First(&NewUser).Error; err != nil {
		// 处理用户不存在

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "用户名不存在",
		})
	}

	// 2. 验证密码（假设密码已哈希存储）
	if err := bcrypt.CompareHashAndPassword([]byte(NewUser.Password), []byte(Req.Password)); err != nil {
		// 处理密码错误

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "密码错误",
		})
	}

	claims := jwt.MapClaims{

		"sub": NewUser.ID, // 必须 - 用户唯一标识

		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("128wang123"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"name": NewUser.Username, "icon": NewUser.AvatarURL, "token": t})
	//return c.SendString("Hello Login")

}



//注册

func Register(c *fiber.Ctx) error {

	newUser := request.User{}
	db := initdb.ReDB

	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// 进行数据验证
	if err := validate.Struct(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "数据不合法",
		})
	}
	//检查name,email是否存在
	var existingUser model.User
	if err := db.Where("email = ? OR username = ?", newUser.Email, newUser.Username).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Username or email already exists",
		})
	}

	//加密密码

	hashPwd, err := hashp.HashPassword(newUser.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "加密密码错误",
		})

	}

	//创建用户
	newUsers := model.User{
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: hashPwd,
	}

	if err := db.Create(&newUsers).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User registered successfully",
		"pwd":     hashPwd,
	})
}

func Test(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{"HELLO": "w"})
}
