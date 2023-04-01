package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main()  {
	conn,err := net.Dial("tcp","127.0.0.1:2000")
	if err != nil {
		return
	}
	defer conn.Close()

	var input string
	fmt.Scan(&input)
	reader := bufio.NewReader(os.Stdin)
	input2,err := reader.ReadString('\n')
	_,err = conn.Write([]byte(input2))
	//_,err = conn.Write([]byte("来么"))
	if err != nil {
		return
	}
}
