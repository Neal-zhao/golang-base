package main

import "fmt"

func main()  {
	s1 := "hello"
	byteArr := []rune(s1)
	//byteArr := []byte(s1)
	s3 := ""
	for i:=0;i<len(byteArr);i++ {
		//强制类型转换
		s3 = s3 + string(byteArr[i])
	}
	fmt.Println(s3)
}
