package main

import "fmt"

func getType(x interface{})  {

	switch x.(type) {
	case  string:
		fmt.Println("这是 类型string")
	}
}

type animal2 interface {
	say()
}
type dog struct {
	name string
}
func (d dog) say()  {
	fmt.Printf("%s 说汪汪汪\n",d.name)
}
type cat struct {
	name string
}
func (c *cat) say()  {
	fmt.Printf("%s 说汪汪汪\n",c.name)
}
func main()  {
	getType("哈哈哈")
	var an animal2
	dog := dog{name: "旺财"}
	cat := cat{name: "喵喵"}
	an = dog
	an.say()
	an = &dog
	an.say()
	an = &cat	//只能使用指针接收者
	an.say()
}
