package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"mylearn/models"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := "saier:Cssbh123..@tcp(192.168.77.128:3306)/cmdb?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

func Close() {
	_ = db.Close()
}

func Login(p *models.User) (err error) {
	oPassword := p.Password
	sqlStr := `select user_id, username, password from user where username=?`
	err = db.Get(p, sqlStr, p.Username)
	if err == sql.ErrNoRows {
		return errors.New("用户不存在")
	}

	if err != nil{
		// 查询数据库失败
		return err
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != p.Password {
		return errors.New("密码错误")
	}
	return err
}


//var db *gorm.DB
//
//func Init() (*gorm.DB, error) {
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
//		viper.GetString("mysql.user"),
//		viper.GetString("mysql.password"),
//		viper.GetString("mysql.host"),
//		viper.GetInt("mysql.port"),
//		viper.GetString("mysql.dbname"),
//	)
//
//	// db, err := gorm.Open("mysql", "saier:Cssbh123..@tcp(192.168.77.128:3306)/cmdb?charset=utf8&parseTime=True")
//	db, err := gorm.Open("mysql", dsn)
//
//	if err != nil{
//		fmt.Errorf("创建数据库连接失败: %v", err)
//	}
//
//	db.DB().SetMaxOpenConns(viper.GetInt("mysql.max_conn"))
//	db.DB().SetMaxIdleConns(viper.GetInt("mysql_idle_conn"))
//	// defer db.Close()
//	return db, err
//}
//
//func Close() {
//	_ = db.Close()
//}