package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Worker(ctx context.Context)  {
	LOOP:
	for  {
		fmt.Println("db connect ...")
		time.Sleep(time.Millisecond * 10)
		select {
		//case <-time.Tick(time.Second * 1):
		//	fmt.Println("time tick...")
		case <-ctx.Done():
			break LOOP
			fmt.Println("ctx done 结束")
		default:

		}

	}
	fmt.Println(" go worker over... ")
	wg.Done()
}
var wg sync.WaitGroup
func main()  {
	//设置一个超时
	ctx,cancel := context.WithTimeout(context.Background(),time.Millisecond * 50)
	//goruntine
	wg.Add(1)
	go Worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over...")
	//通知结束
}
