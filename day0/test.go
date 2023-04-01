package main

import (
	"fmt"
	"strings"
	"time"
)
var (
	pi = 3.1415926
)

func main()  {
	//fmt.Println(pi)
	//var pi = 3.1415927
	pi :=3.1133
	fmt.Println(pi)
	now := time.Now()
	UnixNano := now.UnixNano()
	fmt.Println("now",now)
	fmt.Println("unixNano",int(UnixNano),UnixNano)
	url := "http://www.baidu.com"
	lastIndex := strings.LastIndex(url,"/")
	fmt.Println("lastIndex",lastIndex,url[lastIndex+1:])
}
