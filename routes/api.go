package routes

import "github.com/gin-gonic/gin"

func RegiserRoutes(g *gin.Engine) {
	v1 := g.Group("v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
		})
	}
}
