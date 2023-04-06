package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"greeter/proto/greeter"
	"net"
)

type Hello struct {
}

func (h *Hello)SayHello(c context.Context,req *greeter.HelloReq) (*greeter.HelloRes,error) {
	fmt.Println(req)
	return &greeter.HelloRes{
		Message: "你好"+req.Name,
	},nil
}
func main()  {
	// 初始化 grpc 对象
	grpcServer := grpc.NewServer()

	//	注册服务
	greeter.RegisterGreeterServer(grpcServer,&Hello{})

	//设置监听
	listen,err := net.Listen("tcp","127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	grpcServer.Serve(listen)
}
