package main

import (
	"fmt"
	"strings"
)

func main()  {
	a1 := [...]int{1,2,3}
	b1 := a1
	b1[0] = 111
	fmt.Println(a1,b1)
	abcSlice := strings.Split("a,b,c,d,d,e",",")
	//var abcCount map[string]int
	abcCount := make(map[string]int)
	for _,v := range abcSlice {
		if _,ok := abcCount[v]; !ok {
			abcCount[v] = 1
			continue
		}
		abcCount[v] += 1
	}
	fmt.Println(abcCount)
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("defer err",err)
		}
	}()
	panic("xxx error")
	return
	var mapSlice []map[string]int
	mapSlice = make([]map[string]int,10)
	mapSlice = append(mapSlice, map[string]int{
		"a":1,
		"b":1,
	})
	mapSlice[0]["a"] = 1
	mapSlice[0]["b"] = 11
	fmt.Println(mapSlice)
}
