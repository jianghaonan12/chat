package main

import (
	"GINCHAT/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	db, _ := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	db.AutoMigrate(&models.UserBasic{})
	user := new(models.UserBasic)
	user.Name = "张三"
	db.Create(user)
}
