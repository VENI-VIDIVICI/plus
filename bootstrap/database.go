package bootstrap

import (
	"fmt"

	"github.com/VENI-VIDIVICI/plus/pkg/config"
	"github.com/VENI-VIDIVICI/plus/pkg/database"
	_ "github.com/go-sql-driver/mysql"
)

func SetupDB() {
	// dsn := "root:123456@tcp(127.0.0.1:3308)/blog_plus"
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", config.Get("database.mysql.username"),
		config.Get("database.mysql.password"), config.Get("database.mysql.host"),
		config.Get("database.mysql.port"), config.Get("database.mysql.database"))

	fmt.Println(dsn)
	database.Connect(dsn)
}
