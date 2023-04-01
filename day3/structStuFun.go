package main

import (
	"fmt"
	"os"
	"sort"
)

type student struct {
	name string
	class string
	id int
	sex int8
}
var (
	//studentLists map[int]student = make(map[int]student)
	studentLists = make(map[int]student)
	choic int
	del int = 4
)
//声明 初始化
//新增学生
func (t student) add()  {
	fmt.Println("添加学生")
	stu := t.userInput()
	if _,ok := studentLists[stu.id];!ok {
		studentLists[stu.id] = stu
		fmt.Println("学生添加成功：",stu.name)
		return
	}
	fmt.Println("学生已存在：",stu.name)
}
func (t student) update(stu student)  {}
func (t student) del()  {
	stu := t.userInput()
	if _,ok := studentLists[stu.id];ok {
		delete(studentLists,stu.id)
		return
	}
	fmt.Println("学生不存在")
}
func (t student) userInput() student {
	var (
		name string
		id int
		sex int8
		class string
	)

	if choic == del{
		fmt.Println("请输入学生 id")
		fmt.Scanln(&id)
	}else{
		fmt.Println("请输入学生 id")
		fmt.Scanln(&id)
		fmt.Println("请输入学生name")
		fmt.Scanln(&name)
		fmt.Println("请输入学生 age")
		fmt.Scanln(&sex)
		fmt.Println("请输入学生 class")
		fmt.Scanln(&class)
	}
	return student{
		name: name,
		class: class,
		sex: sex,
		id: id,
	}
}
func (t student) show() {
	fmt.Println("欢迎进入学生BGM")
	showInfo := map[int]string{
		1:"1、查看学生列表",
		2:"2、删除学生",
		3:"3、新增学生",
		4:"4、退出",
	}
	//showKey := []int{}
	var showKey []int
	for k := range showInfo{
		showKey = append(showKey,k)
	}
	sort.Ints(showKey)
	for i:=0;i<=len(showKey);i++ {
		fmt.Println(showInfo[i])
	}
}
func (t student) showStudent()  {
	for k,v := range studentLists{
		fmt.Println(k,v)
	}
}
func main()  {
	//studentLists = make(map[int]student)
	var stu student
	for  {
		stu.show()
		fmt.Scanln(&choic)
		switch choic {
		case 1:
			stu.showStudent()
		case 2:
			//这里是删除
			stu.del()
		case 3://这里是新增
			stu.add()
		case 4:
			os.Exit(0)
		}
		choic = 0
	}
}
