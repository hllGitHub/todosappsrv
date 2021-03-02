package dao

import (
	"github.com/jinzhu/gorm"
	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// DB 初始化数据库变量
	DB *gorm.DB
)

// InitMySQL 初始化数据库
func InitMySQL() (err error) {
	dsn := "root:031599@/Todos?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

// Close 关闭数据库
func Close() {
	DB.Close()
}
