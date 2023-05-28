package service

import (
	"GINCHAT/dao"
	"GINCHAT/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	userVo := new(models.UserBasic)
	c.ShouldBindBodyWith(userVo, binding.JSON)
	userF := dao.GetUserInfo(userVo)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "成功",
		"data": userF,
	})
}
