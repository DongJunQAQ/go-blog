package db

import (
	"GoBlog/models"
	"GoBlog/utils"
	"errors"
	"gorm.io/gorm"
)

func GetUserByName(name string) *models.User { //根据用户名查询用户信息
	db := ConnectMySQL()
	var user models.User
	if err := db.Where("name = ?", name).First(&user).Error; err != nil { //判断查询是否出错，如果有错再判断错误类型
		if !errors.Is(gorm.ErrRecordNotFound, err) { //如果不是"数据库无数据"的错误则打印日志
			utils.LogRus.Errorf("获取%s用户的数据失败:%s", name, err) //遇到了其他错误
		}
		return nil //如果是"数据库无数据"的错误则返回空值
	}
	return &user
}
