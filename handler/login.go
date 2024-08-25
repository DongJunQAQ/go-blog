package handler

import (
	"GoBlog/db"
	"GoBlog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginResponse struct { //返回JSON数据的模板结构体
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Uid   int    `json:"uid"`
	Token string `json:"token"`
}

func LoginHandler(ctx *gin.Context) { //登录
	name := ctx.PostForm("user") //从前端传过来的表单中获取数据
	password := ctx.PostForm("pass")
	if len(name) == 0 {
		ctx.JSON(http.StatusBadRequest, LoginResponse{1, "用户名为空", 0, ""})
		return
	}
	if len(password) != 32 {
		ctx.JSON(http.StatusBadRequest, LoginResponse{2, "无效的密码", 0, ""})
		return
	}
	user := db.GetUserByName(name) //去数据库查一下有无该用户
	if user == nil {
		ctx.JSON(http.StatusForbidden, LoginResponse{3, "该用户不存在", 0, ""})
		return
	}
	if password != user.Password {
		ctx.JSON(http.StatusForbidden, LoginResponse{4, "密码错误", 0, ""})
		return
	}
	utils.LogRus.Infof("用户%s(%d)登录成功", name, user.ID)
	ctx.JSON(http.StatusOK, LoginResponse{0, "登陆成功", 0, ""})
	return
}
