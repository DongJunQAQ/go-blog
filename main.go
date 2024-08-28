package main

import (
	"GoBlog/handler"
	"GoBlog/middleware"
	"GoBlog/utils"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Use(middleware.Metric())                 //全局中间件，每个handler执行之前都会执行该中间件
	router.GET("/metrics", func(ctx *gin.Context) { //该路由用来暴露Prometheus的指标数据
		promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request)
	})
	router.Static("/js", "view/js")                                                  //当用户访问http://yourdomain/js时，服务器会返回view/js目录下的所有文件
	router.StaticFile("/favicon.ico", "view/img/dqq.png")                            //当有请求访问/favicon.ico时，服务器会返回view/img/dqq.png这个文件
	router.LoadHTMLFiles("view/login.html", "view/blog_list.html", "view/blog.html") //加载HTML文件
	router.GET("login/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})
	router.POST("/login/submit", handler.LoginHandler)
	router.GET("/blog/list/:uid", middleware.Auth, handler.BlogListHandler) //:uid为URL中的参数，使用中间件来验证Cookie，用户请求该路由时会先执行middleware.Auth中间件再执行处理程序
	router.GET("/blog/:bid", middleware.Auth, handler.BlogDetailHandler)    //中间件就是在处理程序之前需要执行的handler
	router.POST("/blog/update", middleware.Auth, handler.UpdateBlogHandler)
	err := router.Run(":8080")
	if err != nil {
		utils.LogRus.Errorf("Gin启动失败")
	}
}
