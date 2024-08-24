package main

import (
	"GoBlog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}
func main() {
	router := gin.Default()
	router.Static("/js", "view/js")                                                  //当用户访问http://yourdomain/js时，服务器会返回view/js目录下的所有文件
	router.StaticFile("/favicon.ico", "view/img/dqq.png")                            //当有请求访问/favicon.ico时，服务器会返回view/img/dqq.png这个文件
	router.LoadHTMLFiles("view/login.html", "view/blog_list.html", "view/blog.html") //加载HTML文件
	router.GET("login/", LoginView)
	err := router.Run(":8080")
	if err != nil {
		utils.LogRus.Errorf("Gin启动失败")
	}
}
