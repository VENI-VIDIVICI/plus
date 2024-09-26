package bootstrap

import (
	"net/http"

	"github.com/VENI-VIDIVICI/plus/routes"
	"github.com/gin-gonic/gin"
)

func SetupRouter(g *gin.Engine) {
	registerGlobalMiddleWare(g)
	routes.RegiserRoutes(g)
	setup404Handler(g)
}

func registerGlobalMiddleWare(g *gin.Engine) {
	g.Use(gin.Recovery(), gin.Logger())
}

func setup404Handler(g *gin.Engine) {
	g.NoRoute(func(ctx *gin.Context) {
		acceptStr := ctx.Request.Header.Get("Accept")
		if acceptStr == "text/html" {
			ctx.String(http.StatusNotFound, "页面返回 404")
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code": 404,
				"error_msg":  "路由未定义，请确认 url 和请求方法是否正确",
			})
		}
	})
}
