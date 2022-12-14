package models

// 定义请求的参数结构体

type ParamSignUP struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type LoginSignUP struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}