package main

import (
	"fmt"
	"reflect"
)

func main()  {
	type1 := reflect.TypeOf("abc")
	v1 := reflect.ValueOf("abc")
	fmt.Printf("%v %v %T",v1,type1,type1)
}
