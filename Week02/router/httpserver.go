package router

import (
	"week02/service"

	"github.com/gin-gonic/gin"
)

//HTTPServerRun 启动http服务
func HTTPServerRun() {
	gin.SetMode(gin.DebugMode)
	addr := "127.0.0.1:8080"
	svc := gin.Default()
	svc.GET("/ping", service.Pong)
	svc.GET("/panic", service.Panic) //携程panic的recover
	svc.GET("/GetUserDetails", service.GetUserDetails)
	svc.Run(addr)
}
