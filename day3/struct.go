package main

import (
	"fmt"
)

var (
	lists map[int]string = map[int]string{
		1:"添加",
		2:"修改",
		3:"展示图书",
		4:"退出",
	}
	bookLists map[string]*book
)
//Scanln Scanf Scan的选择
//传值 传指针
//封装方法
type book struct {
	name string
	author string
	price float32
	info string
} 
func displayBooks()  {
	var choice int
	var name string
	var author string
	var price float32
	var info string
	//showLists()
	for  {
		showLists()
		fmt.Scanln(&choice)
		if choice == 3 {
			fmt.Println("书籍列表")
			if len(bookLists) == 0 {
				fmt.Println("暂无书籍")
			}
			for bk,bv := range bookLists{
				fmt.Println(bk,bv)
			}
		}
		if choice == 1 {
			fmt.Println("添加书籍")
			fmt.Println("请输入书籍name")
			fmt.Scanln(&name)
			fmt.Println("请输入书籍author")
			fmt.Scanln(&author)
			fmt.Println("请输入书籍price")
			fmt.Scanln(&price)
			fmt.Println("请输入书籍info")
			fmt.Scanln(&info)
			fmt.Println(name,author,price,info)
			if _,ok := bookLists[name]; !ok {
				bookLists[name] = &book{
					name: name,
					author: author,
					price: price,
					info: info,
				}
				fmt.Printf("添加 %s 成功 \n",name)
			}
		}
		if choice == 2 {
			fmt.Println("修改书籍")
			fmt.Println("请输入书籍name")
			fmt.Scanln(&name)
			fmt.Println("请输入书籍author")
			fmt.Scanln(&author)
			fmt.Println("请输入书籍price")
			fmt.Scanln(&price)
			fmt.Println("请输入书籍info")
			fmt.Scanln(&info)
			if _,ok := bookLists[name]; ok {
				bookLists[name] = &book{
					name: name,
					author: author,
					price: price,
					info: info,
				}
				fmt.Printf("修改 %s 成功 \n",name)
			}else {
				fmt.Printf("修改的书籍 %s 不存在 \n",name)
			}
		}
		if choice == 4 {
			fmt.Println("退出界面")
			break
		}
	}
}
func showLists()  {
	for k,v := range lists{
		fmt.Println(k,":",v)
	}
}
func main()  {
	bookLists = make(map[string]*book)
	displayBooks()
}
