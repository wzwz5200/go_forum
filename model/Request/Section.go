package request

type ReqSection struct {
	Name        string `json:"name" validate:"required,max=10"`
	Description string `json:"description" validate:"max=100"`
}

// 获取当前分区使用帖子请求体
type ReqGetSection struct {
	Name string `json:"name" validate:"required,max=10"`
}
