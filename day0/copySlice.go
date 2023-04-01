package main

import "fmt"

func main()  {
	var arr2 = [...]int{4,5,6}
	slice2 := arr2[1:]
	slice2[0] = 155
	slice2 = append(slice2,1,2,3,4,5,6,7)
	slice2[0] = 551
	fmt.Println(slice2,arr2,cap(arr2),cap(slice2))
	var arr = []int{1,2,3}
	b := arr

	e := arr[1:]
	fmt.Println(e,"\n")
	e[0] = 22
	fmt.Println(e,arr,cap(arr),cap(e))
	var c []int
	c = make([]int,3,3)
	//c := []int{0,0}
	copy(c,arr)
	b[0] = 100
	c[0] = 12345
	arr = append(arr,11,22)
	b[0] = 123
	fmt.Println(arr,"\n",b,"\n",c)

}
