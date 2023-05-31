package models

import (
	"time"
)

type UserBasic struct {
	ID            uint      `gorm:"primarykey;cloumn:id" json:"id"`
	CreatedAt     time.Time `gorm:"cloumn:createAt" json:"createdAt"`
	UpdatedAt     time.Time `gorm:"cloumn:updatedAt" json:"updatedAt"`
	Name          string    `gorm:"cloumn:name" json:"name"`
	Password      string    `gorm:"cloumn:password" json:"password"`
	Phone         string    `gorm:"cloumn:phone" json:"phone"`
	Email         string    `gorm:"cloumn:email" json:"email"`
	Identity      string    `gorm:"cloumn:identity" json:"identity"`
	ClientIp      string    `gorm:"cloumn:clientIp" json:"clientIp"`
	ClientPort    string    `gorm:"cloumn:clientPort" json:"clientPort"`
	LoginTime     time.Time `gorm:"cloumn:loginTime" json:"loginTime"`
	HeartbeatTime time.Time `gorm:"cloumn:heartbeatTime" json:"heartbeatTime"`
	LoginOutTime  time.Time `gorm:"cloumn:loginOutTime" json:"loginOutTime"`
	IsLogout      bool      `gorm:"cloumn:isLogout" json:"isLogout"`
	DeviceInfo    string    `gorm:"cloumn:deviceInfo" json:"deviceInfo"`
}

//func (table *UserBasic) TableName() string {
//	return "user_basic"
//}
