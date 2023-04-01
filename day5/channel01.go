package main

import (
	"fmt"
	"math/rand"
	"time"
)

type item struct {
	id int64
	num int64
}
type result struct {
	item
	sum int64
}
var ch chan *item
var chResult chan *result

func producer(ch chan *item )  {
	var id int64
	for id < 20 {
		id++
		number := rand.Int63()
		number = number % 10
		//fmt.Println("number : ",number)
		tmp := &item{ id: id, num: number}
		ch<-tmp
	}
	close(ch)
}
func consumer(ch chan *item,chResult chan *result) {
	//生产值 得到 消费了
	//tmp := <- ch
	for tmp := range ch {
		sum := calc(tmp.num)
		//fmt.Println("tmp : ",tmp,sum)
		//消费者 消费者卖了多少钱
		result := &result{
			item:*tmp,
			sum:sum,
		}
		chResult <- result
	}
	close(chResult)
}
func calc(num int64) ( sum int64)  {
	//var sum int64
	for num > 0 {
		sum = sum + num % 10
		num = num / 10
 	}
	return sum
}
func printResult(result chan *result )  {
	//for res := range result{
}

func main()  {
	//
	ch := make(chan *item,100)
	chResult1 := make(chan *result,100)
	go producer(ch)
	for i:=1;i<=20;i++ {
		
	}
	go consumer(ch,chResult1)

	printResult(chResult1)
	//for  {
	//	res := <- chRes
	//	fmt.Println(&res)
	//}
	for {
		res := <- chResult1
		if res == nil {
			break
		}
		fmt.Printf("num:%v id:%v sum:%v\n",res.num,res.id,res.sum)
		if(1==0) {
			time.Sleep(time.Second)
		}
	}
}
