package router

import (
	"github.com/gin-gonic/gin"
	login2 "youHUA/internal/login"
)

func Router(r *gin.Engine) {
	// 注册用户
	r.POST("/register", login2.Register)

	// 登录账号
	r.POST("/login", login2.Login)
}
