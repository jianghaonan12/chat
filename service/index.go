package service

import (
	"GINCHAT/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func GetIndex(c *gin.Context) {
	user := new(models.UserBasic)
	c.ShouldBindBodyWith(user, binding.JSON)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "成功",
		"data": user,
	})
}
