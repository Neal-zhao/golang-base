package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type respData struct {
	resp *http.Response
	err error
}
func doCall(ctx context.Context)  {
	transport := http.Transport{}
	client := http.Client{
		Transport: &transport,
	}
	respChan := make(chan *respData, 1)
	//造一个请求对象
	req,err := http.NewRequest("GET","http://127.0.0.1:8080",nil)
	if err != nil {
		fmt.Printf(" new request failed ,err %s\n",err)
		return
	}
	req = req.WithContext(ctx)
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go func() {
		resp,err := client.Do(req)
		rd := &respData{
			resp: resp,
			err: err,
		}
		respChan <- rd
		wg.Done()
	}()
	for  {
		select {
		case result := <-respChan:
			fmt.Println("来了来了",result)
			if result.err != nil {
				fmt.Printf("err : %s",result.err)
				return
			}
			defer result.resp.Body.Close()
			data,_ := ioutil.ReadAll(result.resp.Body)
			fmt.Printf("resp %s \n",string(data))
		case <-ctx.Done():
			transport.CancelRequest(req)
			fmt.Println("超时结束")
		default:
			fmt.Println("default...")
		}
	}
}
func main()  {
	//模拟请求超时处理
	ctx,cancel := context.WithTimeout(context.Background(),time.Second * 1)
	defer cancel()
	doCall(ctx)
}