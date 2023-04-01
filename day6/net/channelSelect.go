package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	var chan1 chan int
	chan1 = make(chan int,1)
	id := 1
	go func() {
		for  id < 1000 {
			id++
			rand1 := rand.Intn(100)
			chan1 <-rand1
		}
		close(chan1)
	}()
	for i := range chan1{
		fmt.Println(i)
	}
	//for  {
	//	select {
	//	case res := <-chan1:
	//		fmt.Println(res)
	//	//default:
	//	//	fmt.Println("default")
	//	}
	//
	//}
}
