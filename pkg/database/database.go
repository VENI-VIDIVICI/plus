package database

import (
	"database/sql"
	"fmt"

	"github.com/VENI-VIDIVICI/plus/pkg/config"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect(dsn string) {
	var err error
	DB, err = sql.Open(config.Get("db.connection"), dsn)
	if err != nil {
		fmt.Println("mysql link is err:", err)
		return
	}
}
