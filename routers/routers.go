package routers

import "github.com/gin-gonic/gin"

func SetupRouter() {
	r := gin.New()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}
