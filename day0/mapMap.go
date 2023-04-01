package main

import "fmt"

func main()  {
	//student := make(map[string]map[string]int)
	//fmt.Println(student,len(student))
	student := map[string]map[interface{}]interface{}{
		"rongQi" : {"id":1,"age":8,"score":100},
		"yiMing" : {"id":2,"age":6,"score":100},
	}
	for k,v := range student{
		fmt.Println(k,v)
	}
	fmt.Println(student,len(student))
}
