package main

import (
	"flag"
	"fmt"

	"github.com/VENI-VIDIVICI/plus/bootstrap"
	btsConfig "github.com/VENI-VIDIVICI/plus/config"
	"github.com/VENI-VIDIVICI/plus/pkg/config"
	"github.com/gin-gonic/gin"
)

func init() {
	btsConfig.Initialize()
}

func main() {
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	bootstrap.SetupLogger()
	config.InitConfig(env)
	bootstrap.SetupDB()
	router := gin.New()
	bootstrap.SetupRouter(router)
	fmt.Println(config.Get("app.port"))
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}

// bootstrap -> route 初始化路由
// routes -> api 注册路由
// pkg -> config -> config
