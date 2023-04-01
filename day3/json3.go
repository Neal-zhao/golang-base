package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	type student struct {
		Name string `json:"name"`
		Age int `json:"age"`
	}

	stu := student{
		"张三",
		22,
	}
	bt,err := json.Marshal(stu)
	fmt.Println(string(bt),err)
	//var out student
	out := student{}
	json.Unmarshal([]byte("{\"Name\":\"张三\",\"Age\":22}"),&out)
	fmt.Println(out.Name)
}
