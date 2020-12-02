package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

//RecoverMiddleware 用于捕获panic，返回错误信息
func RecoverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				//先记录日志
				//log.Println(string(debug.Stack()))
				log.Println("panic happened!!")
				//返回客户端错误
				c.AbortWithStatus(500)
			}
		}()
	}
}
