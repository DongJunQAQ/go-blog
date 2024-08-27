package db

import (
	"GoBlog/models"
	"GoBlog/utils"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func GetBlogById(id uint) *models.Blog { //根据博客id查询博客信息
	db := ConnectMySQL()
	var blog models.Blog
	if err := db.Select("id,title,article,update_time").Where("id = ?", id).First(&blog).Error; err != nil {
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
func UpdateBlog(blog *models.Blog) error { //根据博客ID来修改博客的标题和正文
	if blog.ID <= 0 {
		return fmt.Errorf("无效的博客ID:%d", blog.ID)
	}
	if len(blog.Article) == 0 || len(blog.Title) == 0 {
		return fmt.Errorf("博客的内容或标题不可修改为空")
	}
	db := ConnectMySQL()
	return db.Model(models.Blog{}).Where("id=?", blog.ID).Updates(map[string]any{"title": blog.Title, "article": blog.Article}).Error
}
func CreateBlog(userId uint, title, article string) { //创建新博客
	db := ConnectMySQL()
	newBlog := models.Blog{UserID: userId, Title: title, Article: article}
	if err := db.Create(&newBlog).Error; err != nil {
		utils.LogRus.Errorf("创建博客%s失败:%s", title, err)
	} else {
		utils.LogRus.Infof("创建博客%s成功，博客ID为%d", title, newBlog.ID)
	}
}
