package main

import (
	"fmt"
	"io"
	"net"
)

func main()  {
	//connect
	//写入 Get请求
	//接收响应
	//为什么用for
	conn,err := net.Dial("tcp","www.liwenzhou.com:80")
	if err != nil {
		fmt.Println("conn err")
		return
	}
	defer conn.Close()

	fmt.Fprintf(conn,"GET / HTTP/1.0\r\n\r\n")
	var buf [1024*10]byte
	for  {
		n,err := conn.Read(buf[:])
		if err == io.EOF {
			fmt.Println("读取结束")
			return
		}
		fmt.Printf("%s :",string(buf[:n]))

	}
}
