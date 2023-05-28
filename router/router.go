package router

import (
	"GINCHAT/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/user", service.GetUserInfo)
	return r
}
