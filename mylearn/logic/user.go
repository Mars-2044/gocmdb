package logic

import (
	"errors"
	"mylearn/dao/mysql"
	"mylearn/models"
	"mylearn/pkg/snowflake"
)

// 存放业务逻辑的代码

func SignUp(p *models.ParamSignUP) (err error){
	// 1.判断用户是否存在
	var exist bool
	mysql.CheckUserExist(p.Username)

	if err != nil {
		// 数据库查询出错
		return err
	}

	if exist {
		return errors.New("用户已存在")
	}
	// 生成uid
	userID := snowflake.GetID()

	// 构造一个User实例
	user := &models.User{
		UserID: userID,
		Username: p.Username,
		Password: p.Password,
	}

	// 保存进数据库
	mysql.InsertUser(user)
	return
}
