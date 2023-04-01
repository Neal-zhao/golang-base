package main

import (
	"context"
	"time"
)

func main()  {
	ctx,cancel := context.WithCancel(context.Background())
	context.WithDeadline(context.Background(),time.Now())
}
