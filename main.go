// Package goServer
// @Time  : 2021/10/19 18:03  
// @Author: mouse_zy@qq.com 
// @note  : 登录和注册
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"goServer/app"
	"goServer/docs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// newRoute 注册路由
func newRoute() {
	// 1. 注册根路由
	route := gin.Default()
	{
		// 初始化swagger
		docs.SwaggerInfo.BasePath = "/app"
		route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		// 1.1 设置总路由分组 -- app
		apiRoute := route.Group("/app")
		{
			// 设置功能路由--登录注册
			apiRoute.POST("/register", app.Register)
			apiRoute.POST("/login", app.Login)

		}
		// 1.2 设置基础路由
		baseRoute := apiRoute.Group("/base")
		{
			// 设置验证码路由
			baseRoute.GET("/code", app.Code)
		}

	}

	// 2.启动服务-端口10000-40000
	err := route.Run(":14250")
	if err != nil {
		panic("端口被占用")
	}
}

// createDB 新建或打开数据库
func createDB() {
	// 1.0 打开或者创建数据库
	db, err := gorm.Open(sqlite.Open("userRegisterInfo"), &gorm.Config{})
	if err != nil {
		panic("打开数据库异常")
		return
	}
	//
	err = db.AutoMigrate(&app.UserRegisterInfo{})
	if err != nil {
		panic("创建表失败")
		return
	}
	app.GDB = db
}
func main() {
	fmt.Println("run code start ....")
	// 创建或打开数据库
	createDB()
	// 注册路由
	newRoute()
}
