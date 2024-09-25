package main

import (
	"fmt"

	"github.com/VENI-VIDIVICI/plus/dao/mysql"
)

func main() {
	fmt.Println("hello, World!")
	mysql.Init()
	mysql.Query()
}
