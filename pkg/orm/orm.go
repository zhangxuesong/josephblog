package orm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zhangxuesong/josephblog/pkg/config"
	"log"
)

var db *gorm.DB

func init() {
	log.Println("链接数据库。。。")
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.Mysql.UserName, config.Config.Mysql.Password, config.Config.Mysql.Address, config.Config.Mysql.Port, config.Config.Mysql.DbName))
	if err != nil {
		log.Panic("failed to connect database：%v", err)
	}

	db.SingularTable(true) // 关闭复数表名
	db.LogMode(true)       // 启用Logger，显示详细日志
	// 连接池
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	log.Println("链接数据库成功。。。")
}
