package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	websocker "github.com/gorilla/websocket"
)

type Client struct {
	id string
	socket *websocker.Conn
	send chan []byte
}
func main()  {
	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)

 	//client := &Client{id: uuid.NewV4().String(), socket: conn, send: make(chan []byte)}
}
