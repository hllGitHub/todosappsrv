package routers

import (
	"todos/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化路由
func SetupRouter() *gin.Engine {
	router := gin.Default()
	// 告诉 gin 框架模板文件引用的静态文件目录
	router.Static("/static", "static")
	// 告诉 gin 框架模板文件目录
	router.LoadHTMLGlob("templates/*")
	router.GET("/", controller.IndexHandler)

	// v1
	v1Group := router.Group("v1")
	{
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 获取所有
		v1Group.GET("/todos", controller.GetTodoList)
		// 更新
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除
		v1Group.DELETE("/todos/:id", controller.DeleteATodo)
	}

	return router
}
