package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f1(ch chan int)  {
	var int1 int
	for  {
		int1 = rand.Int()
		//fmt.Println(int1)
		ch <- int1
	}
}
func f2(ch chan int)  {
	var int1 int
	for  {
		int1 = rand.Intn(10000)
		//fmt.Println("f2",int1)
		ch <- int1
	}
}
func main()  {
	ch1,ch2 := make(chan int,100),make(chan int,100)
	go f1(ch1)
	go f2(ch2)
	for  {
		select {
		case ret1 := <- ch1:
			fmt.Println("ret1: ",ret1)
		case ret2 := <- ch2:
			fmt.Println("ret2: ",ret2)
		default:
			fmt.Println("无...")
			if(true) { time.Sleep(time.Second) }
		}
	}
}
