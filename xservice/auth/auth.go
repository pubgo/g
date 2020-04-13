package auth

import (
	"github.com/gin-gonic/gin"
)

// api.github.com

func Init(g *gin.RouterGroup) {
	g.POST("/register", func(c *gin.Context) {})
	g.POST("/password", func(c *gin.Context) {})
	g.PUT("/password", func(c *gin.Context) {
		//	1. 给出新密码
		// 2 给出token
	})
	g.POST("/login", func(c *gin.Context) {
		//	 用户名，密码
		// web 获取cookie 和 session id
		// 移动端 返回token
		// 移动端 返回token
	})
	g.POST("/logout", func(c *gin.Context) {
		//	需要token或者cookie
	})
	g.POST("/verify", func(c *gin.Context) {
		// email phone
		//	 为什么验证，回调url
		// action
		// find password
		// email validate confirm
		// phone validate confirm
		// callback=
	})
}
