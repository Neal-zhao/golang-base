package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/ginchat2023"
	db,err = sqlx.Connect("mysql",dsn)
	if err != nil {
		return err
	}
	db.SetMaxIdleConns(16)
	db.SetMaxOpenConns(100)
	return nil
}