package db

import (
	"GoBlog/models"
	"GoBlog/utils"
	"errors"
	"gorm.io/gorm"
)

func GetBlogById(id uint) *models.Blog { //根据博客id查询博客信息
	db := ConnectMySQL()
	var blog models.Blog
	if err := db.Select("id,user_id,title,article").Where("id = ?", id).First(&blog).Error; err != nil {
		if !errors.Is(gorm.ErrRecordNotFound, err) {
			utils.LogRus.Errorf("获取博客%d的数据失败:%s", id, err)
		}
		return nil
	}
	return &blog
}
func GetBlogByUserIdList(userId uint) []*models.Blog { //根据用户id查询相关博客信息
	db := ConnectMySQL()
	var blogs []*models.Blog
	if err := db.Select("id,title").Where("user_id = ?", userId).Find(&blogs).Error; err != nil {
		if !errors.Is(gorm.ErrRecordNotFound, err) {
			utils.LogRus.Errorf("获取用户%d的博客失败:%s", userId, err)
		}
		return nil
	}
	return blogs
}
