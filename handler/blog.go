package handler

import (
	"GoBlog/db"
	"GoBlog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func BlogListHandler(ctx *gin.Context) {
	uid, err := strconv.Atoi(ctx.Param("uid")) //获取URL中的参数
	if err != nil {
		ctx.String(http.StatusBadRequest, "无效的uid")
		return
	}
	blogs := db.GetBlogByUserIdList(uint(uid))
	utils.LogRus.Debugf("用户%d共有%d篇博客", uid, len(blogs))
	ctx.HTML(http.StatusOK, "blog_list.html", blogs) //给前端返回HTML页面和blogs结构体切片
}
