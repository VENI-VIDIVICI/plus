package middlewares

import (
	"github.com/VENI-VIDIVICI/plus/pkg/app"
	"github.com/VENI-VIDIVICI/plus/pkg/limiter"
	"github.com/VENI-VIDIVICI/plus/pkg/logger"
	"github.com/VENI-VIDIVICI/plus/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func LinitIP(limit string) gin.HandlerFunc {
	if app.IsTesting() {
		limit = "10000-H"
	}
	return func(c *gin.Context) {
		key := limiter.GetKeyIP(c)
		if ok := limitHandler(c, key, limit); !ok {
			return
		}
		c.Next()
	}
}

func LimitFullPath(limit string) gin.HandlerFunc {
	if app.IsTesting() {
		limit = "10000-H"
	}
	return func(c *gin.Context) {
		key := limiter.GetKeyRouteWithIP(c)
		if ok := limitHandler(c, key, limit); !ok {
			return
		}
		c.Next()
	}
}
func limitHandler(c *gin.Context, key string, limit string) bool {
	rate, err := limiter.CheckRate(c, key, limit)
	if err != nil {
		logger.LogIf(err)
		response.Abort500(c)
		return false
	}
	c.Header("X-RateLimit-Limit", cast.ToString(rate.Limit))
	c.Header("X-RateLimit-Remaining", cast.ToString(rate.Remaining))
	c.Header("X-RateLimit-Reset", cast.ToString(rate.Reset))
	if rate.Reached {
		response.Abort500(c, "接口请求太频繁")
		return false
	}
	return true

}
