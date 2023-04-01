package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func mvcInfo(c *gin.Context)  {
	c.HTML(http.StatusOK,"index.html",gin.H{
		"msg" : "hello 我是超越",
	})
}
func xmlInfo(c *gin.Context)  {
	type user struct {
		Name string `json:"name"`
		Password string `json:"password"`
	}
	ret := user{"neal","666888"}
	c.XML(http.StatusOK,ret)
}
func queryInfo(c *gin.Context)  {
	c.DefaultQuery("name","Neal")
	c.DefaultPostForm("password","666888")
	c.Param("id")
	c.Get("sex")
}
func main()  {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/sb","./static")
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"msg":"hello 管大妈",
		})
	})
	router.GET("/mvc",mvcInfo)
	router.GET("/xml",xmlInfo)
	router.Run(":9090")
	fmt.Println("gin")
}