package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main()  {
	//1、用rpc.Dial 与 rpc服务端建立连接
	client,err := rpc.Dial("tcp","127.0.0.1:8000")
	if err != nil {
		log.Printf("rpc.Dial error ：%s", err)
		return
	}
	defer client.Close()

	////1、调用远程函数
	serviceMethod := "World.Hello"		//struct.funName	服务名称.方法名称
	args := "我是客户端哦"	//客户端传参
	reply := ""				//服务端返回	需要传入地址 获取服务端返回的数据
	err = client.Call(serviceMethod,args,&reply)
	if err != nil {
		log.Printf("rpc.Dial error ：%s", err)
		return
	}
	fmt.Println(reply)

	var goodRes addGoodRes
	//err = client.Call("GoodsAbc.AllAddGoods",addGoodReq{ Id: 1, Name: "泳衣", Price: 33.59, ProNum: 100},&goodRes)
	err = client.Call("good.AllAddGoods",addGoodReq{ Id: 1, Name: "泳衣", Price: 33.59, ProNum: 100},&goodRes)
	if err != nil {
		log.Printf("rpc.Dial error ：%s", err)
		return
	}
	fmt.Println("addGoodRes：",goodRes)
}
type addGoodReq struct {
	Id int
	Name string
	Price float64
	ProNum int
}
type addGoodRes struct {
	Id int
	Content string
}