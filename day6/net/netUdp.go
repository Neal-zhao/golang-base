package main

import (
	"fmt"
	"net"
)

func main()  {
	listenner,err := net.ListenUDP("udp",&net.UDPAddr{
		IP: net.ParseIP("127.0.0.1"),
		Port: 3000,
	})
	if err != nil {
		fmt.Println("err 14")
		return
	}
	defer listenner.Close()
	for  {
		var buf [1024]byte
		n,addr,err := listenner.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("err 20")
			return
		}
		fmt.Printf("%v\n",addr,string(buf[:n]))
		n,err = listenner.WriteToUDP([]byte("去吧"),addr)
		if err != nil {
			fmt.Println("err 26")
		}
	}
}
