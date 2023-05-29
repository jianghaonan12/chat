package dao

import (
	"GINCHAT/models"
	"GINCHAT/utils"
)

/*
*
查询用户信息
*/
var DB = utils.DB

func GetUserInfo(userVo *models.UserBasic) *models.UserBasic {
	user := new(models.UserBasic)
	DB.Debug().Table("user_basic").Where("name", userVo.Name).Scan(user)
	return user
}

/*
*
创建用户信息
*/
func CreateUserInfo(userVo *models.UserBasic) *models.UserBasic {
	DB.Debug().Table("user_basic").Create(userVo)
	return userVo
}

/*
*
更新用户信息
*/
func UpdateUserInfo(userVo []*models.UserBasic) bool {
	DB.Debug().Table("user_basic").Updates(userVo)
	return true
}
