package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"mylearn/models"
)

// 把每一层数据库操作封装成函数
// 待logic层根据业务需求调用
const secret = "siming"

func CheckUserExist(username string) (bool, error) {
	var count int

	sqlStr := `select count(user_id) from user where username = ?`

	err := db.Get(&count, sqlStr, username)
	if err != nil {
		//fmt.Println("查询数据失败")
		return false, err
	}
	return count > 0, nil
}

func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
  	user.Password = encryptPassword(user.Password)
	// 执行sql 语句入库
	sqlStr := `insert into user(user_id, username, password) values (?, ?, ?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)

	if err != nil {
		fmt.Printf("sql commit faild")
	}
	// db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

//func Login(p *models.ParamSignUP) error {
//	user := &models.User{
//		Username: p.Username,
//		Password: p.Password,
//	}
//	return mysql.Login(user)
//}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id, username, password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}

	if err != nil{
		// 查询数据库失败
		return err
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return errors.New("密码错误")
	}
	return
}