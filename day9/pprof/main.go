package main

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	ginpprof.Wrap(router)
	router.Run(":8080")
}
