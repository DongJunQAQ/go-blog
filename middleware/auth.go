package middleware

import (
	"GoBlog/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(ctx *gin.Context) {
	if cookieUid := handler.GetUidFromCookie(ctx); cookieUid != "" {
		ctx.Next() //如果符合上面的条件则继续执行后续的中间件或处理程序
	} else {
		ctx.String(http.StatusForbidden, "请先登录")
		ctx.Abort() //如果不符合则停止后续中间件的执行并直接终止请求处理
	}
}
