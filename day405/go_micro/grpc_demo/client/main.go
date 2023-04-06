package main

import (
	"client/proto/greeter"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main()  {
	//fmt.Println(123)
	//	建立连接
	clientConn,err := grpc.Dial("127.0.0.1:8080",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer clientConn.Close()

	//	注册客户端
	client := greeter.NewGreeterClient(clientConn)
	res,err := client.SayHello(context.Background(),&greeter.HelloReq{
		Name: "张三",
	})
	fmt.Printf("%#v\r\n",res)
	fmt.Println( res.Message)
}