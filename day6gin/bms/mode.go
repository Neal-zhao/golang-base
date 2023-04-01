package main

import "fmt"

type Book struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Price float64 `json:"price"`
}

func QueryAllBook() (allBook []*Book, err error) {
	sqlStr := "select * from book"
	err = db.Select(&allBook,sqlStr)
	if err != nil {
		fmt.Println("查询所有书本失败")
		return
	}
	return
}

func insertBook(title string, price float64)  (err error) {
	sql := "INSERT INTO book(title,price) values(?,?)"
	result,err := db.Exec(sql,title,price)
	if err != nil {
		fmt.Printf("inert into err: %s ",err)
		return err
	}
	_,err = result.LastInsertId()
	if err != nil {
		fmt.Printf("LastInsertId err: %s ",err)
		return err
	}
	return err
}

func deleteBook2(id int64)  {
	sql := "delete from book where id=?"
	fmt.Println(sql,id)
	_,err := db.Exec(sql,id)
	if err != nil {
		fmt.Printf("deleteBook err: %s ",err)
		return
	}
	return
}