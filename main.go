package main

import (
	"fmt"

	"github.com/VENI-VIDIVICI/plus/dao/mysql"
	"github.com/VENI-VIDIVICI/plus/routers"
)

func main() {
	fmt.Println("hello, World!")
	mysql.Init()
	mysql.Query()
	routers.SetupRouter()
}
