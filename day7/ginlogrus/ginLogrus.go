package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var log = logrus.New()
func initLogrus()  {
	log.Formatter = &logrus.JSONFormatter{}

	file,er := os.OpenFile("./logrus01.log",os.O_CREATE|os.O_WRONLY|os.O_WRONLY,0777)
	if er != nil {
		panic(er)
	}
	log.Out = file
	gin.DefaultWriter = log.Out
	//gin.SetMode(gin.ReleaseMode)
	log.Level = logrus.InfoLevel
}
func index(c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{
		"msg":"hello",
	})
}
func info(c *gin.Context)  {
	logrus.WithFields(logrus.Fields{
		"action":"info",
	}).Warn("测试信息")
	c.JSON(http.StatusOK,gin.H{
		"msg":"post",
	})
}
func main()  {
	initLogrus()
	r := gin.Default()
	r.GET("/index",index)
	r.GET("/post",info)
	r.Run(":999")
}