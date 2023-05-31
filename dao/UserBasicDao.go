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
	DB.Debug().Table("user_basic").Where("id", userVo.ID).Find(user)
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
func UpdateUserInfo(user *models.UserBasic) *models.UserBasic {
	DB.Debug().Model(user).Update("name", user.Name)
	return user

}
