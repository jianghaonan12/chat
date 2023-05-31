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
	userInfo := dao.GetUserInfo(userVo)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "成功",
		"data": userInfo,
	})
}
func CreateUserInfo(c *gin.Context) {
	userVo := new(models.UserBasic)
	c.ShouldBindBodyWith(userVo, binding.JSON)
	user := dao.CreateUserInfo(userVo)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "成功",
		"data": user,
	})
}
func UpdateUserInfo(c *gin.Context) {
	user := new(models.UserBasic)
	//c.Header()
	c.ShouldBindBodyWith(user, binding.JSON)
	userInfo := dao.UpdateUserInfo(user)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "成功",
		"data": userInfo,
	})
}
