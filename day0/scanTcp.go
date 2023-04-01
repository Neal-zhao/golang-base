package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func main()  {
	startPort := flag.Int("startPort",1,"start port")
	endPort := flag.Int("endPort",100,"end port")
	timeOut := flag.Duration("tinmeOut",time.Millisecond * 200," time out")
	hostName := flag.String("hostName","baidu.com"," host name")
	flag.Parse()
	ports := []int{}

	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	//timeOut := time.Millisecond * 200
	for port:=*startPort;port<*endPort;port++ {
		wg.Add(1)
		go func(p int) {
			open := isOpen(*hostName,p,*timeOut)
			if open == true {
				mutex.Lock()
				ports = append(ports,p)
				mutex.Unlock()
			}
			wg.Done()
		}(port)
	}
	wg.Wait()
	fmt.Printf("Open ports %v\n",ports)
}
func isOpen(host string,port int,timeOut time.Duration) bool {
	time.Sleep(time.Millisecond * 1)
	//conn,err := net.Dial("tcp",fmt.Sprintf("%s:%d",host,port))
	conn,err := net.DialTimeout("tcp",fmt.Sprintf("%s:%d",host,port),timeOut)
	if err  == nil {
		_ = conn.Close()
		return true
	}
	return false
}
