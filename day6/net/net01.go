package main

import (
	"fmt"
	"net"
)

func main()  {
	//监听
	listenner,_ := net.Listen("tcp","127.0.0.1:2000")
	//建立连接
	for  {
		conn,err := listenner.Accept()//阻塞接收
		if err != nil {

		}
		//创建处理连接
		go process(conn)
	}
}
func process(conn net.Conn)  {
	defer conn.Close()
	var buf [1024]byte
	n,err := conn.Read(buf[:])
	if err != nil {

	}
	fmt.Println("获取到的消息：", n, string(buf[:]))
}
