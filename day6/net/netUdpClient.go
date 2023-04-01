package main

import (
	"fmt"
	"net"
)

func main()  {
	conn,err := net.Dial("udp","127.0.0.1:3000")
	if err != nil {
		fmt.Println("err 11")
		return
	}
	defer conn.Close()
	var input string
	_,_ = fmt.Scanf("%s",&input)
	//n,err := conn.Write([]byte("hello "))
	n,err := conn.Write([]byte(input))
	if err != nil {
		fmt.Println("err 17")
		return
	}
	var buf [1024]byte
	n,err = conn.Read(buf[:])
	if err != nil {
		fmt.Println("err 23")
	}
	fmt.Println(string(buf[:n]))
}
