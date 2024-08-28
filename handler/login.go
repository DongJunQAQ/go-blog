package handler

import (
	"GoBlog/db"
	"GoBlog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginResponse struct { //返回JSON数据的模板结构体
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Uid   uint   `json:"uid"`
	Token string `json:"token"`
}

const CookieName = "auth"

func LoginHandler(ctx *gin.Context) { //登录
	name := ctx.PostForm("user") //从前端传过来的表单中获取数据
	password := ctx.PostForm("pass")
	user := db.GetUserByName(name) //去数据库查一下有无该用户
	if len(name) == 0 {
		ctx.JSON(http.StatusBadRequest, LoginResponse{1, "用户名为空", user.ID, ""})
		return
	}
	if len(password) != 32 {
		ctx.JSON(http.StatusBadRequest, LoginResponse{2, "无效的密码", user.ID, ""})
		return
	}
	if user == nil {
		ctx.JSON(http.StatusForbidden, LoginResponse{3, "该用户不存在", 0, ""})
		return
	}
	if password != user.Password {
		ctx.JSON(http.StatusForbidden, LoginResponse{4, "密码错误", user.ID, ""})
		return
	}
	payload := utils.JwtPayload{Sub: "auth", Uid: user.ID, Iat: int(time.Now().Unix())}
	token, _ := utils.GenerateJWT(utils.DefaultHeader, payload, utils.Secret)
	ctx.SetCookie(CookieName, token, 86400, "/", "", false, true) //在设置响应内容之前设置Cookie，才能使Cookie被写入到浏览器
	//Secure:指示Cookie是否仅限于HTTPS连接，如果为false则通过HTTP连接发送
	//HttpOnly:指示Cookie是否仅限于HTTP请求，即不允许通过JavaScript访问
	ctx.JSON(http.StatusOK, LoginResponse{0, "登陆成功", user.ID, ""})
	utils.LogRus.Infof("用户%s(%d)登录成功", name, user.ID)
}
func GetUidFromCookie(ctx *gin.Context) uint { //JWT会被保存到客户端，因此无需将其写入Redis数据库
	for _, cookie := range ctx.Request.Cookies() {
		if cookie.Name == CookieName {
			if _, payload, err := utils.VerifyJWT(cookie.Value, utils.Secret); err == nil {
				return payload.Uid
			}
		}
	}
	return 0
}
