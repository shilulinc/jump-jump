package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//建立数据库连接
var dbclient *sql.DB

func init() {
	var err error
	dbclient, err = sql.Open("mysql", "root:123456@/jump-jump?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}
}
