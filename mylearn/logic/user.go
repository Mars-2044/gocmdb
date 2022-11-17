package logic

import (
	"mylearn/dao/mysql"
	"mylearn/pkg/snowflake"
)

// 存放业务逻辑的代码

func SignUp() {
	// 1.判断用户是否存在
	mysql.QueryUserByUsername()
	// 生成uid
	snowflake.GetID()

	// 保存进数据库
	mysql.InsertUser()
}
