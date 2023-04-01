package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/ginchat2023"
	DB,err = sqlx.Connect("mysql",dsn)
	if err != nil {
		panic(err)
	}
	DB.SetMaxIdleConns(100)
	DB.SetMaxOpenConns(16)
	return
}