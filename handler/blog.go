package handler

import (
	"GoBlog/db"
	"GoBlog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func BlogListHandler(ctx *gin.Context) { //获取某用户的博客列表
	uid, err := strconv.Atoi(ctx.Param("uid")) //获取URL中的参数
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "无效的用户id"})
		return
	}
	blogs := db.GetBlogByUserIdList(uint(uid))
	utils.LogRus.Debugf("用户%d共有%d篇博客", uid, len(blogs))
	ctx.HTML(http.StatusOK, "blog_list.html", blogs) //给前端返回HTML页面和blogs结构体切片
}
func BlogDetailHandler(ctx *gin.Context) { //获取某篇博客的详情
	bid, err := strconv.Atoi(ctx.Param("bid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "无效的博客id"})
		return
	}
	blog := db.GetBlogById(uint(bid))
	if blog == nil {
		ctx.JSON(http.StatusOK, gin.H{"msg": "该博客不存在"})
		return
	}
	utils.LogRus.Debugf("该博客的正文:%s", blog.Article)
	ctx.HTML(http.StatusOK, "blog.html", gin.H{"bid": blog.ID, "title": blog.Title, "article": blog.Article, "update_time": blog.UpdateTime.Format("2006-01-02 15:04:05")}) //给前端返回一个map，.Format将日期转换为字符串
}
