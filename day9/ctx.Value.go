package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)
func worker2(ctx context.Context)  {
	key := TraceCode("TRACE_CODE")
 	traceCode,ok := ctx.Value(key).(string)
	fmt.Println(ok,traceCode)
	if !ok {
		fmt.Println("invalid trace code ...")
	}
	log.Printf("worker trace code :%s\n",traceCode)
Loop:
	for  {
		fmt.Println("db connect ...")
		//time.Sleep(time.Second * 1)
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break Loop
			fmt.Println("ctx.Done over...")
		default:
		}
	}
	wgV.Done()
	fmt.Println("worker over ....")
}
var wgV sync.WaitGroup
type TraceCode string
func main()  {
	//设置一个50毫秒的超时
	ctx,cancel := context.WithTimeout(context.Background(),time.Millisecond * 50)
	//设置变量传递个goroutine
	ctx = context.WithValue(ctx,TraceCode("TRACE_CODE"),"abc111")
	log.Printf("%s main fun","abc111")
	wgV.Add(1)
	go worker2(ctx)
	//休息5秒
	time.Sleep(time.Second * 5)
	//通知 结束
	cancel()
	wgV.Wait()
	fmt.Println("main over ....")
}
