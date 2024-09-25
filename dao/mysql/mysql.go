package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init() {
	// localhost
	dsn := "root:123456@tcp(127.0.0.1:3308)/blog_plus"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return
	}
	fmt.Println("mysql 已经准备好")
	// db.SetMaxIdleConns()
}

func Close() {
	db.Close()
}
