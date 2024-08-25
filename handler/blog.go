package handler

import (
	"GoBlog/db"
	"GoBlog/models"
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
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "该博客不存在"})
		return
	}
	utils.LogRus.Debugf("该博客的正文:%s", blog.Article)
	ctx.HTML(http.StatusOK, "blog.html", gin.H{"bid": blog.ID, "title": blog.Title, "article": blog.Article, "update_time": blog.UpdateTime.Format("2006-01-02 15:04:05")}) //给前端返回一个map，.Format将日期转换为字符串
}

type UpdateBlogRequest struct { //校验修改博客时的参数
	BlogId  uint   `form:"bid" binding:"gt=0"`    //bid的值必须大于0
	Title   string `form:"title" binding:"min=1"` //title的字符串长度至少为1
	Article string `form:"article" binding:"min=1"`
}

func UpdateBlogHandler(ctx *gin.Context) { //更新博客
	var request UpdateBlogRequest
	err := ctx.ShouldBind(&request) //根据请求的内容（如URL参数、查询参数、表单数据或JSON）自动填充request结构体的字段，根据请求的Content-Type自动选择绑定方式
	if err != nil {
		ctx.String(http.StatusBadRequest, "无效参数")
		return
	}
	bid := request.BlogId
	title := request.Title
	article := request.Article
	blog := db.GetBlogById(bid)
	if blog == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "该博客不存在"})
		return
	}
	newBlog := models.Blog{ID: bid, Title: title, Article: article}
	if err = db.UpdateBlog(&newBlog); err != nil {
		utils.LogRus.Errorf("博客%d更新失败:%s", bid, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "博客更新失败"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "博客更新成功"})
}
