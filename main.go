package main

import (
	"github.com/gin-gonic/gin"
	"gotest/config"
	"gotest/middlewire"
	"gotest/router"
)

func main() {

	gin.SetMode(gin.DebugMode) // 默认为 debug 模式，设置为发布模式
	engine := gin.Default()
	engine.Use(middlewire.LoggerToFile())
	router.InitRouter(engine) // 设置路由
	engine.Run(config.PORT)
}