package main

import (
	dao "todos/dao"
	models "todos/models"
	"todos/routers"
)

func main() {
	// 创建数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	// 程序退出关闭数据库
	defer dao.Close()
	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	// 注册路由
	router := routers.SetupRouter()
	router.Run()
}
