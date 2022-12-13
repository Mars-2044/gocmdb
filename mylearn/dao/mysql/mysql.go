package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var db *gorm.DB

func Init() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)

	// db, err := gorm.Open("mysql", "saier:Cssbh123..@tcp(192.168.77.128:3306)/cmdb?charset=utf8&parseTime=True")
	db, err := gorm.Open("mysql", dsn)

	if err != nil{
		fmt.Errorf("创建数据库连接失败: %v", err)
	}

	db.DB().SetMaxOpenConns(viper.GetInt("mysql.max_conn"))
	db.DB().SetMaxIdleConns(viper.GetInt("mysql_idle_conn"))
	// defer db.Close()
	return db, err
}

func Close() {
	_ = db.Close()
}