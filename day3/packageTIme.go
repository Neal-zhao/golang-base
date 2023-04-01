package main

import (
	"fmt"
	"time"
)

func printTime(t time.Time)  {
	timeStr := t.Format("2006-01-02 15:04:05")
	fmt.Println(timeStr)
}
func timeTick()  {
	tick := time.Tick(time.Second * 60)
	for t := range tick{
		fmt.Println(t)
	}
}
func countTime(){
	start := time.Now().UnixNano() / 1000
	fmt.Println("千里江陵一日还")
	time.Sleep(time.Millisecond * 30)
	end := time.Now().UnixNano() / 1000
	fmt.Println("运行了微妙：",end - start)
	time.Now().Year()
	time.Now().Date()
}
func main()  {
	timeTick()
	tnow := time.Now()
	printTime(tnow)
	countTime()

}
