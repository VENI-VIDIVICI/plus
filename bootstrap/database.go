package bootstrap

import (
	"fmt"

	"github.com/VENI-VIDIVICI/plus/pkg/config"
	"github.com/VENI-VIDIVICI/plus/pkg/database"
	_ "github.com/go-sql-driver/mysql"
)

func SetupDB() {
	// dsn := "root:123456@tcp(127.0.0.1:3308)/blog_plus"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Get("db.mysql.username"),
		config.Get("db.mysql.password"), config.Get("db.mysql.host"),
		config.Get("db.mysql.pory"), config.Get("db.mysql.database"))

	database.Connect(dsn)
}
