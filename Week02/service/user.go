package service

import (
	"log"
	"week02/dao"
	"week02/sync"

	"github.com/gin-gonic/gin"
)

//Pong 测试
func Pong(c *gin.Context) {
	c.String(200, "pong")
}

//Panic 测试Panic
func Panic(c *gin.Context) {
	sync.Go(func() {
		panic("gorotine panic")
	})
}

//GetUserDetails 获取用户信息的handler
func GetUserDetails(c *gin.Context) {
	u, e := dao.GetUserDetails("u1111")
	if e != nil { //出现其他错误了
		//记录日志
		log.Printf("[INFO] error occured:%+v\n", e)
		//返回错误
		c.String(200, "get user detail error")
		return
	}
	c.JSON(200, u)
}
