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
	if err := db.Where("name = ?", name).First(&user).Error; err != nil { //判断查询是否出错，如果有错再判断错误类型；为什么查询时要传入user的指针，因为在上一行定义的user目前是个空的结构体，查询后需要往这个空结构体内填入数据需要修改它，所以需要使用指针
		if !errors.Is(gorm.ErrRecordNotFound, err) { //如果不是"数据库无数据"的错误则打印日志
			utils.LogRus.Errorf("获取%s用户的数据失败:%s", name, err) //遇到了其他错误
		}
		return nil //如果是"数据库无数据"的错误则返回空值
	}
	return &user
}
func CreateUser(name, password string) { //创建用户
	db := ConnectMySQL()
	newUser := models.User{Name: name, Password: password}
	if err := db.Create(&newUser).Error; err != nil { //而在这里我需要将数据插入数据库并不需要修改newUser这个结构体，为什么这里也需要传入指针呢？这是因为ID字段为自增，Create()函数就会自动为newUser结构体插入id字段，因此该结构体发生了修改，所以需要使用指针
		utils.LogRus.Errorf("创建用户%s失败:%s", name, err)
	} else {
		utils.LogRus.Infof("创建用户%s成功，用户ID为%d", name, newUser.ID)
	}
}
func DeleteUserByName(name string) { //通过用户名删除用户
	db := ConnectMySQL()
	if err := db.Where("name = ?", name).Delete(models.User{}).Error; err != nil { //通过models.User{}结构体获取到对应的表名
		utils.LogRus.Errorf("删除用户%s失败:%s", name, err)
	}
}
