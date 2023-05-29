package router

import (
	"GINCHAT/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("/select", service.GetUserInfo)
		user.POST("/save", service.CreateUserInfo)
		user.POST("/update", service.UpdateUserInfo)
	}
	return r
}
