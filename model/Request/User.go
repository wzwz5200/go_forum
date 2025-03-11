package request

type User struct {
	Username string `gorm:"size:50;uniqueIndex" validate:"required,min=3,max=50"` // 用户名：必填，3-50 字符
	Email    string `gorm:"size:100;uniqueIndex" validate:"required,email"`       // 邮箱：必填，格式必须是 email
	Password string `gorm:"size:100" validate:"required,min=6,max=100"`           // 密码：必填，最少 6 个字符

}
