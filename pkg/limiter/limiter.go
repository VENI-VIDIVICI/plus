package limiter

import (
	"strings"

	"github.com/VENI-VIDIVICI/plus/pkg/config"
	"github.com/VENI-VIDIVICI/plus/pkg/logger"
	"github.com/VENI-VIDIVICI/plus/pkg/redis"
	"github.com/gin-gonic/gin"
	limiterlib "github.com/ulule/limiter/v3"
	limiterRedis "github.com/ulule/limiter/v3/drivers/store/redis"
)

func GetKeyIP(c *gin.Context) string {
	return c.ClientIP()
}

func GetKeyRouteWithIP(c *gin.Context) string {
	return routeToKeyString(c.FullPath()) + c.ClientIP()
}

func CheckRate(c *gin.Context, key string, formatted string) (limiterlib.Context, error) {
	var context limiterlib.Context
	rate, err := limiterlib.NewRateFromFormatted(formatted)
	if err != nil {
		logger.LogIf(err)
		return context, err
	}
	// store, err := limiterRedis.NewStore(&redis.RedisInstance.Client)
	store, err := limiterRedis.NewStoreWithOptions(&redis.RedisInstance.Client, limiterlib.StoreOptions{
		Prefix: config.Get("app.name") + ":limiter",
	})
	if err != nil {
		logger.LogIf(err)
		return context, err
	}
	limiterObj := limiterlib.New(store, rate)

	if c.GetBool("limiter-once") {
		return limiterObj.Peek(c, key)
	} else {
		c.Set("limiter-once", true)
		return limiterObj.Get(c, key)
	}
}

func routeToKeyString(routeName string) string {
	routeName = strings.ReplaceAll(routeName, "/", "-")
	routeName = strings.ReplaceAll(routeName, ":", "_")
	return routeName
}
