package dao

import (
	"GINCHAT/models"
	"GINCHAT/utils"
)

func GetUserInfo(userVo *models.UserBasic) *models.UserBasic {
	DB := utils.DB
	user := new(models.UserBasic)
	DB.Debug().Table("user_basic").Where("name", userVo.Name).Scan(user)
	return user
}
