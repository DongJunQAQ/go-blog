package main

import (
	"GoBlog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.html", nil)
}
func main() {
	router := gin.Default()
	router.LoadHTMLFiles("view/home.html") //加载HTML文件
	router.GET("/", Home)
	err := router.Run(":8080")
	if err != nil {
		utils.LogRus.Errorf("Gin启动失败")
	}
}
