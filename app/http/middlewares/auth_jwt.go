package middlewares

import (
	"fmt"

	"github.com/VENI-VIDIVICI/plus/app/models/user"
	"github.com/VENI-VIDIVICI/plus/pkg/config"
	"github.com/VENI-VIDIVICI/plus/pkg/jwt2"
	"github.com/VENI-VIDIVICI/plus/pkg/response"
	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		calims, err := jwt2.NewJWT().ParserToken(c)
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("版本错误 % v", config.Get("app.name")))
			return
		}
		// calims.UserID
		userModel := user.Get(calims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(c, "找不到对应用户，用户可能已删除")
			return
		}
		c.Set("current_user_id", userModel.GetStringID())
		c.Set("current_user_name", userModel.Name)
		c.Set("current_user", userModel)
		c.Next()
	}
}

func AuthGuest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.GetHeader("Authorization")) > 0 {
			_, err := jwt2.NewJWT().ParserToken(ctx)
			if err != nil {
				response.Unauthorized(ctx, "请用游客模式访问")
				ctx.Abort()
				return
			}
			ctx.Next()
		}
	}
}
