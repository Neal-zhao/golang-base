package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)
type addGoodRes1 struct {
	Id int
	Content string
}
func main()  {
	//建立 net tcp 连接
	conn,err := net.Dial("tcp","127.0.0.1:8000")
	if err != nil {
		fmt.Println("net Dial err: ",err)
		return
	}

	//使用 json 格式作为 rpc 传输格式
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	//rpc 调用
	var reply addGoodRes1
	err = client.Call("good.AllAddGoods",nil,&reply)
	if err != nil {
		fmt.Println("net Dial err: ",err)
		return
	}
	fmt.Printf("%s - %v - %T",reply,reply,reply)
}
