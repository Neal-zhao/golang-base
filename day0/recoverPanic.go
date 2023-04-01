package main

import "fmt"

func main()  {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("get err info ",err)
		}
	}()
	var a []int
	a = make([]int,1)
	a[0] = 10
	b := []int{1,23,4}
	fmt.Println(a,b)
}
