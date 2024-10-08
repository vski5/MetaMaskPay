package main

import (
	"back/controllers"
	"back/ini"
	"back/routers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	ini.InitConfig("./ini/config.ini")

	router := gin.Default()

	// 配置CORS中间件
	config := cors.DefaultConfig()

	// 设置允许的源
	config.AllowOrigins = ini.GetCORSOrigins()

	// 设置允许的HTTP方法
	config.AllowMethods = ini.GetCORSMethods()

	// 设置允许的HTTP头
	config.AllowHeaders = ini.GetCORSHeaders()

	// 设置是否允许携带凭证
	config.AllowCredentials = ini.GetCORSAllowCredentials()

	// 应用CORS中间件到路由器
	router.Use(cors.New(config))

	// 初始化控制器
	var purchaseController controllers.PurchaseController = controllers.NewPurchaseController()

	// 设置路由
	routers.SetupPurchaseRoutes(router, purchaseController)

	// 启动服务器
	serverPort := ini.GetServerPort()
	router.Run(":" + serverPort) // 监听并在指定端口上启动服务
}
